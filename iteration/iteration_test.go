package iteration

import (
	"fmt"
	"go-with-tdd/coverage"
	"testing"
)

// Coverage enforcement
func TestMain(m *testing.M) {
	coverage.EnforceCoverage(m, 1, "iteration")
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if expected != got {
		t.Errorf("expected: %q, got: %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}
