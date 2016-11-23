package euler

func findFirstPrimes(n int) int {
	var lrgs_fctr int
	var isPrime bool
	for i := 2; i < n; i++ {
		isPrime = true
		for j := 2; j < i; j++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			lrgs_fctr = i
		}
	}
	return lrgs_fctr
}
