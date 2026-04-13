// espresso-spy is a lightweight tool that polls HotShot and logs every
// EspressoBatch it sees for a given namespace, without touching the channel
// manager or L1.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	tagged_base64 "github.com/EspressoSystems/espresso-network/sdks/go/tagged-base64"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	url := flag.String("url", "http://localhost:24000", "Espresso query service URL")
	namespace := flag.Uint64("namespace", 0, "Espresso namespace (L2 chain ID)")
	start := flag.Uint64("start", 0, "HotShot block height to start from")
	batchSize := flag.Uint64("batch", 100, "Number of HotShot blocks to fetch per request")
	pollInterval := flag.Duration("poll", 500*time.Millisecond, "Poll interval when caught up")
	flag.Parse()

	if *namespace == 0 {
		fmt.Fprintln(os.Stderr, "error: --namespace is required")
		os.Exit(1)
	}

	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))

	client := espressoClient.NewClient(*url)

	ctx := context.Background()
	pos := *start
	logger.Info("espresso-spy starting", "url", *url, "namespace", *namespace, "startHeight", pos)

	for {
		latest, err := client.FetchLatestBlockHeight(ctx)
		// logger.Info("starting", "latest", latest)
		if err != nil {
			logger.Warn("failed to fetch latest block height", "err", err)
			time.Sleep(*pollInterval)
			continue
		}

		if pos >= latest {
			time.Sleep(*pollInterval)
			continue
		}

		end := pos + *batchSize
		if end > latest+1 {
			end = latest + 1
		}

		logger.Info("starting", "s", pos, "end", end)

		blocks, err := client.FetchNamespaceTransactionsInRange(ctx, pos, end, *namespace)
		if err != nil {
			logger.Warn("failed to fetch namespace transactions", "start", pos, "end", end, "err", err)
			time.Sleep(*pollInterval)
			continue
		}

		for i, block := range blocks {
			hotShotHeight := pos + uint64(i)
			if len(block.Transactions) == 0 {
				continue
			}
			for j, txn := range block.Transactions {
				batch, err := derive.UnmarshalEspressoTransaction(txn.Payload)
				if err != nil {
					logger.Warn("failed to unmarshal batch",
						"hotShotHeight", hotShotHeight,
						"txIndex", j,
						"err", err,
					)
					continue
				}
				// if batch.Number() >= 20_000 { // && batch.Number() <= 7258 {
				commitment := txn.Commit()
				txHash, _ := tagged_base64.New("TX", commitment[:])
				// block := batch.ToBlock()
				logger.Info("batch seen",
					"hotShotHeight", hotShotHeight,
					"txIndex", j,
					"blockNr", batch.Number(),
					"hash", batch.Hash(),
					"headerHash", batch.Header().Hash().Hex(),
					"parentHash", batch.Header().ParentHash.Hex(),
					"timestamp", batch.Batch.Timestamp,
					"l1OriginNum", batch.Batch.EpochNum,
					"l1OriginHash", batch.Batch.EpochHash.Hex(),
					"hotshotTxHash", txHash,
					"numTxs", len(batch.Batch.Transactions),
				)
			}
		}

		pos = end
	}
}

func ToBlock(b *derive.EspressoBatch) (*types.Block, error) {
	// Re-insert the deposit transaction
	txs := []*types.Transaction{b.L1InfoDeposit}
	for i, opaqueTx := range b.Batch.Transactions {
		var tx types.Transaction
		err := tx.UnmarshalBinary(opaqueTx)
		if err != nil {
			return nil, fmt.Errorf("could not decode tx %d: %w", i, err)
		}
		txs = append(txs, &tx)
	}
	return types.NewBlockWithHeader(b.BatchHeader).WithBody(types.Body{
		Transactions: txs,
	}), nil
}
