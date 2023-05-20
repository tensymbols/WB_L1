package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	const N = 5
	ch := make(chan any)
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second) // через N секунд после создания контекста произойдет его отмена
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done(): // в случае отмены контекста завершаем горутину
				return
			default:
				ch <- 1 // записываем в канал данные
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done(): // в случае отмены контекста завершаем горутину
				return
			default:
				fmt.Println(<-ch) // печатаем в stdout данные прочитанные из канала
			}
		}
	}()
	<-ctx.Done() // main будет заблокирован на данной точке до тех пор пока контекст не будет отменен

}
