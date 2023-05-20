package main

import (
	"fmt"
	"reflect"
)

func PrintType(v any) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("it's an integer")
	case reflect.String:
		fmt.Println("it's a string")
	case reflect.Bool:
		fmt.Println("it's a boolean")
	case reflect.Chan:
		fmt.Println("it's a channel")
	}
}
func main() {
	v := interface{}(123)
	PrintType(v)
}

// В switch v.(type) нет типа channel, поэтому только через reflect
