package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func startWorkingPool(q int, inCh chan int, wg *sync.WaitGroup) chan int {
	outCh := make(chan int)
	for i := 0; i < q; i++ {
		go func() {
			defer wg.Done()
			for v := range inCh {

				outCh <- v * v
			}
		}()
	}
	return outCh
}

func main() {
	const workersQuantity = 4

	nums := []int{2, 4, 6, 8, 10}
	var sum int
	bw := bufio.NewWriter(os.Stdout)

	inCh := make(chan int)

	go func() {
		defer close(inCh)
		for _, v := range nums {
			inCh <- v
		}

	}()

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	wg1.Add(workersQuantity)
	outCh := startWorkingPool(workersQuantity, inCh, &wg1)
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		for v := range outCh {
			sum += v
		}
	}()
	wg1.Wait()
	close(outCh)
	wg2.Wait()
	_, err := bw.WriteString(fmt.Sprintln(sum))
	if err != nil {
		panic("error writing to stdout")
	}
	_ = bw.Flush()

}
