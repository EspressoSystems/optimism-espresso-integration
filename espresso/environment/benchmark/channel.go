package benchmark

// TeeChan is a helper function that takes a channel that is expected to be
// send to, (the source channel) and returns two channels that should be
// submitted to.
func TeeChan[T any](src <-chan T) (<-chan T, <-chan T) {
	dst1 := make(chan T, cap(src))
	dst2 := make(chan T, cap(src))

	go func() {
		for v := range src {
			dst1 <- v
			dst2 <- v
		}

		close(dst1)
		close(dst2)
	}()

	return dst1, dst2
}
