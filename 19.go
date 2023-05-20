package main

import "fmt"

func revertString(s string) string {
	rStr := []rune(s)            // создаем слайс рун из входной строки
	rOut := make([]rune, len(s)) // создаем выходной слайс рун длины входной строки
	for i, v := range rStr {     // проходимся по всем рунам строки
		rOut[len(rOut)-i-1] = v // каждому символу выходного слайса рун под индексом (длина строки - текущий индекс -1) присваиваем текущее значение
	}
	return string(rOut) // кастим слайс рун к строке и возвращаем
}

func main() {
	inputString := "абырвалг" // главрыба
	fmt.Println(revertString(inputString))
}
