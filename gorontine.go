package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	runTime := 0

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine Done")
				return
			default:
				fmt.Printf("Goroutine Running TImes : %d\n", runTime)
				runTime = runTime + 1
			}
			if runTime > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
}
