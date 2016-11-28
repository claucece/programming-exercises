package main

type bidemensionalArray struct {
	Row_Length    int
	Col_Length    int
	ZerothElement *int
}

func MakeBiArray() [][]int {
	array := make([]int, bidemensionalArray.Row_Length)
	for _, i := range array {
		array[i] = make([]int, bidemensionalArray.Col_Length)
	}
	return array
}
