package main

import "fmt"

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high) // распределение

		quickSort(arr, low, pi-1) // рекурсивно делаем то же самое, пока low < high не перестанет выполняться
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high] // берем за пивот последнее значение с индексом high
	i := low - 1

	for j := low; j < high; j++ { // в этом цикле мы распределяем значения меньшие пивота слева, а большие справа
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1 // возвращаем индекс
}

func main() {
	arr := []int{1, 3, -5, 78, -2, 44, 4, 0}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
