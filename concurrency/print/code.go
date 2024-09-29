// 有三个函数，分别打印 "cat", "fish","dog" 要求每一个函数都用一个 goroutine，按照顺序打印 100 次。

// 此题目考察 channel，用三个无缓冲 channel，如果一个 channel 收到信号则通知下一个。

package main

import (
	"fmt"
	"time"
)

var cat = make(chan struct{})
var fish = make(chan struct{})
var dog = make(chan struct{})

func Cat() {
	<-cat
	fmt.Println("cat")
	fish <- struct{}{}
}

func Fish() {
	<-fish
	fmt.Println("fish")
	dog <- struct{}{}
}

func Dog() {
	<-dog
	fmt.Println("dog")
	cat <- struct{}{}
}

func main() {
	for i := 0; i < 100; i++ {
		go Cat()
		go Fish()
		go Dog()
	}
	cat <- struct{}{}
	time.Sleep(2 * time.Second)
}