package batcher

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	espressoLightClient "github.com/EspressoSystems/espresso-network/sdks/go/light-client"
	op "github.com/EspressoSystems/espresso-streamers/op"
	"github.com/EspressoSystems/espresso-streamers/op/derivation"
	opcrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// EspressoDriverSetup groups all Espresso-specific runtime state plumbed from
// BatcherService into DriverSetup. Defined here to keep the upstream Optimism
// DriverSetup field block compact (see driver.go).
//
// All fields are nil/zero when --espresso.enabled is false except for the
// fallback batcher's ChainSigner/SequencerAddress, which are always populated
// by applyEspressoDriverSetup.
type EspressoDriverSetup struct {
	Client           *espressoClient.MultipleNodesClient
	LightClient      *espressoLightClient.LightclientCaller
	ChainSigner      opcrypto.ChainSigner
	SequencerAddress common.Address
	Attestation      []byte
}

// batcherL1Adapter wraps the batcher's L1Client to implement espresso.L1Client
// (HeaderHashByNumber + bind.ContractCaller).
type batcherL1Adapter struct {
	L1Client L1Client
}

func (a *batcherL1Adapter) HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error) {
	h, err := a.L1Client.HeaderByNumber(ctx, number)
	if err != nil {
		return common.Hash{}, err
	}
	return h.Hash(), nil
}

func (a *batcherL1Adapter) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return a.L1Client.CodeAt(ctx, contract, blockNumber)
}

func (a *batcherL1Adapter) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return a.L1Client.CallContract(ctx, call, blockNumber)
}

// batcherL2Adapter wraps the batcher's L2 eth client to implement op.L2Client.
// The streamer uses it once, at construction, to resolve the block hash of the
// batch position it is anchored to. dial.EthClientInterface exposes no
// header-only accessor, so this fetches the full block and takes its hash.
type batcherL2Adapter struct {
	EthClient dial.EthClientInterface
}

func (a *batcherL2Adapter) HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error) {
	block, err := a.EthClient.BlockByNumber(ctx, number)
	if err != nil {
		return common.Hash{}, err
	}
	return block.Hash(), nil
}

// setupEspressoStreamer constructs the Espresso streamer for a BatchSubmitter that
// is starting up; no-op when --espresso.enabled is false.
//
// Called from StartBatchSubmitting rather than NewBatchSubmitter: the streamer
// resolves its anchor batch's hash from the L2 client while constructing, so it
// needs a context and an L2 node that has reached genesis. It also returns an error
// rather than panicking, which construction inside NewBatchSubmitter could not do.
func (l *BatchSubmitter) setupEspressoStreamer(ctx context.Context) error {
	if !l.Config.Espresso.Enabled {
		return nil
	}

	ethClient, err := l.EndpointProvider.EthClient(ctx)
	if err != nil {
		return fmt.Errorf("getting the L2 eth client for the Espresso streamer: %w", err)
	}

	// Convert typed nil pointer to untyped nil interface to avoid typed-nil interface panic
	// in confirmEspressoBlockHeight when EspressoLightClient is not configured.
	var lightClientIface op.LightClientCallerInterface
	if l.Espresso.LightClient != nil {
		lightClientIface = l.Espresso.LightClient
	}

	streamer, err := op.NewStreamer(
		ctx,
		l.Espresso.Client,
		&batcherL1Adapter{L1Client: l.L1Client},
		&batcherL2Adapter{EthClient: ethClient},
		lightClientIface,
		l.RollupConfig.BatchAuthenticatorAddress,
		l.RollupConfig.L2ChainID.Uint64(),
		derivation.CreateEspressoBatchUnmarshaler(),
		l.getSyncStatus,
		l.Config.Espresso.PollInterval,
		l.Log,
		l.Config.Espresso.CaffeinationHeightEspresso,
		l.Config.Espresso.CaffeinationHeightL2,
	)
	if err != nil {
		return fmt.Errorf("failed to create Espresso streamer: %w", err)
	}
	l.espressoStreamer = streamer
	return nil
}

