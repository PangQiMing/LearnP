package main

import (
	"fmt"
	"sync"
)

func action() {
	fmt.Println("Hello Goroutine!")
}

func main() {
	//go action()
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("Hello Goroutine! %d\n", num)
			wg.Done()
		}(i)
	}
	//time.Sleep(2 * time.Second)
	wg.Wait()
}
