package main

import (
	"strings"
)

func findAverage(global_total []int) int {
	var total int
	for _, value := range global_total {
		total += value
	}
	return total / len(global_total)
}

func findString(str_chain, letter string) int {
	return strings.Count(str_chain, letter)
}
