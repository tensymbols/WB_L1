package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left := 0          // левый индекс
	right := len(arr)  // правый индекс равный длине массива
	var pivot int      // пивот (наш "указатель")
	for left < right { // пока левый индекс меньше правого
		pivot = (left + right) / 2 // присваиваем пивоту значение между левым и правым
		if arr[pivot] == target {  // если мы нашли таргет то возвращаем пивот(индекс под которым найденное значение)
			return pivot
		} else if arr[pivot] > target { // если значение под индексом пивота больше искомого, то теперь наш правый указатель будет указывать на 1 меньше пивота
			right = pivot - 1
		} else {
			left = pivot + 1 // если меньше, то на 1 больше
		}
	}
	return -1 // возвращаем индекс -1 если искомого значения нет в массиве
}

func main() {
	arr := []int{
		5, 4, 700, 89, 95, 66, 89561, -667, 298, 517, // случайные входные данные
	}
	sort.Slice(arr, func(i, j int) bool { // сортируем массив перед поиском
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	target := arr[rand.Intn(len(arr))] // выбираем за таргет случайное значение из массива
	fmt.Println("chosen target:", target)
	notExistingTarget := 1 // создаем несуществующее в массиве значение
	fmt.Println("target index:", binarySearch(arr, target))
	fmt.Println("non-existing target index:", binarySearch(arr, notExistingTarget))
}
