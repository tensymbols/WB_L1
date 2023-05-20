package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func startWorkers(q int, mainCh chan any) {

	var wg sync.WaitGroup
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM) // отмена контекста по сигналам
	// sig interrupt и sig terminate

	wg.Add(q) // добавляем в вейтгруппу q воркеров
	for i := 0; i < q; i++ {
		go func() {
			defer wg.Done()
			for { // бесконечный селект
				select {
				case <-ctx.Done(): // в случае есть контекст завершен
					fmt.Println("Stopping goroutine")
					return // завершаем горутину
				case v := <-mainCh: // если пришло значение из главного канала
					fmt.Println(v)
				}
			}
		}()
	}
	wg.Wait() // ждем пока счетчик вейтгруппы станет равен 0
}

func main() {
	var workersQuantity int  // количество воркеров
	mainCh := make(chan any) // главный канал
	n := 10                  // любое число входящих значений

	flag.IntVar(&workersQuantity, "workers", 4, "quantity of workers")
	flag.Parse()

	go func() { // записываем значения в главный канал в горутине
		for i := 0; i < n; i++ {
			mainCh <- i
		}
	}()

	startWorkers(workersQuantity, mainCh)

}
