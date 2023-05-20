package main

import (
	"fmt"
	"strings"
)

func RevertWords(s string) string {
	tokens := strings.Split(s, " ")
	sOut := ""
	for i := len(tokens) - 1; i >= 0; i-- {
		sOut = fmt.Sprintf("%s %s", sOut, tokens[i])
	}
	sOut = strings.TrimSpace(sOut)
	return sOut
}

func main() {
	inputString := "snow dog sun"
	fmt.Println(RevertWords(inputString))
}
