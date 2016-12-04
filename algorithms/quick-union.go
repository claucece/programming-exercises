package main

import "fmt"

var n []int

func setID(s int) []int {
	n := make([]int, s)
	for i := 0; i < s; i++ {
		n[i] = i
	}
	return n
}

func root(i int) int {
	for i != n[i] {
		i = n[i]
	}
	return i
}

func isConnected(p, q int) bool {
	return root(p) == root(q)
}

func union(p, q int) []int {
	i := root(p)
	j := root(q)
	n[i] = j
	return n
}

func main() {
	fmt.Println(len(n))
}
