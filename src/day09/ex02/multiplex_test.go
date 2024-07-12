package ex02

import (
	"fmt"
	"testing"
)

func TestMultiplex1(t *testing.T) {
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

func TestMultiplex2(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		ch2 <- 6
		ch2 <- 7
		ch2 <- 8
		ch2 <- 9
		ch2 <- 10
		close(ch2)
	}()

	out := multiplex(ch1, ch2)

	for val := range out {
		fmt.Println(val)
	}
}
