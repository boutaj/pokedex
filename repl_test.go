package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input string
		expected []string
	}{
		{
			input:    "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "one TWO three",
			expected: []string{"one", "two", "three"},
		},
		{
			input:    "   hEllO              WoRLD ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "          ",
			expected: []string{},
		},
		{
			input:    "   HELLO  ",
			expected: []string{"hello"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected '%v', got %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected '%s', got %s", expectedWord, word)
			}
		}
	}
}