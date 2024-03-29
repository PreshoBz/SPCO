package main

import (
	"fmt"
	"math"
)

// Creating a Struct called Triangle
type Triangle struct {
	base   float64
	heigth float64
}

// Creating method called area
func (t Triangle) Area() float64 {
	return 1.5 * (t.base * t.heigth)
}

// Creating method called perimeter
func (t Triangle) Perimeter() float64 {
	return 2*t.base + t.heigth
}

// Creating a Struct called Circle
type Circle struct {
	radius float64
}

// Creating method called area
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Creating method called perimeter
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Creating a function called square
func square(side float64) (float64, float64) {
	area := side * side
	perimeter := 4 * side
	return area, perimeter
}

func main() {

	//Creatng a variable of type 'triangle'
	tri1 := Triangle{
		base:   5.0,
		heigth: 4.0,
	}

	//Calling the methods on triangle
	fmt.Println(tri1.Area())
	fmt.Println(tri1.Perimeter())
}
