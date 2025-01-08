package main 

import (
	"math"
	"fmt"
) 

type Shape interface {
	area() float64
	perimeter() float64
}


type Circle struct {
	Radius float64
} 

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.Radius
}


type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}


func PrintShapeDetails(s interface{}) {
	/*if s.(type) == Shape {
		s := s.(Shape)
		fmt.Printf("Area: %f, Perimeter: %f\n", s.area(), s.perimeter())		
	}

	fmt.Println("Unknown shape")*/

	switch s.(type) {
		case Circle:
			c := s.(Circle)
			fmt.Printf("Radius: %f, Area: %f, Perimeter: %f\n", c.Radius, c.area(), c.perimeter())
		case Shape:
			s := s.(Shape)
			fmt.Printf("Area: %f, Perimeter: %f\n", s.area(), s.perimeter())
		default:
			fmt.Println("Unknown shape")
	}
}


func main(){
	c := Circle{Radius: 5}
	r := Rectangle{Width: 5, Height: 7}
	t := Triangle{SideA: 3, SideB: 4, SideC: 5}

	PrintShapeDetails(c)
	PrintShapeDetails(r)
	PrintShapeDetails(t)
}