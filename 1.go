package main
 
import (
    "fmt"
    "math"
)
 
type Shape interface {
    Area() float64
}
 
type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
 
func (c Circle) Area() float64    { return math.Pi * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
 
func PrintArea(s Shape) {
    fmt.Printf("Площадь: %.2f\n", s.Area())
}
 
func main() {
    PrintArea(Circle{Radius: 5})
    PrintArea(Rectangle{Width: 3, Height: 4})
}
