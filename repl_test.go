package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   hello    world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Actual length does not match expected length.\nActual: %d\nExpected: %d", len(actual), len(c.expected))
			continue
		}

		for i := range c.expected {
			if actual[i] != c.expected[i] {
				t.Errorf("Actual word doesn't match Expected word at index %d.\nActual: %s\nExpected: %s", i, actual[i], c.expected[i])
			}
		}
	}
}
