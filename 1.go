// Package geometry определяет геометрические фигуры и их характеристики.
package main

import (
    "fmt"
    "math"
)

// Shape представляет интерфейс геометрической фигуры.
type Shape interface {
    // Area возвращает площадь фигуры
    Area() float64
}

// Circle представляет круг с заданным радиусом
type Circle struct{ Radius float64 }

// Area вычисляет площадь круга
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// Rectangle представляет прямоугольник с заданными шириной и высотой
type Rectangle struct{ Width, Height float64 }

// Area вычисляет площадь прямоугольника
func (r Rectangle) Area() float64 { return r.Width * r.Height }

// PrintArea выводит площадь фигуры в стандартный вывод
func PrintArea(s Shape) {
    fmt.Printf("Площадь: %.2f\n", s.Area())
}
