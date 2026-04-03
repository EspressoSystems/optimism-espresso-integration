#!/usr/bin/env python3
"""
Decode a BatchInbox calldata transaction from the Espresso-integrated OP stack.

Usage:
    python3 decode_batch_calldata.py <calldata_hex>
    python3 decode_batch_calldata.py 0x0050546c...

Output:
  - Parsed frame header (channel_id, frame_number, frame_data_length, is_last)
  - keccak256 commitment (to match against authenticateBatchInfo on BatchAuthenticator)
  - Decompressed size
  - RLP-decoded batch attributes (l2 block numbers, timestamps, epoch info)

Requirements: pip install rlp eth-hash[pycryptodome]
  or just:     pip install pysha3   (for keccak256 without the full eth stack)
"""

import sys
import zlib
import struct
import json
import hashlib

# ---------------------------------------------------------------------------
# keccak256 — try eth_hash, fall back to pysha3, fall back to a clear error
# ---------------------------------------------------------------------------
try:
    from Crypto.Hash import keccak as _keccak
    def keccak256(data: bytes) -> str:
        k = _keccak.new(digest_bits=256)
        k.update(data)
        return k.hexdigest()
except ImportError:
    try:
        import sha3 as _sha3  # pysha3 patches hashlib with keccak_256
        _ = _sha3  # mark used
        def keccak256(data: bytes) -> str:
            return hashlib.new("keccak_256", data).hexdigest()
    except Exception:
        def keccak256(_data: bytes) -> str:
            return "(install pycryptodome or pysha3 for keccak256)"

# ---------------------------------------------------------------------------
# RLP decoder (minimal, no external deps)
# ---------------------------------------------------------------------------
def rlp_decode(data: bytes):
    """Minimal RLP decoder, returns Python objects (bytes, lists)."""
    def _decode(data, pos):
        b = data[pos]
        if b < 0x80:
            return bytes([b]), pos + 1
        elif b < 0xB8:
            length = b - 0x80
            return data[pos+1:pos+1+length], pos+1+length
        elif b < 0xC0:
            len_of_len = b - 0xB7
            length = int.from_bytes(data[pos+1:pos+1+len_of_len], "big")
            start = pos + 1 + len_of_len
            return data[start:start+length], start+length
        elif b < 0xF8:
            length = b - 0xC0
            end = pos + 1 + length
            items, cur = [], pos + 1
            while cur < end:
                item, cur = _decode(data, cur)
                items.append(item)
            return items, end
        else:
            len_of_len = b - 0xF7
            length = int.from_bytes(data[pos+1:pos+1+len_of_len], "big")
            start = pos + 1 + len_of_len
            end = start + length
            items, cur = [], start
            while cur < end:
                item, cur = _decode(data, cur)
                items.append(item)
            return items, end

    result, _ = _decode(data, 0)
    return result


# ---------------------------------------------------------------------------
# Batch attribute decoder (SingularBatch / SpanBatch detection)
# ---------------------------------------------------------------------------
BATCH_TYPE_SINGULAR = 0
BATCH_TYPE_SPAN     = 1

def int_from_bytes(b):
    return int.from_bytes(b, "big") if isinstance(b, bytes) and len(b) > 0 else 0


def decode_singular_batch(payload: bytes):
    """Decode a SingularBatch RLP payload: [parent_hash, epoch_num, epoch_hash, timestamp, txs]"""
    decoded = rlp_decode(payload)
    if not isinstance(decoded, list) or len(decoded) < 5:
        return {"decode_error": f"expected RLP list of 5+, got {type(decoded).__name__}"}

    tx_count = len(decoded[4]) if isinstance(decoded[4], list) else "?"
    return {
        "type":        "SingularBatch",
        "parent_hash": decoded[0].hex() if isinstance(decoded[0], bytes) else "?",
        "epoch_num":   int_from_bytes(decoded[1]),
        "epoch_hash":  decoded[2].hex() if isinstance(decoded[2], bytes) else "?",
        "timestamp":   int_from_bytes(decoded[3]),
        "tx_count":    tx_count,
    }


