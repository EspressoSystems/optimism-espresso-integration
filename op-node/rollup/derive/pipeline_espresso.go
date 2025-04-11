// go:build caffnode
package derive

import (
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/log"
)

// NewDerivationPipeline creates a DerivationPipeline, to turn L1 data into L2 block-inputs.
func NewEspressoDerivationPipeline(log log.Logger, rollupCfg *rollup.Config, l1Fetcher L1Fetcher, l1Blobs L1BlobsFetcher,
	altDA AltDAInputFetcher, l2Source L2Source, metrics Metrics, managedMode bool,
) *DerivationPipeline {
	spec := rollup.NewChainSpec(rollupCfg)
	// Stages are strung together into a pipeline,
	// results are pulled from the stage closed to the L2 engine, which pulls from the previous stage, and so on.
	var l1Traversal l1TraversalStage
	if managedMode {
		l1Traversal = NewL1TraversalManaged(log, rollupCfg, l1Fetcher)
	} else {
		l1Traversal = NewL1Traversal(log, rollupCfg, l1Fetcher)
	}
	dataSrc := NewDataSourceFactory(log, rollupCfg, l1Fetcher, l1Blobs, altDA) // auxiliary stage for L1Retrieval
	l1Src := NewL1Retrieval(log, dataSrc, l1Traversal)
	frameQueue := NewFrameQueue(log, rollupCfg, l1Src)
	channelMux := NewChannelMux(log, spec, frameQueue, metrics)
	chInReader := NewChannelInReader(rollupCfg, log, channelMux, metrics)
	batchMux := NewBatchMux(log, rollupCfg, chInReader, l2Source)
	attrBuilder := NewFetchingAttributesBuilder(rollupCfg, l1Fetcher, l2Source)
	attributesQueue := NewAttributesQueue(log, rollupCfg, attrBuilder, batchMux)

	// Reset from ResetEngine then up from L1 Traversal. The stages do not talk to each other during
	// the ResetEngine, but after the ResetEngine, this is the order in which the stages could talk to each other.
	// Note: The ResetEngine is the only reset that can fail.
	stages := []ResettableStage{l1Traversal, l1Src, altDA, frameQueue, channelMux, chInReader, batchMux, attributesQueue}

	return &DerivationPipeline{
		log:       log,
		rollupCfg: rollupCfg,
		l1Fetcher: l1Fetcher,
		altDA:     altDA,
		resetting: 0,
		stages:    stages,
		metrics:   metrics,
		traversal: l1Traversal,
		attrib:    attributesQueue,
		l2:        l2Source,
	}
}
