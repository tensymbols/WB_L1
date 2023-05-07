package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func mustWriteOut(bw *bufio.Writer, mu *sync.Mutex, msg any) {
	mu.Lock()
	_, err := bw.WriteString(fmt.Sprintln(msg))

	if err != nil {
		panic(err.Error())
	}
	_ = bw.Flush()
	mu.Unlock()
}

func startWorkers(q int, mainCh chan any) {

	var mu sync.Mutex
	var wg sync.WaitGroup

	bw := bufio.NewWriter(os.Stdout)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	wg.Add(q)
	for i := 0; i < q; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					mustWriteOut(bw, &mu, "Stopping goroutine")
					return
				case v := <-mainCh:
					mustWriteOut(bw, &mu, v)
				}
			}
		}()
	}
	wg.Wait()
}

func main() {
	var workersQuantity int
	mainCh := make(chan any)
	n := 10 // any number

	flag.IntVar(&workersQuantity, "workers", 4, "quantity of workers")
	flag.Parse()

	go func() {
		for i := 0; i < n; i++ {
			mainCh <- i
		}
	}()

	startWorkers(workersQuantity, mainCh)

}
