package main

import (
    "testing"
)

func TestCircleArea(t *testing.T) {
    c := Circle{Radius: 5}
    expected := math.Pi * 25
    if c.Area() != expected {
        t.Errorf("Ожидалась площадь %.2f, получено %.2f", expected, c.Area())
    }
}

func TestRectangleArea(t *testing.T) {
    r := Rectangle{Width: 3, Height: 4}
    expected := 12.0
    if r.Area() != expected {
        t.Errorf("Ожидалась площадь %.2f, получено %.2f", expected, r.Area())
    }
}

func ExamplePrintArea() {
    PrintArea(Circle{Radius: 5})
    PrintArea(Rectangle{Width: 3, Height: 4})
}
