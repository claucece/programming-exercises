package main

// TODO: manually test
// TODO: use the struct

type bidimensionalArray struct {
	rowLength     int
	colLength     int
	zerothElement *int
}

func (b *bidimensionalArray) makeBiArray() [][]int {
	array := make([][]int, b.colLength, b.rowLength)
	return array
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
