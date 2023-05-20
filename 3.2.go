package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func startWorkingPoolAtomic(q int, inCh chan int64, sum *int64, wg *sync.WaitGroup) {
	for i := 0; i < q; i++ { // цикл на запуск воркеров
		go func() {
			defer wg.Done()
			for v := range inCh { // пока канал inCh не закрыт берем из него значение
				atomic.AddInt64(sum, v*v) // атомарно (потокобезопасно) добавляем к переменной суммы квадрат значения v
			}
		}()
	}
}

func main() {

	// Все аналогично 1 и 2 заданиям
	const workersQuantity = 4

	nums := []int64{2, 4, 6, 8, 10} // входные данные

	inCh := make(chan int64)

	go func() { // горутина на запись в канал
		defer close(inCh)
		for _, v := range nums {
			inCh <- v
		}

	}()

	var wg sync.WaitGroup   // вейтгруппа для воркеров
	wg.Add(workersQuantity) // увеличиваем счетчик вейтгруппы на workersQuantity
	var sum int64
	startWorkingPoolAtomic(workersQuantity, inCh, &sum, &wg)
	wg.Wait() // ожидание завершения горутин  вейтгруппы

	fmt.Println(sum)

}
