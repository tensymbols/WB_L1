package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr)
	var pivot int
	for left < right {
		pivot = (left + right) / 2
		if arr[pivot] == target {
			return pivot
		} else if arr[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}

func main() {
	arr := []int{
		5, 4, 700, 89, 95, 66, 89561, -667, 298, 517,
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	target := arr[rand.Intn(len(arr))]
	fmt.Println("chosen target:", target)
	notExistingTarget := 1
	fmt.Println("target index:", binarySearch(arr, target))
	fmt.Println("non-existing target index:", binarySearch(arr, notExistingTarget))
}
