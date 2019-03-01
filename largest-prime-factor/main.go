package main

func LargestPrimeFactor(prime int64) int64 {
	var pf int64 = 2

	for prime > 1 {
		for prime%pf != 0 {
			pf += 1
		}

		prime /= pf
	}

	return pf
}

func main() {
}
