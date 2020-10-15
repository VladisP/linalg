package gauss

import (
	"fmt"
	"linalg/matrix"
	"linalg/permutation"
	"linalg/vector"
	"math"
)

func checkConstraints(a *matrix.Matrix, f *vector.Vector) error {
	if a.RowCount != a.ColumnCount {
		return fmt.Errorf(ErrorSquareMatrix)
	}
	if a.RowCount != f.Size {
		return fmt.Errorf(ErrorVectorDimension)
	}
	return nil
}

func eliminate(eliminateIndex int, a *matrix.Matrix, f *vector.Vector) error {
	if a.Value[eliminateIndex][eliminateIndex] == 0 {
		return fmt.Errorf(ErrorNullDiagonal)
	}

	divider := a.Value[eliminateIndex][eliminateIndex]

	for j := 0; j < a.ColumnCount; j++ {
		a.Value[eliminateIndex][j] /= divider
	}

	f.Value[eliminateIndex] /= divider

	for i := eliminateIndex + 1; i < a.RowCount; i++ {
		factor := a.Value[i][eliminateIndex]

		for j := 0; j < a.ColumnCount; j++ {
			a.Value[i][j] -= a.Value[eliminateIndex][j] * factor
		}

		f.Value[i] -= f.Value[eliminateIndex] * factor
	}

	return nil
}

func substitution(a *matrix.Matrix, f *vector.Vector) (*vector.Vector, error) {
	if !a.IsTriangle() {
		return nil, fmt.Errorf(ErrorTriangleMatrix)
	}

	res := make([]float64, f.Size)

	for i := a.RowCount - 1; i >= 0; i-- {
		x := f.Value[i]

		for j := i + 1; j < a.ColumnCount; j++ {
			x -= a.Value[i][j] * res[j]
		}

		res[i] = x
	}

	return vector.NewVector(res), nil
}

func rowPivoting(step int, a *matrix.Matrix, p permutation.Permutation) error {
	max := math.Abs(a.Value[step][step])
	maxIndex := step

	for j := step; j < a.ColumnCount; j++ {
		elem := math.Abs(a.Value[step][j])

		if elem > max {
			max, maxIndex = elem, j
		}
	}

	if err := a.SwapColumns(step, maxIndex); err != nil {
		return err
	}

	tr := permutation.Id(len(p))
	if err := tr.Swap(step, maxIndex); err != nil {
		return err
	}

	return p.Compose(tr)
}

func columnPivoting(step int, a *matrix.Matrix, f *vector.Vector) error {
	max := math.Abs(a.Value[step][step])
	maxIndex := step

	for i := step; i < a.RowCount; i++ {
		elem := math.Abs(a.Value[i][step])

		if elem > max {
			max, maxIndex = elem, i
		}
	}

	if err := f.SwapComponents(step, maxIndex); err != nil {
		return err
	}
	return a.SwapRows(step, maxIndex)
}
