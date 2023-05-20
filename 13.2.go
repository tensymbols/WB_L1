package main

import "fmt"

func main() {
	a, b := 3, 4
	fmt.Println(a, b)
	a = a * b
	b = a / b
	a = a / b // своп умножением
	fmt.Println(a, b)
}
