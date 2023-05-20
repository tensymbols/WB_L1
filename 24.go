package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) GetX() float64 { // геттер для поля x
	return p.x
}
func (p *Point) GetY() float64 { // геттер для поля y
	return p.y
}
func NewPoint(x float64, y float64) *Point { // конструктор точки
	return &Point{
		x: x,
		y: y,
	}
}

func FindDistanceBetweenPoints(p1 Point, p2 Point) float64 { // находим расстояние
	firstX := p1.GetX()
	firstY := p1.GetY()
	secondX := p2.GetX()
	secondY := p2.GetY()                                                                    // инициализируем временные переменные заранее чтобы не вызывать геттеры много раз
	return math.Sqrt((firstX-secondX)*(firstX-secondX) + (firstY-secondY)*(firstY-secondY)) // формула Пифагора

}
func main() {
	p1 := NewPoint(1, 5) // инициализируем две точки с заданными координатами
	p2 := NewPoint(17, 4)
	fmt.Println(FindDistanceBetweenPoints(*p1, *p2))
}
