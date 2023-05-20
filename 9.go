package main

import "fmt"

type Stage func(<-chan any) <-chan any // определяем тип стейдж нашего большого пайплайна

func Stage1(in <-chan any) <-chan any { // функция первой стадии пайплайна
	// получаем на вход канал ин
	out := make(chan any) // создаем выходной канал
	go func() {           // горутина нашей стадии
		defer close(out)
		for v := range in {
			out <- v.(int) * 2 // само преобразование данных
		}
	}()
	return out // возвращаем выходной канал
}

type Out <-chan any

func Pipeline(in <-chan any) <-chan any { // функция исполняющая конвеер
	stages := []Stage{Stage1}  // слайс всех стадий (тут только одна)
	for _, v := range stages { // проходим по всем стадиям
		in = v(in)
		// присваиваем к in выходной канал каждой стадии, так получится что входной канал каждой следующей стадии
		// будет выходным каналом текущей
	}
	return in
}

func main() {
	in := make(chan any) // входной канал
	var out Out          // выходной канал
	data := []int{
		1, 2, 3, 4, 777, 2e8, -6, 15, // данные
	}
	go func() {
		defer close(in)
		for _, v := range data {
			in <- v // пишем данные во входной канал
		}
	}()
	out = Pipeline(in)   // присваиваем каналу out последний выходной канал конвеера
	for v := range out { // пока out открыт печатаем все значения из него
		fmt.Println(v)
	}
}
