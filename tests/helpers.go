package tests

import "math/rand"

func fillMatrix(n int) [][]float64 {
	mx := make([][]float64, n)

	for i := 0; i < n; i++ {
		mx[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			mx[i][j] = rand.Float64() * 1000
		}
	}

	return mx
}

func fillVector(n int) []float64 {
	v := make([]float64, n)

	for i := 0; i < n; i++ {
		v[i] = rand.Float64() * 1000
	}

	return v
}
