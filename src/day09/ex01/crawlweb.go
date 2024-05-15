package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	urls := make(chan string)

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-cancelChan
		fmt.Println("\nStopping...\nGraceful shutdown")
		cancel()
	}()

	go func() {
		urlList := []string{
			"https://www.google.com",
		}

		for _, url := range urlList {
			urls <- url
		}

		close(urls)
	}()

	results := crawlWeb(ctx, urls)

	for result := range results {
		fmt.Println(result)
	}
}

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
				case results <- getResultString(body, err):
				}

			}(url)
		}

		wg.Wait()
		close(semaphore)
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

func getResultString(body string, err error) string {
	if err != nil {
		return fmt.Sprintf("error getting web page body: %v", err)
	}

	return body
}
