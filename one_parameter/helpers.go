package one_parameter

import (
	"fmt"
	"linalg/matrix"
	"linalg/vector"
)

func checkConstraints(a *matrix.Matrix, f *vector.Vector) error {
	if a.RowCount != a.ColumnCount {
		return fmt.Errorf(ErrorSquareMatrix)
	}
	if a.RowCount != f.Size {
		return fmt.Errorf(ErrorVectorDimension)
	}
	if isSymmetric, err := a.IsSymmetric(); err != nil || !isSymmetric {
		return fmt.Errorf(ErrorSymmetricMatrix)
	}
	if !a.IsPositive() {
		return fmt.Errorf(ErrorPositiveMatrix)
	}
	return nil
}
