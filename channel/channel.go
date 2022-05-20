package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("channel = ", <-ch)
	}()
	ch <- 1
	close(ch)
}
