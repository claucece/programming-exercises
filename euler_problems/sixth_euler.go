package euler

const N = 100

// by brute force
func sumSquares() int {
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i * i
	}
	return sum
}

func squareSum() int {
	square := 0
	for i := 1; i <= N; i++ {
		square += i
	}
	return square * square
}

func difference() int {
	return squareSum() - sumSquares()
}

// by Gauss formula and extended
func calculateDifferenceOfSquares() int {
	sum := N * (N + 1) / 2
	square := (N * (N + 1) * (2*N + 1)) / 6
	difference := sum*sum - square
	return difference
}
