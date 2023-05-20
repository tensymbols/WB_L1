package main

import (
	"fmt"
	"math/rand"
)

// Последствия неправильной реализации в том, что делая срез в 100 элементов мы делаем срез в 100 байт,
// соотвественно мы можем обрезать нашу строку так что она станет невалидной

func createHugeString(n int) string {
	alphabet := []rune("абвгдеёжзийклмнопрстуфхцчшщыьъэюяabcdefghijklmnopqrstuvwxyz0123456789") // создаем алфавит с переменной размерностью рун
	s := make([]rune, n)                                                                        // создаем слайс рун и выдяляем ему память на н значений
	for i := 0; i < n; i++ {
		s[i] = alphabet[rand.Int31n(int32(len(alphabet)))] // берем случайное значение из алфавита и записываем в слайс рун
	}
	return string(s) // кастим слайс рун в строку и возвращаем  ее
}

var justString string

func someFunc() {
	v := []rune(createHugeString(1 << 10)) // исправленная версия кастит строку в слайс рун, объяснение почему сверху
	justString = string(v[:100])
}
func main() {
	someFunc()
	fmt.Println(justString)
}
