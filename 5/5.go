package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	const N = 5
	ch := make(chan any)
	ctx, _ := context.WithTimeout(context.Background(), N*time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- 1
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(<-ch)
			}
		}
	}()
	<-ctx.Done()
}