def decode_batches(decompressed: bytes):
    """
    Decode all batches from decompressed channel data.

    Each batch item is one of:
      - type_byte=0x00 followed by RLP(SingularBatch fields)
      - type_byte=0x01 followed by SpanBatch encoding
      - plain RLP list (legacy, no type prefix — used in some OP versions)
    """
    results = []
    pos = 0
    idx = 0

    while pos < len(decompressed):
        first = decompressed[pos]

        # OP spec: type byte 0x00 or 0x01 followed by batch payload
        if first in (BATCH_TYPE_SINGULAR, BATCH_TYPE_SPAN):
            batch_type = first
            payload = decompressed[pos+1:]
            if batch_type == BATCH_TYPE_SINGULAR:
                try:
                    info = decode_singular_batch(payload)
                    # advance past the full RLP item to find the next batch
                    _, end = _rlp_decode_item(decompressed, pos + 1)
                    results.append(info)
                    pos = end
                except Exception as e:
                    results.append({"type": "SingularBatch", "decode_error": str(e),
                                    "raw_hex": decompressed[pos:pos+32].hex()})
                    break
            else:
                results.append({"type": "SpanBatch",
                                "note": "SpanBatch full decode not implemented",
                                "raw_hex": decompressed[pos:pos+32].hex()})
                break

        # RLP-wrapped batch (no type prefix byte, or starts with RLP prefix like 0xb8/0xf8)
        elif first >= 0x80:
            try:
                item, end = _rlp_decode_item(decompressed, pos)
                # item is either bytes (wrapping inner batch data) or a list
                if isinstance(item, bytes) and len(item) > 0:
                    inner = item
                    # Inner may start with a type byte
                    if inner[0] in (BATCH_TYPE_SINGULAR, BATCH_TYPE_SPAN):
                        batch_type = inner[0]
                        payload = inner[1:]
                        if batch_type == BATCH_TYPE_SINGULAR:
                            try:
                                info = decode_singular_batch(payload)
                                results.append(info)
                            except Exception as e:
                                results.append({"type": "SingularBatch(inner)",
                                                "decode_error": str(e),
                                                "raw_hex": inner[:32].hex()})
                        else:
                            results.append({"type": "SpanBatch(inner)",
                                            "note": "SpanBatch full decode not implemented",
                                            "raw_hex": inner[:32].hex()})
                    else:
                        results.append({"type": "RLPBytes",
                                        "note": "inner data doesn't start with known batch type",
                                        "first_byte": f"0x{inner[0]:02x}",
                                        "raw_hex": inner[:32].hex()})
                elif isinstance(item, list):
                    results.append({"type": "RLPList",
                                    "note": "top-level RLP list (may be legacy SingularBatch)",
                                    "items": len(item)})
                pos = end
            except Exception as e:
                results.append({"decode_error": str(e),
                                "raw_hex": decompressed[pos:pos+32].hex()})
                break
        else:
            results.append({"note": f"unexpected first byte 0x{first:02x} at offset {pos}",
                            "raw_hex": decompressed[pos:pos+32].hex()})
            break

        idx += 1
        if idx > 1000:
            results.append({"note": "stopped after 1000 batches"})
            break

    return results


def _rlp_decode_item(data: bytes, pos: int):
    """Return (decoded_item, new_pos) for a single RLP item."""
    b = data[pos]
    if b < 0x80:
        return bytes([b]), pos + 1
    elif b < 0xB8:
        length = b - 0x80
        return data[pos+1:pos+1+length], pos+1+length
    elif b < 0xC0:
        len_of_len = b - 0xB7
        length = int.from_bytes(data[pos+1:pos+1+len_of_len], "big")
        start = pos + 1 + len_of_len
        return data[start:start+length], start+length
    elif b < 0xF8:
        length = b - 0xC0
        end = pos + 1 + length
        items, cur = [], pos + 1
        while cur < end:
            item, cur = _rlp_decode_item(data, cur)
            items.append(item)
        return items, end
    else:
        len_of_len = b - 0xF7
        length = int.from_bytes(data[pos+1:pos+1+len_of_len], "big")
        start = pos + 1 + len_of_len
        end = start + length
        items, cur = [], start
        while cur < end:
            item, cur = _rlp_decode_item(data, cur)
            items.append(item)
        return items, end


# ---------------------------------------------------------------------------
# AltDA commitment decoder (DerivationVersion1 = 0x01)
# ---------------------------------------------------------------------------
# Wire format: 0x01 | commitment_type (1 byte) | commitment_data
#   commitment_type 0x00 = Keccak256Commitment — 32-byte keccak256 of the frame data
#   commitment_type 0x01 = GenericCommitment   — opaque DA provider bytes (e.g. EigenDA cert)
#
# The BatchAuthenticator.validBatchInfo mapping is keyed on keccak256(full_calldata),
# i.e. keccak256 of the entire tx data field including the leading 0x01 byte.
# This is the "commitment" arg passed to authenticateBatchInfo().
# ---------------------------------------------------------------------------

