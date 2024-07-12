package ex01

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func TestCrawlweb(t *testing.T) {
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
