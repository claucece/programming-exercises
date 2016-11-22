package main

import "fmt"

// brute force
func sumSquares(n int) int {
	var sum int
	for n := 1; n <= 100; n++ {
		sum += n * n
	}
	return sum
}

func squareSum(n int) int {
	var square int
	for n := 1; n <= 100; n++ {
		square += n
	}
	return square * square
}

func difference(n, x int) int {
	return squareSum(n) - sumSquares(x)
}

func main() {
	fmt.Println(difference(100, 100))
}
