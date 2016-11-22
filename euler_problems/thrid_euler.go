package main

import "fmt"

func findPrimes(n int) int {
	var max int
	for i := 1; i < n; i++ {
		if n%i == 0 && i < max {
			max = i
		}
	}
	return max
}

func main() {
	fmt.Println(findPrimes(10))
}
