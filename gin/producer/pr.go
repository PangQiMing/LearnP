package main

import (
	"fmt"
	"time"
)

var (
	infos = make(chan int, 10)
)

func producer(index int) {
	infos <- index
	fmt.Printf("producer %d sent %d\n", index, index)
}

func consumer(index int) {
	fmt.Printf("consumer %d receive %d\n", index, <-infos)
}

func main() {
	for i := 0; i < 10; i++ {
		go producer(i)
	}

	for i := 0; i < 100; i++ {
		go consumer(i)
	}
	time.Sleep(20 * time.Second)
}