// startEspressoLoops registers the batcher with the BatchAuthenticator
// contract, resolves the TEE verifier address, spawns the Espresso transaction
// submitter, and starts the four Espresso-specific batcher goroutines (in
// addition to the upstream receiptsLoop and publishingLoop). Replaces the
// upstream three-goroutine pattern when --espresso.enabled is set.
func (l *BatchSubmitter) startEspressoLoops(receiptsCh chan txmgr.TxReceipt[txRef], publishSignal chan pubInfo) error {
	if err := l.registerBatcher(l.killCtx); err != nil {
		return fmt.Errorf("could not register with BatchAuthenticator contract: %w", err)
	}

	// The streamer drives itself from its own poll loop, so it is started here rather
	// than being pumped by espressoBatchLoadingLoop. Bound to shutdownCtx so it stops
	// fetching before the publish path winds down.
	if err := l.espressoStreamer.Start(l.shutdownCtx); err != nil {
		return fmt.Errorf("could not start the Espresso streamer: %w", err)
	}

	// Resolve the TEE verifier address from the BatchAuthenticator contract.
	if err := l.resolveTEEVerifierAddress(); err != nil {
		return fmt.Errorf("could not resolve TEE verifier address: %w", err)
	}

	l.espressoSubmitter = NewEspressoTransactionSubmitter(
		WithContext(l.shutdownCtx),
		WithWaitGroup(l.wg),
		WithEspressoClient(l.Espresso.Client),
		WithVerifyReceiptMaxBlocks(l.Config.Espresso.VerifyReceiptMaxBlocks),
		WithVerifyReceiptSafetyTimeout(l.Config.Espresso.VerifyReceiptSafetyTimeout),
		WithVerifyReceiptRetryDelay(l.Config.Espresso.VerifyReceiptRetryDelay),
	)
	l.espressoSubmitter.SpawnWorkers(4, 4)
	l.espressoSubmitter.Start()

	// Limit teeAuthGroup to at most 128 concurrent goroutines as an arbitrary
	// not-too-big limit for the number of BatchInbox transactions that can be
	// simultaneously waiting for corresponding BatchAuthenticator transaction to be
	// confirmed before submission to L1.
	l.teeAuthGroup.SetLimit(128)

	l.wg.Add(4)
	go l.receiptsLoop(l.wg, receiptsCh) // ranges over receiptsCh channel
	go l.espressoBatchQueueingLoop(l.shutdownCtx, l.wg)
	go l.espressoBatchLoadingLoop(l.shutdownCtx, l.wg, publishSignal)
	go l.publishingLoop(l.killCtx, l.wg, receiptsCh, publishSignal) // ranges over publishSignal, spawns routines which send on receiptsCh. Closes receiptsCh when done.
	return nil
}

// waitForTEEAuthGroup blocks until all in-flight TEE authentication goroutines
// (kicked off by sendTx for the Espresso TEE batcher and by the fallback
// batcher post-EspressoEnforcement) have completed. Called from
// publishingLoop's tail; blocks until killCtx is cancelled if any auth retries
// are still in flight.
func (l *BatchSubmitter) waitForTEEAuthGroup() {
	if err := l.teeAuthGroup.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			l.Log.Error("error waiting for transaction authentication requests to complete", "err", err)
		}
	}
}

