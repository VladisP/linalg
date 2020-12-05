package one_parameter

import (
	"linalg/matrix"
	"linalg/vector"
)

const Eps = 1e-12

func OneParameter(a *matrix.Matrix, f *vector.Vector, tau float64) (*vector.Vector, int, error) {
	if err := checkConstraints(a, f); err != nil {
		return nil, 0, err
	}

	id := matrix.IdentityMatrix(a.RowCount, a.ColumnCount)
	p := id.Sub(a.MulScalar(tau))
	g := f.MulScalar(tau)
	x := vector.EmptyVector(f.Size)
	xNext := vector.EmptyVector(f.Size)
	iterCount := 0

	for {
		iterCount++

		px, _ := p.MulVector(x)
		xNext, _ = px.Sum(g)

		if delta, _ := x.Sub(xNext); delta.UniformNorm() < Eps {
			return xNext, iterCount, nil
		}

		x = vector.Of(xNext)
	}
}
