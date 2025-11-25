package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLLo WORld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "foo bar blah fasel A B",
			expected: []string{"foo", "bar", "blah", "fasel", "a", "b"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %d, got: %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected word: %s, got %s", expectedWord, word)
			}
		}
	}
}
