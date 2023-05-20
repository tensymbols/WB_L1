package main

import "fmt"

func main() {
	sequence := []string{
		"cat", "cat", "dog", "cat", "tree",
	}
	set := map[string]bool{}
	for _, v := range sequence {
		set[v] = true
	}
	for k := range set {
		fmt.Println(k)
	}
}