// shouldSkipPublishForActiveSeq returns true if publishStateToL1 should skip
// publishing because this batcher is not the on-chain "active" batcher. The
// Espresso TEE batcher always honors the on-chain activeIsEspresso flag (it is
// fundamentally a post-fork actor); the fallback batcher honors the flag only
// post-EspressoEnforcement (pre-fork it must run as a vanilla upstream
// Optimism batcher with no BatchAuthenticator coupling).
func (l *BatchSubmitter) shouldSkipPublishForActiveSeq(ctx context.Context) bool {
	if !l.hasBatchAuthenticator() {
		return false
	}
	consultActiveFlag := l.Config.Espresso.Enabled
	if !consultActiveFlag {
		fallbackAuthRequired, err := l.isFallbackAuthRequired(ctx)
		if err != nil {
			l.Log.Warn("Failed to evaluate fallback-auth gate, skipping publish", "err", err)
			return true
		}
		consultActiveFlag = fallbackAuthRequired
	}
	if !consultActiveFlag {
		return false
	}
	isActive, err := l.isBatcherActive(ctx)
	if err != nil {
		l.Log.Warn("Failed to check if batcher is active, skipping publish", "err", err)
		return true
	}
	return !isActive
}

// resetEspressoStreamer re-anchors the Espresso streamer to the safe L2 head when
// --espresso.enabled is set; no-op otherwise. Called from clearState alongside the
// upstream channel-manager reset so the streamer's view of "next batch" matches the
// freshly-cleared channel state.
//
// The streamer is nil until StartBatchSubmitting constructs it, and clearState runs
// before that on the startup path, so the nil check is load-bearing.
func (l *BatchSubmitter) resetEspressoStreamer(ctx context.Context) {
	if !l.Config.Espresso.Enabled || l.espressoStreamer == nil {
		return
	}
	syncStatus, err := l.getSyncStatus(ctx)
	if err != nil {
		l.Log.Warn("Failed to fetch sync status to re-anchor the Espresso streamer, keeping the current tip", "err", err)
		return
	}
	l.espressoStreamer.ResetToSafeBatch(syncStatus)
}

// dispatchAuthenticatedSendTx routes sendTx through the Espresso (TEE) auth
// path or the fallback-batcher post-fork auth path, returning true when the tx
// has been handed off to teeAuthGroup. Returns false to mean "fall through to
// the upstream queue.Send path" — pre-fork fallback batcher and any cancel tx.
//
// The TEE batcher (Config.Espresso.Enabled == true) always authenticates. The
// fallback batcher consults isFallbackAuthRequired to gate authentication
// behind the EspressoEnforcement hardfork: pre-fork the verifier accepts plain
// sender-authenticated batches, and the BatchAuthenticator contract is
// irrelevant; calling authenticateBatchInfo pre-fork would also revert against
// the default activeIsEspresso=true contract state.
func (l *BatchSubmitter) dispatchAuthenticatedSendTx(txdata txData, isCancel bool, candidate *txmgr.TxCandidate, queue TxSender[txRef], receiptsCh chan txmgr.TxReceipt[txRef]) bool {
	if isCancel {
		return false
	}
	// Espresso batcher: authenticate via BatchAuthenticator.
	if l.Config.Espresso.Enabled {
		l.teeAuthGroup.Go(
			func() error {
				l.sendTxWithEspresso(txdata, isCancel, candidate, queue, receiptsCh)
				return nil
			},
		)
		return true
	}
	// Fallback batcher: authenticate via BatchAuthenticator only after the
	// EspressoEnforcement hardfork has activated. Pre-fork, the chain runs
	// pure upstream Optimism semantics — the verifier accepts plain
	// sender-authenticated batches, and the BatchAuthenticator contract is
	// irrelevant. Calling authenticateBatchInfo pre-fork would also revert
	// against the default activeIsEspresso=true contract state.
	if !l.hasBatchAuthenticator() {
		return false
	}
	fallbackAuthRequired, err := l.isFallbackAuthRequired(l.killCtx)
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  txRef{id: txdata.ID(), isCancel: isCancel, isBlob: txdata.daType == DaTypeBlob, daType: txdata.daType, size: txdata.Len()},
			Err: fmt.Errorf("failed to evaluate fallback-auth gate: %w", err),
		}
		return true
	}
	if fallbackAuthRequired {
		l.teeAuthGroup.Go(
			func() error {
				l.sendTxWithFallbackAuth(txdata, isCancel, candidate, queue, receiptsCh)
				return nil
			},
		)
		return true
	}
	return false
}
