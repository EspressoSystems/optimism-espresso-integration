//! Batch Authentication via L1 Event Scanning
//!
//! This module implements event-based batch authentication for the Espresso integration.
//! Instead of relying on an on-chain `BatchInbox` contract to verify batches, the derivation
//! pipeline scans L1 receipts for `BatchInfoAuthenticated(bytes32 indexed commitment)` events
//! emitted by the `BatchAuthenticator` contract within a lookback window.
//!
//! Whether batches must be authorized by an event is gated by the
//! [`espresso_time`](kona_genesis::HardForkConfig::espresso_time)
//! hardfork. The fork is conceptually an L2-timestamp hardfork, but the per-L1-block decision
//! made at the data source layer is gated on the L1 origin time of the block being scanned —
//! mirroring the upstream `ecotoneTime` precedent.
//!
//! - **Pre-fork (or fork unset):** the pipeline runs vanilla OP Stack semantics. A batch is
//!   authorized iff its sender matches `batcher_address`. The `BatchAuthenticator` event
//!   lookback is bypassed entirely.
//! - **Post-fork:** a batch is authorized iff its commitment hash matches a
//!   `BatchInfoAuthenticated(bytes32 indexed commitment)` event emitted by the configured
//!   `BatchAuthenticator` contract within the lookback window. Sender-based fallback is
//!   rejected.
//!
//! This matches the verifier-side semantics of upstream Go PR #409
//! (`EspressoSystems/optimism-espresso-integration`).
//!
//! Using event scanning (rather than L1 contract state reads) keeps the derivation pipeline
//! compatible with the op-program fault proof environment, which can only access L1 block headers,
//! transactions, receipts, and blobs — not contract state.

use crate::ChainProvider;

use alloc::{collections::BTreeSet, vec::Vec};
use alloy_consensus::transaction::SignerRecoverable;
use alloy_consensus::{Receipt, TxEnvelope, TxReceipt};
use alloy_primitives::{Address, B256, b256, keccak256};
use kona_protocol::BlockInfo;
use lru::LruCache;

/// The `keccak256("BatchInfoAuthenticated(bytes32)")` event topic.
///
/// This is the event emitted by the `BatchAuthenticator` contract when a batch is authenticated.
/// The first indexed topic is the commitment hash.
pub(crate) const BATCH_INFO_AUTHENTICATED_TOPIC: B256 =
    b256!("ee0d07d204d979d28885955e59a46f754c4db7378b7df1a95123525aac6e3f80");

/// Configuration for event-based batch authentication.
#[derive(Debug, Clone)]
pub struct BatchAuthConfig {
    /// The L1 address of the `BatchAuthenticator` contract.
    pub authenticator_address: Address,
}

/// Computes `keccak256(calldata)`, matching the `BatchAuthenticator` contract's calldata batch
/// validation path.
pub(crate) fn compute_calldata_batch_hash(data: &[u8]) -> B256 {
    keccak256(data)
}

/// Computes `keccak256(concat(blob_hashes))`, matching the `BatchAuthenticator` contract's blob
/// batch validation path.
pub(crate) fn compute_blob_batch_hash(blob_hashes: &[B256]) -> B256 {
    let mut concatenated = Vec::with_capacity(32 * blob_hashes.len());
    for hash in blob_hashes {
        concatenated.extend_from_slice(hash.as_slice());
    }
    keccak256(&concatenated)
}

/// Extracts all authenticated batch commitment hashes from a single block's receipts.
///
/// Scans for `BatchInfoAuthenticated` events emitted by `authenticator_addr` in successful
/// receipts. Returns the set of commitment hashes found.
pub(crate) fn collect_auth_events_from_receipts(
    receipts: &[Receipt],
    authenticator_addr: Address,
) -> BTreeSet<B256> {
    let topic0 = BATCH_INFO_AUTHENTICATED_TOPIC;
    let mut result = BTreeSet::new();
    for receipt in receipts {
        if !receipt.status() {
            continue;
        }
        for log in &receipt.logs {
            if log.address != authenticator_addr {
                continue;
            }
            if log.topics().len() >= 2 && log.topics()[0] == topic0 {
                result.insert(log.topics()[1]);
            }
        }
    }
    result
}

