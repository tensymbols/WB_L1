package main

import (
	"fmt"
	"sort"
)

type temperatures []float64 // определяем тип temperatures как слайс float64

func main() {

	tempsByInterval := map[int64]temperatures{} // создаем мапу интервал-температуры

	tempData := temperatures{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, // входные значения температур
	}
	sort.Float64s(tempData) // сортируем данные

	currInterval := int64(tempData[0]) / 10 * 10 // вычисляем и берем первый интервал за текущий
	var currVals temperatures                    // текущие температуры
	for _, v := range tempData {
		tempInterval := int64(v) / 10 * 10 // вычисляем интервал к которому принадлежит значение v
		if tempInterval == currInterval {  // если оно принадлежит к текущему интервалу то добавляем в текущие значения это значение
			currVals = append(currVals, v)
		} else {
			tempsByInterval[currInterval] = currVals // иначе все текущие значения саписываем в мапу под ключом текущего интервала
			currVals = nil                           // обнуляем текущие значения
			currVals = append(currVals, v)           // добавляем в слайс текущих значений значение v
			currInterval = tempInterval              // теперь текущий интервал это временный интервал этой итерации ( интервал от v)
		}
	}
	tempsByInterval[currInterval] = currVals // добавляем последнее значение в мапу

	var allIntervals []int64 // теперь создаем слайс всех интервалов, это нужно для того чтобы их хранить в отсортированном порядке

	for k := range tempsByInterval {
		allIntervals = append(allIntervals, k) // добавляем все ключи из мапы в наш слайс
	}
	sort.Slice(allIntervals, func(i, j int) bool { // сортируем интервалы в слайсе
		return allIntervals[i] < allIntervals[j]
	})
	for _, v := range allIntervals {
		fmt.Println(v, tempsByInterval[v], "\n------") // выводим значения из мапы
	}

}
