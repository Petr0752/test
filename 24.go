package main

import (
	"fmt"
	"math"
)

// Point - структура для представления точки на плоскости
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания экземпляра структуры Point
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo - метод для вычисления расстояния между двумя точками
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// создание двух экземпляров структуры Point
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	// вычисление расстояния между двумя точками
	distance := p1.DistanceTo(p2)

	fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) равно %.2f\n", p1.x, p1.y, p2.x, p2.y, distance)
}
