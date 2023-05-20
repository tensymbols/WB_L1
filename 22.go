package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a big.Int
	var b big.Int

	a.Exp(big.NewInt(10), big.NewInt(20), nil)
	b.Exp(big.NewInt(7), big.NewInt(25), nil)
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Add(&a, &b)
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Mul(&b, big.NewInt(25))
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Sub(&b, &a)
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Div(&a, &b)
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
}
