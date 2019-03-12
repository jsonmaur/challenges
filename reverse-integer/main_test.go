package main

import "testing"

func TestReverseInteger(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}

	for _, test := range tests {
		result := ReverseInteger(test.input)
		if result != test.output {
			t.Errorf("Expected %v for %v, got %v", test.output, test.input, result)
		}
	}
}
