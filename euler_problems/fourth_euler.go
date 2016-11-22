package main

import "fmt"

func findLargestFactor(num int) int {
	counter := 2
	var largest_fact int
	for counter*counter <= num {
		if num%counter == 0 {
			num /= counter
			largest_fact = counter
		} else {
			counter++
		}
	}
	if num > 0 {
		largest_fact = num
		return largest_fact
	}
	return largest_fact
}

func main() {
	fmt.Println(findLargestFactor(600851475143))
}
