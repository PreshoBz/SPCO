package main

// Creating a Struct called Triangle
type Triangle struct {
	base   float64
	heigth float64
}

// Creating method called area
func (t Triangle) Area() float64 {
	return 1.5 * (t.base * t.heigth)
}
