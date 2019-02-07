package main

import "fmt"

// interface is a collection of methods
type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 0
}

// Function argument from interface
func getShapeArea(s Shape) float64 {
	return s.area()
}

func main()  {
	var sampleRect = Rect{1, 2}

	// To check Rect struct totally implements Shape interface
	if _, ok := interface{}(sampleRect).(Shape); ok {
		fmt.Println("It works")
	}

	fmt.Println(sampleRect.area())

	fmt.Println(getShapeArea(sampleRect))
}