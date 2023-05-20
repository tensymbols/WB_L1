package main

import (
	"fmt"
	"sync"
)

type AsyncMap struct { // структура асинхронной мапы хранящая внутри данные(обычная мапа) и мьютекс
	data map[any]any
	mu   sync.Mutex
}

func (am *AsyncMap) Put(k any, v any) {
	defer am.mu.Unlock() // анлок мьютекса после выхода из функции
	am.mu.Lock()         // лок мьютекса для записи в мапу
	am.data[k] = v       // запись в мапу

}
func (am *AsyncMap) Get(k any) any {
	defer am.mu.Unlock() // анлок мьютекса после выхода из функции
	am.mu.Lock()         // блокируем мьютекс
	return am.data[k]    // возвращаем значение, после чего будет выход из функции и анлок мьютекса
}
func (am *AsyncMap) PrintAll() {
	defer am.mu.Unlock()
	am.mu.Lock()
	for k, v := range am.data {
		fmt.Println(k, v)
	}

}
func NewAsyncMap() *AsyncMap { // конструктор мапы
	return &AsyncMap{
		data: map[any]any{},
		mu:   sync.Mutex{},
	}
}

func main() {
	m := NewAsyncMap()
	var wg sync.WaitGroup
	wg.Add(2)
	// Далее в двух горутинах добавляем значения в мапу, то есть асинхронно пишем в нее
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			m.Put(i*i, true)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			m.Put(i*i*2, true)
		}
	}()
	wg.Wait()
	m.PrintAll()

}
