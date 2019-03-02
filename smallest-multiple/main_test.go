package main

import "testing"

func TestSmallestMultiple(t *testing.T) {
	tests := []struct {
		inputMin int
		inputMax int
		output   int
	}{
		{1, 10, 2520},
		{1, 20, 232792560},
	}

	for _, test := range tests {
		result, err := SmallestMultiple(test.inputMin, test.inputMax)
		if err != nil {
			t.Error(err)
		}
		if result != test.output {
			t.Errorf("Expected %v, got %v", test.output, result)
		}
	}
}
