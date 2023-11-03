package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Height + rec.Width)
}

func Area(rec Rectangle) float64 {
	return rec.Height * rec.Width
}
