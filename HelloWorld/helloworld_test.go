package helloworld

import (
	"go-with-tdd/coverage"
	"testing"
)

// Coverage enforcement
func TestMain(m *testing.M) {
	coverage.EnforceCoverage(m, 1, "helloworld")
}

// Helper test assertion
func assertMessage(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q want %q", got, expected)
		t.Fail()
	}
}

// Unit tests in the form of explicit subtests
func TestHelloWorldSub(t *testing.T) {
	t.Run("Should say hello to people when name is provided", func(t *testing.T) {
		got := Hello("Asad", "English")
		expected := "Hello Asad"
		assertMessage(t, got, expected)
	})

	t.Run("Should say hello to world when name is not provided", func(t *testing.T) {
		got := Hello("", "English")
		expected := "Hello World"
		assertMessage(t, got, expected)
	})

	t.Run("Should say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Asad", "Spanish")
		expected := "Hola Asad"

		assertMessage(t, got, expected)
	})

	t.Run("Should say hello to people in French", func(t *testing.T) {
		got := Hello("Asad", "French")
		expected := "Bonjour Asad"

		assertMessage(t, got, expected)
	})
}

// Unit tests in the form of table tests
func TestHelloWorldTable(t *testing.T) {
	type parameters struct{ name, language string }
	tests := []struct {
		testName string
		param    parameters
		expected string
	}{
		{
			testName: "Should say hello to people in English when name is provided",
			param:    parameters{name: "Asad", language: "English"},
			expected: "Hello Asad",
		},
		{
			testName: "Should say hello to world in English when name is not provided",
			param:    parameters{name: "", language: "English"},
			expected: "Hello World",
		},
		{
			testName: "Should say hello to people in Spanish when name is provided",
			param:    parameters{name: "Asad", language: "Spanish"},
			expected: "Hola Asad",
		},
		{
			testName: "Should say hello to world in Spanish when name is provided",
			param:    parameters{name: "", language: "Spanish"},
			expected: "Hola World",
		},
		{
			testName: "Should say hello to people in French when name is provided",
			param:    parameters{name: "Asad", language: "French"},
			expected: "Bonjour Asad",
		},
		{
			testName: "Should say hello to people in French when name is provided",
			param:    parameters{name: "", language: "French"},
			expected: "Bonjour World",
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := Hello(test.param.name, test.param.language)
			assertMessage(t, got, test.expected)
		})
	}
}

func ExampleSayGreeting() {
	SayGreeting()
	// Output: Hello Asad
}
