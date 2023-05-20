package main

import "fmt"

type Client struct{} // клиент

type Walker interface { // интерфейс
	Walk()
}

type Man struct { // структура Man (Human уже в этом пакете определена), содержащая произвольные поля
	name string
	age  int
}
type Lizard struct { // структура Lizard содержащая другие произвольные поля и не удовлятворяющая Walker
	color string
}

type LizardAdapter struct { // адаптер содержащий Lizard, которую нужно адаптировать под интерфейс Walker
	Lizard
}

func (h *Man) Walk() { // Man имеет метод Walk(), поэтому удовлетворяет интерфейсу Walker
	fmt.Printf("Human %s is walking\n", h.name)
}

func (l *Lizard) Crawl() { // несовместимый метод
	fmt.Printf("%s lizard is crawling\n", l.color)
}

func (la *LizardAdapter) Walk() { // метод делающий адаптер удовлетворяющим интерфейсу Walker
	la.Lizard.Crawl()
}
func (c *Client) WalkSomewhere(w ...Walker) { // метод который принимает на вход только интерфейс(ы) Walker
	for _, v := range w {
		v.Walk()
	}
}

func main() {
	man := &Man{name: "Bob", age: 25}                       // создание экземпляра Man
	lizardAdapter := &LizardAdapter{Lizard{color: "Green"}} // создание адаптера LizardAdapter содержащего внутри Lizard
	client := Client{}
	client.WalkSomewhere(man, lizardAdapter) // вызываем WalkSomewhere и передаем туда Walker'ы
	// Теперь наша ящерица умеет ходить как человек
}
