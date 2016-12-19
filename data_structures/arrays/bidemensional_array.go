package main

// TODO: manually test
// TODO: use the struct

type bidimensionalArray struct {
	rowLength     int // has to be 2
	colLength     int
	zerothElement *int
}

func (b *bidimensionalArray) makeBiArray() [][]int {
	array := make([][]int, b.colLength)
	for i := range array {
		array[i] = make([]int, b.rowLength)
	}
	return array
}

func (b *bidimensionalArray) makeBiArrayByData() (s [][]int) {
	raw := make([]int, b.colLength*b.rowLength)
	for i := range raw {
		raw[i] = i
	}

	s = make([][]int, b.colLength)
	for j := range s {
		s[j] = raw[j*b.rowLength : j*b.rowLength+b.rowLength]
	}

	return
}

func readBidimensionalArray(array [][]int) (r int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			r = array[i][j]
		}
	}
	return
}

// refactor this
func writeInBidimensionalArray(array [][]int, a, x, y int) (s [][]int) {
	if array[x][y] == 0 {
		s[x][y] = a
	}
	return s
}
