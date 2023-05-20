package main

import (
	"fmt"
	"math/rand"
)

// Последствия неправильной реализации в том что делая срез в 100 элементов мы делаем срез в 100 байт,
// соотвественно мы можем обрезать нашу строку так что она станет невалидной

func createHugeString(n int) string {
	alphabet := []rune("абвгдеёжзийклмнопрстуфхцчшщыьъэюяabcdefghijklmnopqrstuvwxyz0123456789")
	s := make([]rune, n)
	for i := 0; i < n; i++ {
		s[i] = alphabet[rand.Int31n(int32(len(alphabet)))]
	}
	return string(s)
}

var justString string

func someFunc() {
	v := []rune(createHugeString(1 << 10))
	justString = string(v[:100])
}
func main() {
	someFunc()
	fmt.Println(justString)
}
