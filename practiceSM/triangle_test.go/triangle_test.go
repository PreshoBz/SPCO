package main

import (
	"testing"
)
// Testing the area function for Triangle
func TestingTriangleArea(t *testing.T) {
	t1 := triangle {
		base : 3,
		heigth : 4
	}
	expectedArea := 0.5 * 3 * 4
	area := t1.area()
	if area != expectedArea {
		t.Errorf("Area calculation is incorrect, got: %f, want: %f.", area, expectedArea)
	}

}