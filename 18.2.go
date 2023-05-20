package main

import (
	"fmt"
	"sync"
)

type CounterMutex struct { //структур счетчика содержащая сам счетчик типа uint32 и мьютекса для избежания data race
	c  uint32
	mu sync.Mutex
}

func NewCounterMutex(n uint32) *CounterMutex { // конструктор счетчика с мьютексом
	return &CounterMutex{c: n}
}

func (c *CounterMutex) Inc() {
	c.mu.Lock()   // лок мьютекса
	c.c++         // икремент
	c.mu.Unlock() // анлок мьютекса
}
func (c *CounterMutex) GetValue() uint32 {
	defer c.mu.Unlock() // анлок мьютекса после выхода из функции
	c.mu.Lock()         // лочим мьютекс
	return c.c          // возвращаем значение счетчика, после чего произойдет анлок мьютекса
}

func main() {
	cnt := NewCounterMutex(0)

	limit1 := uint32(1e6)
	limit2 := uint32(1e8)

	// все аналогично первой реализации

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := uint32(0); i < limit1; i++ {
			cnt.Inc()
		}
	}()
	go func() {
		defer wg.Done()
		for i := uint32(0); i < limit2; i++ {
			cnt.Inc()
		}
	}()
	wg.Wait()
	if cnt.GetValue() != limit1+limit2 {
		panic("Counting failed")
	}
	fmt.Println("Counter value:", cnt.GetValue())

}
