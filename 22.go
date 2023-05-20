package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a big.Int // создание двух переменных типа big.Int
	var b big.Int

	a.Exp(big.NewInt(10), big.NewInt(20), nil) // инициализация первой переменной как 10^20
	b.Exp(big.NewInt(7), big.NewInt(25), nil)  // инициализация второй переменной как 7^25
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Add(&a, &b) // сложение больших чисел
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Mul(&b, big.NewInt(25)) // умножение b на 25
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Sub(&b, &a) // вычитание а из b
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
	a.Div(&a, &b) // деление a на b
	fmt.Println("First big number:", a.String(), "\n"+
		"Second big number:", b.String())
}
