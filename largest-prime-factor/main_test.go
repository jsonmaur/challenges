package main

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	tests := []struct {
		input  int64
		output int64
	}{
		{600851475143, 6857},
		{13195, 29},
	}

	for _, test := range tests {
		result := LargestPrimeFactor(test.input)
		if result != test.output {
			t.Errorf("%v should return %v, got %v", test.input, test.output, result)
		}
	}
}
