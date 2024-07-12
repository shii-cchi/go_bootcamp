package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

const maxConcurrency = 8

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		<-cancelChan
		fmt.Println("\nStopping...\nGraceful shutdown")
		cancel()
	}()

	urlList := []string{
		"https://www.google.com",
		"https://edu.21-school.ru/",
		"https://github.com/",
		"https://translate.yandex.ru/",
		"https://web.telegram.org/",
		"https://www.youtube.com/",
		"https://rocketchat-student.21-school.ru/",
		"https://shikimori.one/",
		"https://mail.google.com/",
	}

	urls := make(chan string, len(urlList))

	for _, url := range urlList {
		urls <- url
	}

	close(urls)

	results := crawlWeb(ctx, urls)

	for result := range results {
		fmt.Println(result)
	}
}

func crawlWeb(ctx context.Context, inputChan chan string) chan string {
	outChan := make(chan string)
	semaphore := make(chan struct{}, maxConcurrency)

	var wg sync.WaitGroup

	go func() {
		defer close(outChan)

		for url := range inputChan {
			select {
			case <-ctx.Done():
				return
			case semaphore <- struct{}{}:
				wg.Add(1)

				go func(url string) {
					defer wg.Done()
					defer func() { <-semaphore }()

					body, err := fetch(url)

					if err != nil {
						outChan <- fmt.Sprintf("URL: %s Result: error getting web page body: %v", url, err)
					} else {
						outChan <- fmt.Sprintf("URL: %s Result: %p", url, &body)
					}
				}(url)
			}
		}

		wg.Wait()
	}()

	return outChan
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
