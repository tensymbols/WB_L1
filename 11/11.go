package main

import "fmt"

func getSetsIntersection(s1 []any, s2 []any) []any {
	var res []any
	s2Map := map[any]int{}
	for _, v := range s2 {
		s2Map[v] = s2Map[v] + 1
	}
	for _, v := range s1 {
		if s2Map[v] != 0 {
			res = append(res, v)
			s2Map[v]--
		}
	}
	return res
}

func main() {
	set1 := []any{
		661, 2, 2, 1, 3, 9, 7, 6, 100,
	}
	set2 := []any{
		2, 8, 8, 16, 3, 661, 1, 448, 1e5,
	}
	intersectionSet := getSetsIntersection(set1, set2)
	for _, v := range intersectionSet {
		fmt.Println(v)
	}
}
