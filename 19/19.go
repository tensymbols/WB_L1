package main

import "fmt"

func revertString(s string) string {
	rStr := []rune(s)
	rOut := make([]rune, len(s))
	for i, v := range rStr {
		rOut[len(rOut)-i-1] = v
	}
	return string(rOut)
}

func main() {
	inputString := "абырвалг"
	fmt.Println(revertString(inputString))
}
