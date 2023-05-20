package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Завершение контекста
	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст с функцией отмены
	var wg sync.WaitGroup                                   // вейтгруппа
	fmt.Println("Start")
	wg.Add(2) //
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context was canceled") // в случае отмены контекста (в данной ситуации из другой горутины) выходим
				return
			default:
				// do smth
			}
		}

	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second) // симуляция бурной деятельности горутины которая занимает 2 секунды
		cancel()                    // отменяем контекст
	}()
	wg.Wait()
	// программа успешно завершается, значит вейтгруппа разблокировала мейн, то есть все горутины успешно закончили работу
}
