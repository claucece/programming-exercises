package main

// TODO: manually test
// TODO: use the struct

type bidemensionalArray struct {
	Row_Length    int
	Col_Length    int
	ZerothElement *int
}

func (b *bidemensionalArray) MakeBiArray() [][]int {
	array := make([]int, bidemensionalArray.Row_Length)
	for _, i := range array {
		array[i] = make([]int, bidemensionalArray.Col_Length)
	}
	return array
}

func read(array [][]int) int {
	var r int
	for _, i := range array {
		for _, j := range array {
			r = array[i][j]
			return r
		}
	}
	return nil
}

func write(array [][]int, a, b, x, y int) [][]int {
	if array[x] == 0 || array[x] == nil {
		if array[y] == 0 || array[y] == nil {
			s[x][y] = a, b
			return s
		}
	}
	return 0
}
