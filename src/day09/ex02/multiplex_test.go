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

func TestMultiplex3(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	ch4 := make(chan interface{})
	ch5 := make(chan interface{})
	ch6 := make(chan interface{})

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	go func() {
		ch3 <- 5
		ch3 <- 6
		close(ch3)
	}()

	go func() {
		ch4 <- 7
		ch4 <- 8
		close(ch4)
	}()

	go func() {
		ch5 <- 9
		ch5 <- 10
		close(ch5)
	}()

	go func() {
		ch6 <- 11
		ch6 <- 12
		close(ch6)
	}()

	out := multiplex(ch1, ch2, ch3, ch4, ch5, ch6)

	for val := range out {
		fmt.Println(val)
	}
}
