package main

import "fmt"

const bufferSize = 10

func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < bufferSize; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int, done chan<- struct{}) {
	for i := range ch {
		fmt.Println(i)
	}
	done <- struct{}{}
}

func main() {
	ch := make(chan int, bufferSize)
	done := make(chan struct{})
	go producer(ch)
	go consumer(ch, done)
	<-done
}
