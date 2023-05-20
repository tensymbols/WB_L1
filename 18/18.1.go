package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterAtomic struct {
	c uint32
}

func NewCounterAtomic(n uint32) *CounterAtomic {
	return &CounterAtomic{c: n}
}

func (c *CounterAtomic) Inc() {
	atomic.AddUint32(&c.c, 1)
}
func (c *CounterAtomic) GetValue() uint32 {
	return atomic.LoadUint32(&c.c)
}

func main() {
	cnt := NewCounterAtomic(0)

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
