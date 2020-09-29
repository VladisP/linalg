package gauss

import (
	"linalg/matrix"
	"linalg/vector"
)

func Gauss(a *matrix.Matrix, f *vector.Vector) (*vector.Vector, error) {
	trMatrix, trVector, err := triangularMatrix(a, f)
	if err != nil {
		return nil, err
	}

	result, err := substitution(trMatrix, trVector)
	return result, err
}
