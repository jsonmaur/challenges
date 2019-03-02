package main

import "testing"

func TestSumSquareDiff(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{10, 2640},
		{100, 25164150},
	}

	for _, test := range tests {
		result, err := SumSquareDiff(test.input)
		if err != nil {
			t.Error(err)
		}
		if result != test.output {
			t.Errorf("Expected %v for %v, got %v", test.output, test.input, result)
		}
	}
}
