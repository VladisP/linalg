package kholez

import (
	"fmt"
	"linalg/matrix"
)

func checkConstraints(a *matrix.Matrix) error {
	if isSymmetric, err := a.IsSymmetric(); err != nil || !isSymmetric {
		return fmt.Errorf(ErrorSymmetricMatrix)
	}
	if !a.IsPositive() {
		return fmt.Errorf(ErrorPositiveMatrix)
	}
	if !a.IsDominance() {
		return fmt.Errorf(ErrorDominanceMatrix)
	}
	return nil
}
