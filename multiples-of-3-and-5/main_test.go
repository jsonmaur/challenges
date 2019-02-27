package main

import "testing"

func TestFindSumBelow(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1000, 233168},
		{10, 23},
		{1, 0},
		{-1, 0},
		{0, 0},
	}

	for _, test := range tests {
		result := FindSumBelow(test.input)
		if result != test.output {
			t.Errorf("Expected %v, got %v", test.output, result)
		}
	}
}
