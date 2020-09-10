package gauss

import (
	"fmt"
	"linalg/matrix"
	"linalg/vector"
)

func Substitution(a *matrix.Matrix, f *vector.Vector) (*vector.Vector, error) {
	if a.RowCount != a.ColumnCount {
		return nil, fmt.Errorf(ErrorSquareMatrix)
	}
	if a.RowCount != f.Size {
		return nil, fmt.Errorf(ErrorVectorDimension)
	}
	if !isTriangle(a) {
		return nil, fmt.Errorf(ErrorTriangleMatrix)
	}

	res := make([]float64, f.Size)

	for i := a.RowCount - 1; i >= 0; i-- {
		x := f.Value[i]

		for j := i + 1; j < a.ColumnCount; j++ {
			x -= a.Value[i][j] * res[j]
		}

		res[i] = x / a.Value[i][i]
	}

	return vector.NewVector(res), nil
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
