package main

import "fmt"

func main() {
	sequence := []string{ // создаем слайс неуникальных строк
		"cat", "cat", "dog", "cat", "tree",
	}
	set := map[string]bool{}     // создаем множество
	for _, v := range sequence { // добавляем все элементы из слайса
		set[v] = true // можем не проверять на наличие, тк неуникальные значения просто перезапишутся
	}
	// теперь множество хранит только уникальные значения
	for k := range set {
		fmt.Println(k)
	}
}
