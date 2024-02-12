package main

import (
	"testing"
)

// Testing the square function
func TestSquare(t *testing.T) {
	side := 5.0
	expectedArea := 25.0
	expectedPerimeter := 20.0

	area, perimeter := square(side)

	if area != expectedArea {
		t.Errorf("Area calculation is incorrect. Got: %f, Expected: %f", area, expectedArea)
	}

	if perimeter != expectedPerimeter {
		t.Errorf("Perimeter calculation is incorrect. Got: %f, Expected: %f", perimeter, expectedPerimeter)
	}
}
