package main

import (
	"fmt"
)

func LongestNonrepeatingSubstr(s string) int {
	pos := 0
	longest := 0

	seen := make(map[byte]int)

	for pos < len(s) {
		char := s[pos]

		if _, ok := seen[char]; ok {
			if len(seen) > longest {
				longest = len(seen)
			}

			pos = seen[char] + 1
			seen = make(map[byte]int)
		} else {
			seen[char] = pos
			pos++
		}
	}

	if len(seen) > longest {
		longest = len(seen)
	}

	return longest
}

func main() {
	fmt.Println(LongestNonrepeatingSubstr("dvdp"))
}
