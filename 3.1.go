package main

import (
	"fmt"
	"sync"
)

func startWorkingPool(q int, inCh chan int, wg *sync.WaitGroup) chan int { // пул воркеров проводящих вычисления
	outCh := make(chan int)
	for i := 0; i < q; i++ {
		go func() {
			defer wg.Done()
			for v := range inCh {

				outCh <- v * v // запись квадрата в канал
			}
		}()
	}
	return outCh
}

func main() {

	// Все аналогично 1 и 2 заданиям
	const workersQuantity = 4

	nums := []int{2, 4, 6, 8, 10}
	var sum int
	inCh := make(chan int)

	go func() { // горутина на запись в канал
		defer close(inCh)
		for _, v := range nums {
			inCh <- v
		}

	}()

	var wg1 sync.WaitGroup // вейтгруппа для воркеров
	var wg2 sync.WaitGroup // вейтгруппа для суммирования
	wg1.Add(workersQuantity)
	outCh := startWorkingPool(workersQuantity, inCh, &wg1)
	wg2.Add(1)
	go func() { // горутина на суммирование
		defer wg2.Done()

		for v := range outCh { // до тех пор пока outCh будет открыт горутина не завершит работу
			sum += v
		}
	}()
	wg1.Wait()   // ожидание завершения горутин первой вейтгруппы
	close(outCh) //закрытие выходного канала
	wg2.Wait()   // ожидание завершения горутин ( в нашем случае одной горутины) второй вейтгруппы
	fmt.Println(sum)

}
