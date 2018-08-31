package whatevs

import (
	"math"
)

func NewShape(shapeType string) Shape {
	switch shapeType {
	case "CIRCLE":
		return &Circle{}
	case "SQUARE":
		return &Square{}
	case "RECTANGLE":
		return &Rectangle{}
	}
}

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) SetRadius(r float64) {
	c.radius = r
}

type Square struct {
	sideLength float64
}

func (s *Square) GetArea() float64 {
	return s.sideLength * s.sideLength
}

type Rectangle struct {
	width  float64
	height float64
}
