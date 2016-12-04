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

func isConnected(p, q int) bool {
	return n[p] == n[q]
}

func union(p, q int) []int {
	id_p := n[p]
	id_q := n[q]
	for _, i := range n {
		if n[i] == id_p {
			n[i] = id_q
		}
	}
	return n
}

func main() {
	fmt.Println(len(n))
}
