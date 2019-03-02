package main

import (
	"errors"
	"fmt"
	"math"
)

func SmallestMultiple(min, max int) (int, error) {
	multiple := 1

	if min > max {
		return 0, errors.New("Min cannot be larger than max")
	}

	for i := min; i <= max; i++ {
		if multiple >= math.MaxInt64 {
			return 0, errors.New("Max Int value reached with no result")
		}
		if multiple%i != 0 {
			multiple++
			i = min
		}
	}

	return multiple, nil
}

func main() {
	fmt.Println(SmallestMultiple(1, 20))
}
