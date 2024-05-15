package ex00

import (
	"sync"
	"time"
)

func sleepSort(unsortedNumbers []int) chan int {
	numChan := make(chan int)

	var wg sync.WaitGroup

	go func() {
		defer close(numChan)

		for _, number := range unsortedNumbers {
			wg.Add(1)
			go func(num int) {
				defer wg.Done()
				time.Sleep(time.Duration(num) * time.Second)
				numChan <- num
			}(number)
		}

		wg.Wait()
	}()

	return numChan
}
