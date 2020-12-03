package one_parameter

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"linalg/matrix"
)

func GetEigen(m *matrix.Matrix) (max float64, min float64, opt float64, err error) {
	eigenvalues, err := getEigen(m)
	if err != nil {
		return 0, 0, 0, err
	}

	max, min = eigenvalues[0], eigenvalues[0]
	for _, v := range eigenvalues {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min, 2 / (min + max), nil
}

func getEigen(m *matrix.Matrix) ([]float64, error) {
	var eigen mat.Eigen

	dm := mat.NewDense(m.RowCount, m.ColumnCount, m.Flat())
	eigen.Factorize(dm, mat.EigenLeft)
	rawEigenValues := eigen.Values(nil)

	res := make([]float64, len(rawEigenValues))
	for i, v := range rawEigenValues {
		res[i] = real(v)
		if imag(v) != 0 {
			return nil, fmt.Errorf(ErrorComplex)
		}
	}
	return res, nil
}
