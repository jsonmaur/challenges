package main

import "testing"

func TestLargestPalindromeProduct(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{2, 9009},
		{3, 906609},
	}

	for _, test := range tests {
		result := LargestPalindromeProduct(test.input)
		if result != test.output {
			t.Errorf("Expected %v, got %v", test.output, result)
		}
	}
}
