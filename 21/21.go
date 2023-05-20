package main

import "fmt"

type Client struct{}

type Walker interface {
	Walk()
}

type Human struct {
	name string
	age  int
}
type Lizard struct {
	color string
}

type LizardAdapter struct {
	Lizard
}

func (h *Human) Walk() {
	fmt.Printf("Human %s is walking\n", h.name)
}

func (l *Lizard) Crawl() {
	fmt.Printf("%s lizard is crawling\n", l.color)
}

func (la *LizardAdapter) Walk() {
	la.Lizard.Crawl()
}
func (c *Client) WalkSomewhere(w ...Walker) {
	for _, v := range w {
		v.Walk()
	}
}

func main() {
	human := &Human{name: "Bob", age: 25}
	lizardAdapter := &LizardAdapter{Lizard{color: "Green"}}
	client := Client{}
	client.WalkSomewhere(human, lizardAdapter)
}
