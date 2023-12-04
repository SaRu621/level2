package pkg

import "sync"

func Or(doneChan ...<-chan interface{}) <-chan interface{} {
	merged := make(chan interface{})

	var wg sync.WaitGroup

	for _, ch := range doneChan {
		wg.Add(1)

		go func(ch <-chan interface{}) {
			defer wg.Done()

			for v := range ch {
				merged <- v
			}
			merged <- nil
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
