package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterAtomic struct { // структура счетчика содержащая внутри сам счетчик типа uint32
	c uint32
}

func NewCounterAtomic(n uint32) *CounterAtomic { // конструктор атомарного счетчика
	return &CounterAtomic{c: n}
}

func (c *CounterAtomic) Inc() {
	atomic.AddUint32(&c.c, 1) // атомарно инкрементируем счетчик
}
func (c *CounterAtomic) GetValue() uint32 {
	return atomic.LoadUint32(&c.c) // атомарно получаем значение счетчика
}

func main() {
	cnt := NewCounterAtomic(0)

	limit1 := uint32(1e6) // первый лимит
	limit2 := uint32(1e8) // второй лимит

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := uint32(0); i < limit1; i++ { // в этой горутине мы limit1 раз инкрементируем счетчик
			cnt.Inc()
		}
	}()
	go func() {
		defer wg.Done()
		for i := uint32(0); i < limit2; i++ { // в этой горутины - limit2 раз
			cnt.Inc()
		}
	}()
	wg.Wait()
	if cnt.GetValue() != limit1+limit2 { // паникуем если по какой-то неведомой причины наш счетчик проинкрементировался неверное количество раз
		panic("Counting failed")
	}
	fmt.Println("Counter value:", cnt.GetValue())

}
