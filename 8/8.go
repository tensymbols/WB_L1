package main

import "fmt"

func MustSetBit(i int64, pos uint, val uint) int64 {
	one := int64(1)
	switch val {
	case 0:
		mask := ^(one << pos)
		return i & mask
	case 1:
		return i | (one << pos)
	}
	panic("Invalid value")
}
func main() {
	fmt.Println(MustSetBit(32, 0, 1))
}
