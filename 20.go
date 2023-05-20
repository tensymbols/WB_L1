package main

import (
	"fmt"
	"strings"
)

func RevertWords(s string) string {
	tokens := strings.Split(s, " ") // разбиваем входную строку на токены
	sOut := ""                      // инициализируем выходную строку как пустую
	for i := len(tokens) - 1; i >= 0; i-- {
		sOut = fmt.Sprintf("%s %s", sOut, tokens[i]) // в обратном порядке конкатенируем токены
	}
	sOut = strings.TrimSpace(sOut) // убираем лишний пробел
	return sOut
}

func main() {
	inputString := "snow dog sun" // sun dog snow
	fmt.Println(RevertWords(inputString))
}