def decode_altda(calldata: bytes, commitment: str):
    print("=== AltDA Commitment (DerivationVersion1) ===")
    print("  This transaction posts a DA commitment pointer, not raw frame data.")
    print()

    if len(calldata) < 2:
        print("Error: altDA calldata too short (missing commitment type byte)", file=sys.stderr)
        sys.exit(1)

    commitment_type = calldata[1]
    commitment_data = calldata[2:]

    COMMITMENT_TYPE_KECCAK256 = 0
    COMMITMENT_TYPE_GENERIC   = 1

    if commitment_type == COMMITMENT_TYPE_KECCAK256:
        type_name = "Keccak256Commitment"
        if len(commitment_data) != 32:
            print(f"  Warning: Keccak256Commitment should be 32 bytes, got {len(commitment_data)}")
        print(f"  commitment_type : 0x00 ({type_name})")
        print(f"  da_commitment   : 0x{commitment_data.hex()}")
        print()
        print("  The DA server holds the raw frame data whose keccak256 equals da_commitment.")
        print("  Fetch it with:")
        print(f"    curl <DA_SERVER_URL>/get/0x{commitment_data.hex()}")
        print()
        print("  To verify integrity, keccak256(fetched_data) must equal:")
        print(f"    0x{commitment_data.hex()}")
    elif commitment_type == COMMITMENT_TYPE_GENERIC:
        type_name = "GenericCommitment (EigenDA)"
        print(f"  commitment_type : 0x01 ({type_name})")
        print(f"  da_commitment   : 0x{commitment_data.hex()}")
        print()
        print("  This is an opaque DA provider commitment (e.g. EigenDA cert/blob ref).")
        print("  Fetch from the EigenDA proxy with:")
        print(f"    curl <EIGENDA_PROXY_URL>/get/0x01{commitment_data.hex()}")
        print()
        print("  Note: the proxy prepends the commitment_type byte (0x01) when fetching.")
    else:
        print(f"  commitment_type : 0x{commitment_type:02x} (unknown)")
        print(f"  da_commitment   : 0x{commitment_data.hex()}")
        print()
        print("  Unknown commitment type — cannot decode further.")

    print()
    print("=== BatchAuthenticator Verification ===")
    print("  The commitment stored in BatchAuthenticator.validBatchInfo is:")
    print(f"    0x{commitment}")
    print("  This equals keccak256(full_calldata) — the entire tx data field.")
    print()
    print("  Check on-chain with cast:")
    print(f"    cast call <BATCH_AUTHENTICATOR_ADDR> 'validBatchInfo(bytes32)(bool)' 0x{commitment}")
    print()
    print("  If it returns true, authenticateBatchInfo() was called for this batch")
    print("  before it was posted on-chain.")
    print()
    print("Done.")


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------
def main():
    import argparse
    parser = argparse.ArgumentParser(description="Decode BatchInbox calldata from the Espresso-integrated OP stack")
    parser.add_argument("calldata", help="Calldata hex string (with or without 0x prefix)")
    parser.add_argument("-v", "--verbose", action="store_true", help="Print all decoded batches (default: summary only)")
    args_parsed = parser.parse_args()

    raw_hex = args_parsed.calldata.strip()
    if raw_hex.startswith("0x") or raw_hex.startswith("0X"):
        raw_hex = raw_hex[2:]

    try:
        calldata = bytes.fromhex(raw_hex)
    except ValueError as e:
        print(f"Error: invalid hex input: {e}", file=sys.stderr)
        sys.exit(1)

    total_size = len(calldata)
    print(f"Total calldata size: {total_size} bytes ({total_size/1024:.1f} KB)")
    print()

    # --- Commitment (for linking to authenticateBatchInfo) ---
    commitment = keccak256(calldata)
    print("=== Commitment (keccak256 of full calldata) ===")
    print(f"  {commitment}")
    print("  Match this against the 'commitment' arg in the authenticateBatchInfo tx")
    print()

    # --- Version dispatch ---
    if total_size < 1:
        print("Error: calldata is empty", file=sys.stderr)
        sys.exit(1)

    version = calldata[0]

    # DerivationVersion1 (0x01) = altDA commitment pointer, not raw frames
    if version == 1:
        decode_altda(calldata, commitment)
        return

    # --- Frame header (DerivationVersion0 = 0x00) ---
    if version != 0:
        print(f"Warning: unexpected version byte 0x{version:02x} (expected 0x00 or 0x01)")

    if total_size < 24:
        print("Error: calldata too short to contain a valid frame header", file=sys.stderr)
        sys.exit(1)

    channel_id         = calldata[1:17]
    frame_number       = struct.unpack_from(">H", calldata, 17)[0]   # uint16 big-endian
    frame_data_length  = struct.unpack_from(">I", calldata, 19)[0]   # uint32 big-endian

    frame_data_start = 23
    frame_data_end   = frame_data_start + frame_data_length

    if frame_data_end + 1 > total_size:
        print(f"Error: calldata too short: need {frame_data_end+1} bytes, have {total_size}", file=sys.stderr)
        sys.exit(1)

    frame_data = calldata[frame_data_start:frame_data_end]
    is_last    = calldata[frame_data_end]

    print("=== Frame Header ===")
    print(f"  version          : {version}")
    print(f"  channel_id       : 0x{channel_id.hex()}")
    print(f"  frame_number     : {frame_number}")
    print(f"  frame_data_length: {frame_data_length} bytes")
    print(f"  is_last          : {bool(is_last)} (0x{is_last:02x})")
    print()

    # --- Decompress ---
    print("=== Decompression ===")
    if frame_data[:2] in (b'\x78\x9c', b'\x78\xda', b'\x78\x01', b'\x78\x5e'):
        algo = "zlib"
        try:
            decompressed = zlib.decompress(frame_data)
        except zlib.error as e:
            print(f"  zlib decompress failed: {e}")
            sys.exit(1)
    elif frame_data[:3] == b'\xff\x06\x00':
        algo = "brotli (not supported — install brotli package)"
        print(f"  {algo}")
        sys.exit(0)
    else:
        algo = f"unknown (magic: {frame_data[:4].hex()})"
        print(f"  Unrecognised compression magic: {algo}")
        sys.exit(1)

    print(f"  algorithm        : {algo}")
    print(f"  compressed size  : {frame_data_length} bytes")
    print(f"  decompressed size: {len(decompressed)} bytes")
    print(f"  ratio            : {frame_data_length/len(decompressed):.3f}")
    print()

    # --- Decode batches ---
    print("=== Batches in channel ===")
    print(f"  decompressed first 32 bytes: {decompressed[:32].hex()}")
    print()
    batches = decode_batches(decompressed)

    if args_parsed.verbose:
        for i, b in enumerate(batches):
            print(f"  [{i}] {json.dumps(b, indent=6)}")
    else:
        # Summary: first batch, last batch, count, and ranges
        singular = [b for b in batches if b.get("type") == "SingularBatch" and "decode_error" not in b]
        errors   = [b for b in batches if "decode_error" in b]
        other    = [b for b in batches if b not in singular and b not in errors]

        print(f"  Total batches decoded : {len(batches)}")
        if singular:
            timestamps = [b["timestamp"] for b in singular]
            epochs     = [b["epoch_num"]  for b in singular]
            tx_counts  = [b["tx_count"]   for b in singular if isinstance(b.get("tx_count"), int)]
            print(f"  SingularBatch count  : {len(singular)}")
            print(f"  epoch_num range      : {min(epochs)} → {max(epochs)}")
            print(f"  timestamp range      : {min(timestamps)} → {max(timestamps)}")
            if tx_counts:
                print(f"  tx_count range       : {min(tx_counts)} → {max(tx_counts)}  (total txs: {sum(tx_counts)})")
            print()
            print("  First batch:")
            print(f"    {json.dumps(singular[0], indent=4)}")
            if len(singular) > 1:
                print("  Last batch:")
                print(f"    {json.dumps(singular[-1], indent=4)}")
        if other:
            print(f"  Other batch types    : {len(other)}")
            for b in other:
                print(f"    {json.dumps(b)}")
        if errors:
            print(f"  Decode errors        : {len(errors)}")
            for b in errors:
                print(f"    {json.dumps(b)}")
        print()
        print("  (use -v / --verbose to print all batches)")

    print()
    print("Done.")


if __name__ == "__main__":
    main()
