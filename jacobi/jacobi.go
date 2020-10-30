package jacobi

import (
	"linalg/matrix"
	"linalg/vector"
)

const Eps = 1e-12

func Jacobi(a *matrix.Matrix, f *vector.Vector) (*vector.Vector, int, error) {
	if err := checkConstraints(a, f); err != nil {
		return nil, 0, err
	}

	x := vector.EmptyVector(f.Size)
	xNext := vector.EmptyVector(f.Size)
	iterCount := 0

	for {
		iterCount++

		for i := 0; i < a.RowCount; i++ {
			xNext.Value[i] = f.Value[i]
			for j := 0; j < a.ColumnCount; j++ {
				if i != j {
					xNext.Value[i] -= a.Value[i][j] * x.Value[j]
				}
			}
			xNext.Value[i] /= a.Value[i][i]
		}

		if delta, _ := x.Sub(xNext); delta.UniformNorm() < Eps {
			return xNext, iterCount, nil
		}

		x = vector.Of(xNext)
	}
}
