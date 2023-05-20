package main

import (
	"fmt"
	"sync"
)

type CounterMutex struct {
	c  uint32
	mu sync.Mutex
}

func NewCounterMutex(n uint32) *CounterMutex {
	return &CounterMutex{c: n}
}

func (c *CounterMutex) Inc() {
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}
func (c *CounterMutex) GetValue() uint32 {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.c
}

func main() {
	cnt := NewCounterMutex(0)

	limit1 := uint32(1e6)
	limit2 := uint32(1e8)

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
