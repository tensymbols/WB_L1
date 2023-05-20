package main

import (
	"fmt"
	"strings"
)

func strUniqueChars(s string) bool {
	sr := []rune(strings.ToLower(s)) // кастим в слайс рун строку приведенную в ловеркейс
	sMap := map[rune]bool{}          // создаем мапу(фактически сет) для хранения символов
	for _, v := range sr {           // проходимся по всем символам строки
		_, ok := sMap[v]
		if ok {
			return false // если символ уже есть в сете, то есть встречался ранее в строке то вовзращаем false
		}
		sMap[v] = true // добавляем символ в сет
	}
	return true // возвращаем true тк до этого уже бы вернули false в случае неуникальных символов
}

func main() {
	inputString1 := "abCdefAaf" // входные строки
	inputString2 := "abcd"
	inputString3 := "aabcd"
	fmt.Println( // печатаем true если все символы в строке уникальные или false если нет
		strUniqueChars(inputString1),
		strUniqueChars(inputString2),
		strUniqueChars(inputString3))
}
