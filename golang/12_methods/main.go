package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func main() {
	c := Circle{3}
	r := Rectangle{4, 5}

	fmt.Println(circleArea(c))
	fmt.Println(rectangleArea(r))

	fmt.Println(c.area())
	fmt.Println(r.area())
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func rectangleArea(r Rectangle) float64 {
	return r.w * r.h
}
