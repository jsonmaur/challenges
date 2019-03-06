package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, test := range tests {
		for index, result := range ThreeSum(test.input) {
			if !reflect.DeepEqual(result, test.output[index]) {
				t.Fail()
			}
		}
	}
}
