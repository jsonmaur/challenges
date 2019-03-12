package main

import (
	"fmt"
	"math"
)

func ReverseInteger(x int) int {
	rev := 0
	min := math.MinInt32
	max := math.MaxInt32

	for x != 0 {
		pop := x % 10
		x /= 10

		if (rev > max/10 || (rev == max/10 && pop > 7)) ||
			(rev < min/10 || (rev == min/10 && pop < -8)) {
			return 0
		}

		rev = rev*10 + pop
	}

	return rev
}

func main() {
	fmt.Println(ReverseInteger(123))
}