/// Scans L1 receipts in the range `[block_ref.number - lookback_window, block_ref.number]`
/// and returns the set of all batch commitment hashes that were authenticated via
/// `BatchInfoAuthenticated` events.
///
/// This is called once per L1 block by the data source, and the returned set is checked
/// against each candidate batch transaction. This avoids rescanning the lookback window
/// for every individual batch transaction.
///
/// Results are cached per block hash in the provided LRU cache. For consecutive L1 blocks
/// the lookback windows overlap by `lookback_window - 1` blocks, so only one new block's
/// receipts need to be fetched on each call. The cache is keyed by block hash (not number)
/// so it is naturally reorg-safe.
pub(crate) async fn collect_authenticated_batches<CP: ChainProvider + Send>(
    provider: &mut CP,
    block_ref: &BlockInfo,
    authenticator_addr: Address,
    lookback_window: u64,
    cache: &mut BatchAuthCache,
) -> Result<BTreeSet<B256>, CP::Error> {
    let mut all_authenticated = BTreeSet::new();
    let mut current_hash = block_ref.hash;
    let mut current_number = block_ref.number;

    loop {
        // Check receipt cache first
        if let Some(cached) = cache.receipts.get(&current_hash) {
            all_authenticated.extend(cached.iter());
        } else {
            // Cache miss: fetch receipts, extract events, cache the result
            let receipts = provider.receipts_by_hash(current_hash).await?;
            let events = collect_auth_events_from_receipts(&receipts, authenticator_addr);
            all_authenticated.extend(events.iter());
            cache.receipts.put(current_hash, events);
        }

        if current_number == 0 || block_ref.number - current_number >= lookback_window {
            break;
        }

        // Walk backward using header to get parent hash
        let parent_hash = if let Some(&cached_parent) = cache.headers.get(&current_hash) {
            cached_parent
        } else {
            let header = provider.header_by_hash(current_hash).await?;
            cache.headers.put(current_hash, header.parent_hash);
            header.parent_hash
        };
        current_hash = parent_hash;
        current_number = current_number.saturating_sub(1);
    }

    Ok(all_authenticated)
}

/// LRU caches used during the batch authentication lookback window traversal.
///
/// Bundles the receipt-event cache and the header (block hash → parent hash) cache.
/// Both caches are sized slightly larger than the configured lookback window to avoid
/// thrashing at the boundary.
#[derive(Debug, Clone)]
pub(crate) struct BatchAuthCache {
    /// Authenticated batch commitment hashes extracted from receipts, keyed by L1 block hash.
    pub receipts: LruCache<B256, BTreeSet<B256>>,
    /// Block parent hashes keyed by block hash
    pub headers: LruCache<B256, B256>,
}

impl BatchAuthCache {
    /// Creates a new [`BatchAuthCache`] with both caches sized to `lookback_window + 2`.
    pub(crate) fn new(lookback_window: u64) -> Self {
        let cap =
            lookback_window.try_into().map(|w: usize| w.saturating_add(2)).unwrap_or(usize::MAX);
        let cap = core::num::NonZeroUsize::new(cap).expect("cache size must be non-zero");
        Self { receipts: LruCache::new(cap), headers: LruCache::new(cap) }
    }
}

/// Checks whether a batch transaction is authorized.
///
/// Behaviour is gated by `espresso_active` (computed by the caller from
/// `RollupConfig::is_espresso_active(block_ref.timestamp)` where `block_ref` is
/// the L1 origin of the block being scanned):
///
/// - **`espresso_active = false`** (pre-fork / vanilla OP Stack): authorized iff the
///   transaction sender matches `batcher_address`. `auth_config` and `authenticated_hashes` are
///   ignored.
/// - **`espresso_active = true`** (post-fork): authorized iff a `BatchAuthenticator` is
///   configured (`auth_config.is_some()`) AND `batch_hash` is present in
///   `authenticated_hashes`. Sender-based fallback is rejected.
///
/// If the gap between the authentication transaction and the batch data exceeds the
/// [`BATCH_AUTH_LOOKBACK_WINDOW`], it's the batcher's responsibility to detect this and
/// re-submit the authentication transaction and batch data.
pub(crate) fn is_batch_authorized(
    tx: &TxEnvelope,
    batch_hash: B256,
    auth_config: Option<&BatchAuthConfig>,
    authenticated_hashes: &BTreeSet<B256>,
    batcher_address: Address,
    espresso_active: bool,
) -> bool {
    if espresso_active {
        // Post-fork: event-only authorization. Sender-based fallback is rejected.
        return auth_config.is_some() && authenticated_hashes.contains(&batch_hash);
    }
    // Pre-fork: vanilla OP Stack sender verification.
    tx.recover_signer().map(|sender| sender == batcher_address).unwrap_or(false)
}

