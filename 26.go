package main

import (
	"fmt"
	"strings"
)

func strUniqueChars(s string) bool {
	sr := []rune(strings.ToLower(s))
	sMap := map[rune]bool{}
	for _, v := range sr {
		_, ok := sMap[v]
		if ok {
			return false
		}
		sMap[v] = true
	}
	return true
}

func main() {
	inputString1 := "abCdefAaf"
	inputString2 := "abcd"
	inputString3 := "aabcd"
	fmt.Println(
		strUniqueChars(inputString1),
		strUniqueChars(inputString2),
		strUniqueChars(inputString3))
}
