package ex02

import "sync"

func multiplex(inputChans ...<-chan interface{}) chan interface{} {
	outChan := make(chan interface{})

	var wg sync.WaitGroup

	for _, inputChan := range inputChans {
		wg.Add(1)

		go func(inputChan <-chan interface{}) {
			defer wg.Done()
			for {
				val, ok := <-inputChan

				if !ok {
					break
				}

				outChan <- val
			}
		}(inputChan)
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	return outChan
}
