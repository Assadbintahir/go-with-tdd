package struct_interface

import (
	"go-with-tdd/coverage"
	"testing"
)

// Enforce coverage
func TestMain(m *testing.M) {
	coverage.EnforceCoverage(m, 1, "struct_interface")
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	}
	t.Run("should return area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("should return area of circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTable(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "should return area of Rectangle", shape: Rectangle{12.0, 6.0}, expected: 72.0},
		{name: "should return area of Circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "should return area of Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()

			if got != test.expected {
				t.Errorf("%#v got %g but expected %g", test.shape, got, test.expected)
			}
		})
	}
}
