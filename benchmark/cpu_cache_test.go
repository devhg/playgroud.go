package benchmark

import (
	"testing"
)

func createMatrix(size int) [][]int64 {
	matrix := make([][]int64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int64, size)
	}
	return matrix
}

const matrixLength = 6400

func BenchmarkMatrixCombination(b *testing.B) {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[i][j]
			}
		}
	}
}

func BenchmarkMatrixReversedCombination(b *testing.B) {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[j][i]
			}
		}
	}
}
