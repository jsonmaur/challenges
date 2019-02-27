package main

import "testing"

func TestFindSumOfEvens(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{4000000, 4613732},
		{2, 2},
		{0, 0},
	}

	for _, test := range tests {
		result := FindSumOfEvens(test.input)
		if result != test.output {
			t.Errorf("%v should return %v", test.input, test.output)
		}
	}
}
