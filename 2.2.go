package main

import (
	"fmt"
	"sync"
)

func startSquaringWorkersAndPrint(q int, inCh chan int, wg *sync.WaitGroup) { // запуск пула и вывод
	for i := 0; i < q; i++ {
		go func() {
			defer wg.Done()
			for v := range inCh {
				v *= v
				fmt.Println(v)
			}
		}()
	}
}

func main() {
	const workersQuantity = 4

	nums := []int{2, 4, 6, 8, 10} // инпут

	inCh := make(chan int)

	go func() { // запись в канал
		defer close(inCh)
		for _, v := range nums {
			inCh <- v
		}
	}()
	var wg sync.WaitGroup // вейтгруппа для воркеров
	wg.Add(workersQuantity)
	startSquaringWorkersAndPrint(workersQuantity, inCh, &wg)
	wg.Wait()

}
