package main

import "fmt"

type Stage func(<-chan any) <-chan any

func Stage1(in <-chan any) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(int) * 2
		}
	}()
	return out
}

type Out <-chan any

func Pipeline(in <-chan any) <-chan any {
	stages := []Stage{Stage1}
	for _, v := range stages {
		in = v(in)
	}
	return in
}

func main() {
	in := make(chan any)
	var out Out
	data := []int{
		1, 2, 3, 4, 777, 2e8, -6, 15,
	}
	go func() {
		defer close(in)
		for _, v := range data {
			in <- v
		}
	}()
	out = Pipeline(in)
	for v := range out {
		fmt.Println(v)
	}
}
