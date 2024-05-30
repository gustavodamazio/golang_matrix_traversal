package main

import (
	"time"
)

const N = 5000

func main() {
	// CreateAndSaveMatrix(N)

	matrix := ReadMatrix(N)

	sumRowMajor, durationRowMajor := sumRowMajor(matrix, N)

	sumColumnMajor, durationColumnMajor := sumColumnMajor(matrix, N)

	PrintResult(durationColumnMajor, durationRowMajor, sumColumnMajor, sumRowMajor)
}

func sumRowMajor(matrix [][]int, N int) (int, time.Duration) {
	start := time.Now()
	sum := 0
	for row := 0; row < N; row++ {
		for column := 0; column < N; column++ {
			sum += matrix[row][column]
		}
	}
	return sum, time.Since(start)
}

func sumColumnMajor(matrix [][]int, N int) (int, time.Duration) {
	start := time.Now()
	sum := 0
	for row := 0; row < N; row++ {
		for column := 0; column < N; column++ {
			sum += matrix[column][row]
		}
	}
	return sum, time.Since(start)
}
