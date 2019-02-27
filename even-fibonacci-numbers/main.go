package main

func FindSumOfEvens(max int) int {
	sum := 0
	current := 0
	prev := 0
	next := 1

	for current < max {
		current = prev + next
		prev = next
		next = current

		if current%2 == 0 {
			sum += current
		}
	}

	return sum
}

func main() {
}
