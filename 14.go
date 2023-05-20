package main

import (
	"fmt"
	"reflect"
)

func PrintType(values ...any) { // функция печатающая типа
	for _, v := range values {
		switch reflect.TypeOf(v).Kind() { // выбираем из типов, чему равно значение типа переменной v
		case reflect.Int:
			fmt.Println("it's an integer") // печатает соответствующий типа
		case reflect.String:
			fmt.Println("it's a string")
		case reflect.Bool:
			fmt.Println("it's a boolean")
		case reflect.Chan:
			fmt.Println("it's a channel")
		}
	}

}
func main() {
	v1 := interface{}(123)            // создаем пустой интерфейс со значением инт
	v2 := interface{}("hello")        // строки
	v3 := interface{}(true)           // с булевым значением
	v4 := interface{}(make(chan any)) // со значением канала типа interface{}
	PrintType(v1, v2, v3, v4)
}

// В switch v.(type) нет типа channel, поэтому только через reflect
