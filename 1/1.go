package main

import "fmt"

type Human struct {
	name string
}
type Action struct {
	Human
}

func (h Human) Walk(dist int) {
	fmt.Printf("%s walked %d steps\n", h.name, dist)
}

func (h Human) Jump() {
	fmt.Printf("%s jumped\n", h.name)
}

func (h Human) SaySomething(s string) {
	fmt.Printf("%s said: %s\n", h.name, s)
}

func main() {
	h := Human{name: "Joe Schmoe"}
	action := Action{h}
	action.Walk(15)
	action.Jump()
	action.SaySomething("Hello!")
}
