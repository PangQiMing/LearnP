package main

import (
	"fmt"
	"sync"
)

func main() {
	eggs := make(chan int, 10)
	for i := 0; i < 10; i++ {
		eggs <- i
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			select {
			case egg := <-eggs:
				fmt.Printf("people %d = %d\n", i, egg)
			default:
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
