package concepts

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() string
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() string {
	area := math.Pi * c.Radius * c.Radius
	return fmt.Sprintf("Circle area: %.2f", area)
}

func (r Rectangle) Area() string {
	area := r.Width * r.Height
	return fmt.Sprintf("Rectangle area: %.2f", area)
}


func PrintArea(s Shape) {
	fmt.Println(s.Area())
}