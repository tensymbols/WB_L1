package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func startWorkersAndPrint(q int, inCh chan int, wg *sync.WaitGroup) {
	bw := bufio.NewWriter(os.Stdout)
	var mu sync.Mutex
	for i := 0; i < q; i++ {
		go func() {
			defer wg.Done()
			for v := range inCh {
				v *= v
				mu.Lock()
				_, err := bw.WriteString(fmt.Sprintln(v))

				if err != nil {
					panic(err.Error())
				}
				_ = bw.Flush()
				mu.Unlock()
			}
		}()
	}
}

func main() {
	const workersQuantity = 4

	nums := []int{2, 4, 6, 8, 10}

	inCh := make(chan int)

	go func() {
		defer close(inCh)
		for _, v := range nums {
			inCh <- v
		}
	}()
	var wg sync.WaitGroup
	wg.Add(workersQuantity)
	startWorkersAndPrint(workersQuantity, inCh, &wg)
	wg.Wait()

}
