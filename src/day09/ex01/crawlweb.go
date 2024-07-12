package ex01

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func crawlWeb(ctx context.Context, inputChan chan string) chan string {
	results := make(chan string)
	semaphore := make(chan struct{}, 8)

	var wg sync.WaitGroup

	go func() {
		defer close(results)

		for url := range inputChan {

			select {
			case <-ctx.Done():
				return
			case semaphore <- struct{}{}:
			}

			wg.Add(1)

			go func(url string) {
				defer wg.Done()
				defer func() { <-semaphore }()

				body, err := fetch(url)

				select {
				case <-ctx.Done():
					return
				case results <- fmt.Sprintf("URL: %s Result: %p", url, getResultString(body, err)):
				}

			}(url)
		}

		wg.Wait()
	}()

	return results
}

func fetch(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getResultString(body string, err error) *string {
	if err != nil {
		result := fmt.Sprintf("error getting web page body: %v", err)
		return &result
	}

	return &body
}
