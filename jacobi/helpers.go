package jacobi

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
	if !a.IsDominance() {
		return fmt.Errorf(ErrorDominanceMatrix)
	}
	return nil
}
