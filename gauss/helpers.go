package gauss

import (
	"linalg/matrix"
	"linalg/vector"
)

func copyMatrix(src *matrix.Matrix) [][]float64 {
	dst := make([][]float64, src.RowCount)

	for i := 0; i < src.RowCount; i++ {
		dst[i] = make([]float64, src.ColumnCount)
		copy(dst[i], src.Value[i])
	}

	return dst
}

func copyVector(src *vector.Vector) []float64 {
	dst := make([]float64, src.Size)
	copy(dst, src.Value)

	return dst
}

func isTriangle(m *matrix.Matrix) bool {
	for i := 1; i < m.RowCount; i++ {
		for j := 0; j < i; j++ {
			if m.Value[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func nullDiagonal(m *matrix.Matrix) bool {
	for i := 0; i < m.RowCount; i++ {
		if m.Value[i][i] == 0 {
			return true
		}
	}

	return false
}
