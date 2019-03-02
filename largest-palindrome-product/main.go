package main

import (
	"math"
	"strconv"
)

func isPalindrome(s string) bool {
	strLen := len(s)
	strMid := strLen / 2

	for i := 0; i < strMid; i++ {
		if s[i] != s[strLen-i-1] {
			return false
		}
	}

	return true
}

func LargestPalindromeProduct(factorLen int) int {
	var largest int

	factorMin := int(math.Pow10(factorLen - 1))
	factorMax := int(math.Pow10(factorLen) - 1)

	for i := factorMin; i <= factorMax; i++ {
		for j := factorMin; j < factorMax; j++ {
			product := i * j
			productStr := strconv.Itoa(product)

			if isPalindrome(productStr) && product > largest {
				largest = product
			}
		}
	}

	return largest
}

func main() {
}
