package main

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Encode(decoded string) (string, error) {
	var encoded bytes.Buffer
	split := strings.Split(decoded, "")

	if len(split) == 0 {
		return "", nil
	}

	isNum, _ := regexp.Compile("[0-9]")
	current := split[0]
	count := 0

	for _, char := range split {
		if isNum.MatchString(char) {
			return "", errors.New("String cannot contain numbers")
		}

		if char == current {
			count++
		} else {
			encoded.WriteString(
				fmt.Sprintf("%v%v", count, current),
			)

			current = char
			count = 1
		}
	}

	encoded.WriteString(
		fmt.Sprintf("%v%v", count, current),
	)

	return encoded.String(), nil
}

func main() {
}
