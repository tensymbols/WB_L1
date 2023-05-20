package main

import "fmt"

func removeIthElement(sliceOfAny []any, index int) []any {
	if len(sliceOfAny) == 0 || index > len(sliceOfAny)-1 {
		panic("invalid index") // если слайс пустой или индекс больше допустимого то паникуем
	}
	copy(sliceOfAny[index:], sliceOfAny[index+1:]) // копируем данные в в диапазон с index до конца из диапазона index+1 до конца
	return sliceOfAny[:len(sliceOfAny)-1]          // возвращаем слайс от входного слайса без последнего элемента
}

func main() {
	sliceOfAny := []any{
		1, 5.0, "abcdefg", struct {
			float64
			int
		}{}, -6, "str", "str", 123, // слайс из.. всего
	}
	fmt.Println(sliceOfAny)
	sliceOfAny = removeIthElement(sliceOfAny, 2) // удаляем элемент под индексом 2
	fmt.Println(sliceOfAny)
}
