package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Закрытие канала
	ch := make(chan any)
	var wg sync.WaitGroup
	fmt.Println("Start")
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					wg.Done()
					fmt.Println("Channel is closed")
					return
				}
				fmt.Println(v)
			default:

			}
		}
	}()
	go func() {
		defer wg.Done()
		ch <- 1
		ch <- "hello"
		ch <- true
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	wg.Wait()

}