#[cfg(test)]
mod tests {
    use super::*;
    use alloc::{vec, vec::Vec};
    use alloy_consensus::{Eip658Value, Receipt, Signed, TxLegacy};
    use alloy_primitives::{Address, Log, LogData, Signature, TxKind, address, b256};

    fn make_auth_receipt(authenticator_addr: Address, commitment: B256) -> Receipt {
        let topic0 = BATCH_INFO_AUTHENTICATED_TOPIC;
        let log = Log {
            address: authenticator_addr,
            data: LogData::new_unchecked(vec![topic0, commitment], Default::default()),
        };
        Receipt { status: Eip658Value::Eip658(true), logs: vec![log], ..Default::default() }
    }

    fn make_failed_auth_receipt(authenticator_addr: Address, commitment: B256) -> Receipt {
        let topic0 = BATCH_INFO_AUTHENTICATED_TOPIC;
        let log = Log {
            address: authenticator_addr,
            data: LogData::new_unchecked(vec![topic0, commitment], Default::default()),
        };
        Receipt { status: Eip658Value::Eip658(false), logs: vec![log], ..Default::default() }
    }

    fn test_legacy_tx(to: Address) -> TxEnvelope {
        let sig = Signature::test_signature();
        TxEnvelope::Legacy(Signed::new_unchecked(
            TxLegacy { to: TxKind::Call(to), ..Default::default() },
            sig,
            Default::default(),
        ))
    }

    #[test]
    fn test_compute_calldata_batch_hash() {
        let data = b"hello world";
        let hash = compute_calldata_batch_hash(data);
        assert_eq!(hash, keccak256(data));
    }

    #[test]
    fn test_compute_blob_batch_hash() {
        let h1 = b256!("0000000000000000000000000000000000000000000000000000000000000001");
        let h2 = b256!("0000000000000000000000000000000000000000000000000000000000000002");
        let hash = compute_blob_batch_hash(&[h1, h2]);

        let mut expected_input = Vec::new();
        expected_input.extend_from_slice(h1.as_slice());
        expected_input.extend_from_slice(h2.as_slice());
        assert_eq!(hash, keccak256(&expected_input));
    }

    #[test]
    fn test_collect_auth_events_from_receipts_success() {
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let commitment = b256!("abcdef0000000000000000000000000000000000000000000000000000000000");

        let receipt = make_auth_receipt(auth_addr, commitment);
        let result = collect_auth_events_from_receipts(&[receipt], auth_addr);

        assert!(result.contains(&commitment));
        assert_eq!(result.len(), 1);
    }

    #[test]
    fn test_collect_auth_events_from_receipts_wrong_address() {
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let wrong_addr = address!("0000000000000000000000000000000000000001");
        let commitment = b256!("abcdef0000000000000000000000000000000000000000000000000000000000");

        let receipt = make_auth_receipt(wrong_addr, commitment);
        let result = collect_auth_events_from_receipts(&[receipt], auth_addr);

        assert!(result.is_empty());
    }

    #[test]
    fn test_collect_auth_events_from_receipts_failed_receipt() {
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let commitment = b256!("abcdef0000000000000000000000000000000000000000000000000000000000");

        let receipt = make_failed_auth_receipt(auth_addr, commitment);
        let result = collect_auth_events_from_receipts(&[receipt], auth_addr);

        assert!(result.is_empty());
    }

    #[test]
    fn test_collect_auth_events_multiple_events() {
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let c1 = b256!("0000000000000000000000000000000000000000000000000000000000000001");
        let c2 = b256!("0000000000000000000000000000000000000000000000000000000000000002");

        let r1 = make_auth_receipt(auth_addr, c1);
        let r2 = make_auth_receipt(auth_addr, c2);
        let result = collect_auth_events_from_receipts(&[r1, r2], auth_addr);

        assert_eq!(result.len(), 2);
        assert!(result.contains(&c1));
        assert!(result.contains(&c2));
    }

