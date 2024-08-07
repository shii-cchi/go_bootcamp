package ex00

import (
	"sync"
	"time"
)

func sleepSort(unsortedNumbers []int) chan int {
	numChan := make(chan int, len(unsortedNumbers))
	defer close(numChan)

	var wg sync.WaitGroup

	for _, number := range unsortedNumbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(time.Duration(num) * time.Second)
			numChan <- num
		}(number)
	}

	wg.Wait()

	return numChan
}
