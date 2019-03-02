package main

import (
	"errors"
	"fmt"
)

func SumSquareDiff(max int) (int, error) {
	if max < 0 {
		return 0, errors.New("Max value must be positive")
	}

	sum := 0
	sumOfSquares := 0

	for i := 1; i <= max; i++ {
		sum += i
		sumOfSquares += i * i
	}

	squareOfSum := sum * sum
	difference := squareOfSum - sumOfSquares

	return difference, nil
}

func main() {
	fmt.Println(SumSquareDiff(100))
}
