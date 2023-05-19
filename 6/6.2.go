package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Завершение контекста
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	fmt.Println("Start")
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context was canceled")
				return
			default:
				// do smth
			}
		}

	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		cancel()
	}()
	wg.Wait()

}
