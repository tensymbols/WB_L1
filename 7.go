package main

import (
	"fmt"
	"sync"
)

type AsyncMap struct {
	data map[any]any
	mu   sync.Mutex
}

func (am *AsyncMap) Put(k any, v any) {
	defer am.mu.Unlock()
	am.mu.Lock()
	am.data[k] = v

}
func (am *AsyncMap) Get(k any) any {
	defer am.mu.Unlock()
	am.mu.Lock()
	return am.data[k]
}
func (am *AsyncMap) PrintAll() {
	defer am.mu.Unlock()
	am.mu.Lock()
	for k, v := range am.data {
		fmt.Println(k, v)
	}

}
func NewAsyncMap() *AsyncMap {
	return &AsyncMap{
		data: map[any]any{},
		mu:   sync.Mutex{},
	}
}

func main() {
	m := NewAsyncMap()
	var wg sync.WaitGroup
	wg.Add(2)
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
