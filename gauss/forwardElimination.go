package gauss

import (
	"fmt"
	"linalg/matrix"
	"linalg/vector"
)

func TriangularMatrix(a *matrix.Matrix, f *vector.Vector) (*matrix.Matrix, *vector.Vector, error) {
	if a.RowCount != a.ColumnCount {
		return nil, nil, fmt.Errorf(ErrorSquareMatrix)
	}
	if a.RowCount != f.Size {
		return nil, nil, fmt.Errorf(ErrorVectorDimension)
	}

	newMatrixValue := copyMatrix(a)
	newVectorValue := copyVector(f)

	for eliminateIndex := 0; eliminateIndex < a.RowCount-1; eliminateIndex++ {
		for i := eliminateIndex + 1; i < a.RowCount; i++ {
			factor1, factor2 := newMatrixValue[i][eliminateIndex], newMatrixValue[eliminateIndex][eliminateIndex]

			for j := 0; j < a.ColumnCount; j++ {
				newMatrixValue[i][j] = newMatrixValue[i][j]*factor2 - newMatrixValue[eliminateIndex][j]*factor1
			}

			newVectorValue[i] = newVectorValue[i]*factor2 - newVectorValue[eliminateIndex]*factor1
		}
	}

	newMatrix, err := matrix.NewMatrix(newMatrixValue)
	if err != nil {
		return nil, nil, err
	}
	newVector := vector.NewVector(newVectorValue)

	return newMatrix, newVector, nil
}
