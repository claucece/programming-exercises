package main

import "fmt"

// common brute force one
func common_divisor(n int) int {
	result := 0
	for i := 1; i < n; i++ {
		if ((i % 3) == 0) || ((i % 5) == 0) {
			result += i
		}
	}
	return result
}

// By this: http://www.wikihow.com/Sum-the-Integers-from-1-to-N
// and https://en.wikipedia.org/wiki/Arithmetic_progression
func sumDivisbleBy(divisor, limit int) int {
	return n * (p / n) * ((p / n) + 1) / 2
}

func solve() int {
	return sumDivisbleBy(3, 999) + sumDivisbleBy(5, 999) - sumDivisbleBy(15, 999)
}

func main() {
	fmt.Println(common_divisor(1000))
	fmt.Println(solve())
}
