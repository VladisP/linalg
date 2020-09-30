package matrix

import (
	"fmt"
	"linalg/vector"
	"strconv"
)

type Matrix struct {
	RowCount    int
	ColumnCount int
	Value       [][]float64
}

func NewMatrix(value [][]float64) (*Matrix, error) {
	if !equalColumnsCount(value) {
		return nil, fmt.Errorf(ErrorColumnsNumber)
	}

	return &Matrix{
		RowCount:    len(value),
		ColumnCount: len(value[0]),
		Value:       value,
	}, nil
}

func (m *Matrix) MulMatrix(secondMatrix *Matrix) (*Matrix, error) {
	if m.ColumnCount != secondMatrix.RowCount {
		return nil, fmt.Errorf(ErrorMulMatricesSize)
	}

	res := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		res[i] = make([]float64, secondMatrix.ColumnCount)

		for k := 0; k < secondMatrix.ColumnCount; k++ {
			elem := float64(0)

			for j := 0; j < m.ColumnCount; j++ {
				elem += m.Value[i][j] * secondMatrix.Value[j][k]
			}

			res[i][k] = elem
		}
	}

	return NewMatrix(res)
}

func (m *Matrix) MulVector(v *vector.Vector) (*vector.Vector, error) {
	if m.ColumnCount != v.Size {
		return nil, fmt.Errorf(ErrorMulVectorDimension)
	}

	res := make([]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		elem := float64(0)

		for j := 0; j < m.ColumnCount; j++ {
			elem += m.Value[i][j] * v.Value[j]
		}

		res[i] = elem
	}

	return vector.NewVector(res), nil
}

func (m *Matrix) IsTriangle() bool {
	for i := 1; i < m.RowCount; i++ {
		for j := 0; j < i; j++ {
			if m.Value[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) SwapColumns(i, j int) error {
	if i < 0 || j < 0 || i >= m.ColumnCount || j >= m.ColumnCount {
		return fmt.Errorf(ErrorSwapIndexOutOfRange)
	}

	for k := 0; k < m.RowCount; k++ {
		m.Value[k][i], m.Value[k][j] = m.Value[k][j], m.Value[k][i]
	}

	return nil
}

func (m *Matrix) SwapRows(i, j int) error {
	if i < 0 || j < 0 || i >= m.RowCount || j >= m.RowCount {
		return fmt.Errorf(ErrorSwapIndexOutOfRange)
	}

	for k := 0; k < m.ColumnCount; k++ {
		m.Value[i][k], m.Value[j][k] = m.Value[j][k], m.Value[i][k]
	}

	return nil
}

func (m *Matrix) Copy() *Matrix {
	rows := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		rows[i] = make([]float64, m.ColumnCount)
		copy(rows[i], m.Value[i])
	}

	newMatrix, _ := NewMatrix(rows)

	return newMatrix
}

func (m *Matrix) String() string {
	s := ""

	for i := 0; i < m.RowCount; i++ {
		row := "|"

		for j := 0; j < m.ColumnCount; j++ {
			row += strconv.FormatFloat(m.Value[i][j], 'f', -1, 64)

			if j != m.ColumnCount-1 {
				row += "\t"
			}
		}

		row += "|\n"
		s += row
	}

	return s
}
