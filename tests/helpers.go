package tests

import (
	"linalg/matrix"
	"linalg/vector"
	"math/rand"
)

func fillRandomMatrix(n int) (*matrix.Matrix, error) {
	value := make([][]float64, n)

	for i := 0; i < n; i++ {
		value[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			value[i][j] = rand.Float64() * 1000
		}
	}

	return matrix.NewMatrix(value)
}

func fillDominanceMatrix(n int) (*matrix.Matrix, error) {
	m, err := fillRandomMatrix(n)
	if err != nil {
		return nil, err
	}

	for i := 0; i < m.RowCount; i++ {
		m.Value[i][i] = sum(m.Value[i]) * 100
	}

	return m, nil
}

func sum(s []float64) float64 {
	sum := float64(0)

	for _, v := range s {
		sum += v
	}

	return sum
}

func fillRandomVector(n int) *vector.Vector {
	v := make([]float64, n)

	for i := 0; i < n; i++ {
		v[i] = rand.Float64() * 1000
	}

	return vector.NewVector(v)
}