    #[test]
    fn test_is_batch_authorized_post_fork_event_path() {
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let config = BatchAuthConfig { authenticator_address: auth_addr };
        let batch_hash = b256!("abcdef0000000000000000000000000000000000000000000000000000000000");
        let mut authenticated = BTreeSet::new();
        authenticated.insert(batch_hash);

        let tx = test_legacy_tx(Address::ZERO);
        // Post-fork, event present: authorized regardless of sender.
        assert!(is_batch_authorized(
            &tx,
            batch_hash,
            Some(&config),
            &authenticated,
            Address::ZERO,
            true,
        ));
    }

    #[test]
    fn test_is_batch_authorized_post_fork_no_event_rejected() {
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let config = BatchAuthConfig { authenticator_address: auth_addr };
        let batch_hash = b256!("abcdef0000000000000000000000000000000000000000000000000000000000");
        let authenticated = BTreeSet::new(); // empty

        let tx = test_legacy_tx(Address::ZERO);
        // Post-fork, no event: rejected even for an empty hash set.
        assert!(!is_batch_authorized(
            &tx,
            batch_hash,
            Some(&config),
            &authenticated,
            Address::ZERO,
            true,
        ));
    }

    #[test]
    fn test_is_batch_authorized_post_fork_sender_fallback_rejected() {
        // Even when sender matches batcher_address, post-fork requires an event.
        let batch_hash = B256::ZERO;
        let authenticated = BTreeSet::new();

        let tx = test_legacy_tx(Address::ZERO);
        let sender = tx.recover_signer().unwrap();
        // No auth_config => no event-based authorization possible => rejected.
        assert!(!is_batch_authorized(&tx, batch_hash, None, &authenticated, sender, true));

        // With an auth_config but no matching event: still rejected.
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let config = BatchAuthConfig { authenticator_address: auth_addr };
        assert!(
            !is_batch_authorized(&tx, batch_hash, Some(&config), &authenticated, sender, true,)
        );
    }

    #[test]
    fn test_is_batch_authorized_pre_fork_vanilla_sender_path() {
        let batch_hash = B256::ZERO;
        let authenticated = BTreeSet::new();

        let tx = test_legacy_tx(Address::ZERO);
        let sender = tx.recover_signer().unwrap();
        // Pre-fork: sender matches batcher_address => authorized.
        assert!(is_batch_authorized(&tx, batch_hash, None, &authenticated, sender, false));
        // Pre-fork: sender mismatch => rejected.
        assert!(!is_batch_authorized(
            &tx,
            batch_hash,
            None,
            &authenticated,
            address!("0000000000000000000000000000000000000001"),
            false,
        ));
    }

    #[test]
    fn test_is_batch_authorized_pre_fork_ignores_auth_event() {
        // Pre-fork, even if a matching auth event exists, only the sender check is honored.
        let auth_addr = address!("1234567890123456789012345678901234567890");
        let config = BatchAuthConfig { authenticator_address: auth_addr };
        let batch_hash = b256!("abcdef0000000000000000000000000000000000000000000000000000000000");
        let mut authenticated = BTreeSet::new();
        authenticated.insert(batch_hash);

        let tx = test_legacy_tx(Address::ZERO);
        let sender = tx.recover_signer().unwrap();
        let wrong_addr = address!("0000000000000000000000000000000000000001");

        // Sender mismatch: even with an event present, we reject (event path is gated off).
        assert!(!is_batch_authorized(
            &tx,
            batch_hash,
            Some(&config),
            &authenticated,
            wrong_addr,
            false,
        ));
        // Sender match: authorized.
        assert!(
            is_batch_authorized(&tx, batch_hash, Some(&config), &authenticated, sender, false,)
        );
    }

    #[test]
    fn test_batch_info_authenticated_topic_is_correct() {
        assert_eq!(BATCH_INFO_AUTHENTICATED_TOPIC, keccak256("BatchInfoAuthenticated(bytes32)"));
    }

    #[test]
    fn test_new_batch_auth_cache() {
        let lookback = 100u64;
        let cache = BatchAuthCache::new(lookback);
        let expected_cap = (lookback as usize) + 2;
        assert_eq!(cache.receipts.len(), 0);
        assert_eq!(cache.receipts.cap().get(), expected_cap);
        assert_eq!(cache.headers.len(), 0);
        assert_eq!(cache.headers.cap().get(), expected_cap);
    }
}
