package kholez

import (
	"linalg/matrix"
	"math"
)

func Kholez(a *matrix.Matrix) (*matrix.Matrix, error) {
	if err := checkConstraints(a); err != nil {
		return nil, err
	}

	l := matrix.IdentityMatrix(a.RowCount, a.ColumnCount)

	l.Value[0][0] = math.Sqrt(a.Value[0][0])

	for j := 1; j < a.RowCount; j++ {
		l.Value[j][0] = a.Value[j][0] / l.Value[0][0]
	}

	for i := 1; i < a.RowCount; i++ {
		sum := float64(0)
		for p := 0; p < i; p++ {
			sum += l.Value[i][p] * l.Value[i][p]
		}
		l.Value[i][i] = math.Sqrt(a.Value[i][i] - sum)
	}

	for i := 1; i < a.RowCount-1; i++ {
		for j := i + 1; j < a.ColumnCount; j++ {
			sum := float64(0)
			for p := 0; p < i; p++ {
				sum += l.Value[i][p] * l.Value[j][p]
			}
			l.Value[j][i] = (a.Value[j][i] - sum) / l.Value[i][i]
		}
	}

	return l, nil
}
