package gauss

import (
	"fmt"
	"linalg/matrix"
	"linalg/vector"
)

func triangularMatrix(a *matrix.Matrix, f *vector.Vector) (*matrix.Matrix, *vector.Vector, error) {
	if a.RowCount != a.ColumnCount {
		return nil, nil, fmt.Errorf(ErrorSquareMatrix)
	}
	if a.RowCount != f.Size {
		return nil, nil, fmt.Errorf(ErrorVectorDimension)
	}
	if nullDiagonal(a) {
		return nil, nil, fmt.Errorf(ErrorNullDiagonal)
	}

	newMatrixValue := copyMatrix(a)
	newVectorValue := copyVector(f)

	for eliminateIndex := 0; eliminateIndex < a.RowCount; eliminateIndex++ {
		divider := newMatrixValue[eliminateIndex][eliminateIndex]

		for j := 0; j < a.ColumnCount; j++ {
			newMatrixValue[eliminateIndex][j] /= divider
		}

		newVectorValue[eliminateIndex] /= divider

		for i := eliminateIndex + 1; i < a.RowCount; i++ {
			factor := newMatrixValue[i][eliminateIndex]

			for j := 0; j < a.ColumnCount; j++ {
				newMatrixValue[i][j] -= newMatrixValue[eliminateIndex][j] * factor
			}

			newVectorValue[i] -= newVectorValue[eliminateIndex] * factor
		}
	}

	newMatrix, err := matrix.NewMatrix(newMatrixValue)
	if err != nil {
		return nil, nil, err
	}
	newVector := vector.NewVector(newVectorValue)

	return newMatrix, newVector, nil
}
