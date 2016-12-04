//TODO : this is still to mathematical

package main

import "fmt"

var n []int
var sz []int

func setID(s int) []int {
	n := make([]int, s)
	sz := make([]int, s)
	for i := 0; i < s; i++ {
		n[i] = i
		sz[i] = 1
	}
	return n
}

func root(i int) int {
	for i != n[i] {
		n[i] = n[n[i]]
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
	if i == j {
		return n
	}
	if sz[i] < sz[j] {
		n[i] = j
		sz[j] += sz[i]
	} else {
		n[j] = i
		sz[i] += sz[j]
	}
	return n
}

func main() {
	fmt.Println(len(n))
}
