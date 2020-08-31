package integers

import (
	"fmt"
	"go-with-tdd/coverage"
	"testing"
)

// Coverage enforcement
func TestMain(m *testing.M) {
	coverage.EnforceCoverage(m, 1, "integers")
}

func TestAdder(t *testing.T) {
	type parameters struct{ first, second int }
	tests := []struct {
		name     string
		params   parameters
		expected int
	}{
		{
			name:     "should return 4 when first input is 2 and second is 2",
			params:   parameters{first: 2, second: 2},
			expected: 4,
		},
		{
			name:     "should return 90 when first input is 50 and second is 40",
			params:   parameters{first: 50, second: 40},
			expected: 90,
		},
		{
			name:     "should return 15 when first input is 7 and second is 8",
			params:   parameters{first: 7, second: 8},
			expected: 15,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Adder(test.params.first, test.params.second)

			if got != test.expected {
				t.Errorf("expected: %d, got: %d", test.expected, got)
			}
		})
	}
}

func ExampleAdder() {
	sum := Adder(2, 2)
	fmt.Println(sum)
	// Output: 4
}
