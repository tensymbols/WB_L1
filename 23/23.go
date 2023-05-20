package main

import "fmt"

func removeIthElement(sliceOfAny []any, index int) []any {
	if len(sliceOfAny) == 0 || index > len(sliceOfAny)-1 {
		panic("invalid index")
	}
	copy(sliceOfAny[index:], sliceOfAny[index+1:])
	return sliceOfAny[:len(sliceOfAny)-1]
}

func main() {
	sliceOfAny := []any{
		1, 5.0, "abcdefg", struct {
			float64
			int
		}{}, -6, "str", "str", 123,
	}
	fmt.Println(sliceOfAny)
	sliceOfAny = removeIthElement(sliceOfAny, 2)
	fmt.Println(sliceOfAny)
}
