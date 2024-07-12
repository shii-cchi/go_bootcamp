package ex02

import (
	"sync"
)

func multiplex(inputChans ...chan interface{}) chan interface{} {
	outChan := make(chan interface{})

	var wg sync.WaitGroup

	go func() {
		defer close(outChan)

		for _, inputChan := range inputChans {
			wg.Add(1)

			go func(inputChan chan interface{}) {
				defer wg.Done()

				for value := range inputChan {
					outChan <- value
				}
			}(inputChan)
		}

		wg.Wait()
	}()

	return outChan
}
