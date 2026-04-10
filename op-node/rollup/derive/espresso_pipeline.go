//go:build !mips64

package derive

import op "github.com/EspressoSystems/espresso-streamers/op"

// EspressoStreamer returns the underlying Espresso batch streamer.
// Used by op-node to expose the streamer for metrics and service-level access.
func (dp *DerivationPipeline) EspressoStreamer() *op.BatchStreamer[EspressoBatch] {
	return dp.attrib.espresso.espressoStreamer
}
