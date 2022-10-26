package main

import "math"

type Rectangle struct {
	W float64
	H float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	S float64
}
type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.H + rec.W)
}

func (r Rectangle) Area() float64 {
	return r.H * r.W
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s Square) Area() float64 {
	return s.S * s.S
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5

}
