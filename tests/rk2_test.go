package tests

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"linalg/gauss"
	"linalg/matrix"
	"linalg/vector"
	"testing"
)

func TestRK2(t *testing.T) {
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := vector.NewVector([]float64{0, 1, 2, 3, 0, 5, 6, 7, 8, 9, 10})
	m := createSystemMatrix(x)

	a1, _ := gauss.ClassicGauss(m, f)
	fn1 := createFn(x, a1.Value)

	a2, _ := gauss.PivotingGauss(m, f)
	fn2 := createFn(x, a2.Value)

	libM := mat.NewDense(m.RowCount, m.ColumnCount, m.Flat())
	libF := mat.NewVecDense(len(f.Value), vector.Of(f).Value)
	a3 := mat.NewVecDense(len(f.Value), nil)
	_ = a3.SolveVec(libM, libF)
	fn3 := createFn(x, a3.RawVector().Data)

	fmt.Println(fn1)
	fmt.Println(fn2)
	fmt.Println(fn3)
	for i, v := range x {
		fmt.Printf("(%.15f,%.15f)\n", v, f.Value[i])
	}
}

func createFn(x, val []float64) string {
	fn := ""
	for i, v := range val {
		fn += fmt.Sprintf("%.15f", v)
		for j := 0; j < i; j++ {
			fn += fmt.Sprintf("(x-%.15f)", x[j])
		}
		if i < len(val)-1 {
			fn += "+"
		}
	}
	return fn
}

func createSystemMatrix(x []float64) *matrix.Matrix {
	m := matrix.IdentityMatrix(len(x), len(x))

	for i := 0; i < m.RowCount; i++ {
		m.Value[i][0] = 1
	}

	for i := 1; i < m.RowCount; i++ {
		for j := 1; j <= i; j++ {
			mul := float64(1)
			for k := 0; k < j; k++ {
				mul *= x[i] - x[k]
			}
			m.Value[i][j] = mul
		}
	}

	return m
}
