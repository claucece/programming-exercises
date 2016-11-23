package euler

func findFirstPrimes(n int) int {
	var max int
	for i := 1; i < n; i++ {
		if n%i == 0 && i < max {
			max = i
		}
	}
	return max
}
