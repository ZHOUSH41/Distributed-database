package struct_method

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter()  float64 {
	return 2*(r.width + r.height)
}

func (r Rectangle)Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (t Triangle)Area() float64 {
	return t.Base * t.Height * 0.5
}