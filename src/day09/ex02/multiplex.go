package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		ch1 <- "1"
		ch1 <- "2"
		ch1 <- "3"
		ch1 <- "4"
		ch1 <- "5"
		close(ch1)
	}()

	go func() {
		ch2 <- "6"
		ch2 <- "7"
		ch2 <- "8"
		ch2 <- "9"
		ch2 <- "10"
		close(ch2)
	}()

	out := multiplex(ch1, ch2)

	for val := range out {
		fmt.Println(val)
	}

}

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
