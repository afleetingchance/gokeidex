package main

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"split words": {
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		"strip whitespace": {
			input:    "  hello  world ",
			expected: []string{"hello", "world"},
		},
		"lowercase": {
			input:    "HELLO WoRlD",
			expected: []string{"hello", "world"},
		},
		"single word": {
			input:    "thisisalloneword",
			expected: []string{"thisisalloneword"},
		},
		"empty": {
			input:    "",
			expected: []string{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(test.input)
			diff := cmp.Diff(test.expected, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
