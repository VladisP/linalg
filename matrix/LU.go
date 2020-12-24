package matrix

import (
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack/lapack64"
	"gonum.org/v1/gonum/mat"
	"linalg/vector"
)

func (m *Matrix) LU() (*Matrix, *Matrix) {
	L := IdentityMatrix(m.RowCount, m.ColumnCount)
	U := IdentityMatrix(m.RowCount, m.ColumnCount)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			U.Value[i][j] = 0
			L.Value[i][j] = 0
		}
		L.Value[i][i] = 1
	}

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if i <= j {
				sum := 0.
				for k := 0; k <= i-1; k++ {
					sum += L.Value[i][k] * U.Value[k][j]
				}
				U.Value[i][j] = m.Value[i][j] - sum
			} else {
				sum := 0.
				for k := 0; k <= j-1; k++ {
					sum += L.Value[i][k] * U.Value[k][j]
				}
				L.Value[i][j] = (m.Value[i][j] - sum) / U.Value[j][j]
			}
		}
	}

	return L, U
}

func SolveSystemLU(m *Matrix, f *vector.Vector) *vector.Vector {
	L, U := m.LU()
	y := vector.EmptyVector(f.Size)
	x := vector.EmptyVector(f.Size)

	for i := 0; i < y.Size; i++ {
		sum := 0.
		for k := 0; k <= i-1; k++ {
			sum += L.Value[i][k] * y.Value[k]
		}
		y.Value[i] = f.Value[i] - sum
	}

	for i := x.Size - 1; i >= 0; i-- {
		sum := 0.
		for k := i + 1; k < x.Size; k++ {
			sum += U.Value[i][k] * x.Value[k]
		}
		x.Value[i] = (y.Value[i] - sum) / U.Value[i][i]
	}

	return x
}

func SolveSystemGonum(m *Matrix, f *vector.Vector) *vector.Vector {
	libA := mat.NewDense(m.RowCount, m.ColumnCount, m.Flat())
	libF := mat.NewVecDense(len(f.Value), vector.Of(f).Value)
	rawX2 := mat.NewVecDense(len(f.Value), nil)
	_ = rawX2.SolveVec(libA, libF)
	return vector.NewVector(rawX2.RawVector().Data)
}

func (m *Matrix) DetLU() float64 {
	L, U := m.LU()
	det := 1.

	for i := 0; i < m.RowCount; i++ {
		det *= L.Value[i][i] * U.Value[i][i]
	}

	return det
}

func InverseGonum(m *Matrix) (*Matrix, error) {
	libM := blas64.General{
		Rows:   m.RowCount,
		Cols:   m.ColumnCount,
		Data:   m.Flat(),
		Stride: m.ColumnCount,
	}

	v := make([]int, m.RowCount)
	_ = lapack64.Getrf(libM, v)
	work := make([]float64, 4*m.RowCount)
	lapack64.Getri(libM, v, work, -1)

	if int(work[0]) > 4*m.RowCount {
		work = make([]float64, int(work[0]))
	} else {
		work = work[:4*m.RowCount]
	}

	lapack64.Getri(libM, v, work, len(work))

	data := make([][]float64, libM.Rows)
	for i := 0; i < libM.Rows; i++ {
		data[i] = make([]float64, libM.Cols)
		for j := 0; j < libM.Cols; j++ {
			data[i][j] = libM.Data[i*libM.Cols+j]
		}
	}
	return NewMatrix(data)
}

func (m *Matrix) InverseLU() *Matrix {
	columns := IdentityMatrix(m.RowCount, m.ColumnCount).GetColumns()
	inverseColumns := make([]*vector.Vector, 0, m.RowCount)

	for _, f := range columns {
		res := SolveSystemLU(m, f)
		inverseColumns = append(inverseColumns, res)
	}

	return FromColumns(inverseColumns)
}
