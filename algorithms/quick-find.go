package main

func setID(s int) []int {
	n := make([]int, s)
	for i := 0; i < s; i++ {
		n[i] = i
	}
	return n
}

func isConnected(p, q, c int, n []int) bool {
	n = setID(c)
	return n[p] == n[q]
}

func union(p, q, c int, n []int) []int {
	n = setID(c)
	idP := n[p]
	idQ := n[q]
	for _, i := range n {
		if n[i] == idP {
			n[i] = idQ
		}
	}
	return n
}
