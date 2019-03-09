package main

import "testing"

func TestLongestNonrepeatingSubstr(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdp", 3},
		{" ", 1},
		{"", 0},
	}

	for _, test := range tests {
		result := LongestNonrepeatingSubstr(test.input)
		if result != test.output {
			t.Errorf("Expected %v for %v, got %v", test.output, test.input, result)
		}
	}
}
