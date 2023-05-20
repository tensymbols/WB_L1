package main

import (
	"fmt"
	"sync"
)

func startSquaringWorkers(q int, inCh chan int, wg *sync.WaitGroup) chan int { // Запускаем пул
	outCh := make(chan int) // Выходной канал
	for i := 0; i < q; i++ {
		go func() { // Каждая из q горутин будет считывать значения из канала inCh и записывать результат преобразования в outCh
			defer wg.Done()
			for v := range inCh {
				outCh <- v * v
			}
		}()
	}
	return outCh
}

func main() {
	const workersQuantity = 4 // количество воркеров

	nums := []int{2, 4, 6, 8, 10} // входные данные

	inCh := make(chan int) // канал в который будут записываться входные данные

	go func() { // запись в канал
		defer close(inCh)
		for _, v := range nums {
			inCh <- v
		}

	}()
	var wg1 sync.WaitGroup // вейтгруппа для воркеров
	var wg2 sync.WaitGroup // вейтгруппа для вывода
	wg1.Add(workersQuantity)
	outCh := startSquaringWorkers(workersQuantity, inCh, &wg1)
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		for v := range outCh {
			fmt.Println(v)
		}
	}()
	wg1.Wait()
	close(outCh)
	wg2.Wait()

}
