package matrix

import (
	"fmt"
	"linalg/vector"
	"math"
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

func IdentityMatrix(rowCount, columnCount int) *Matrix {
	value := make([][]float64, rowCount)

	for i := 0; i < rowCount; i++ {
		value[i] = make([]float64, columnCount)
		for j := 0; j < columnCount; j++ {
			value[i][j] = 0
		}
		value[i][i] = 1
	}

	matrix, _ := NewMatrix(value)
	return matrix
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

func (m *Matrix) MulScalar(s float64) *Matrix {
	value := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		value[i] = make([]float64, m.ColumnCount)

		for j := 0; j < m.ColumnCount; j++ {
			value[i][j] = m.Value[i][j] * s
		}
	}

	newM, _ := NewMatrix(value)
	return newM
}

func (m *Matrix) Sum(matrix *Matrix) (*Matrix, error) {
	if m.RowCount != matrix.RowCount || m.ColumnCount != matrix.ColumnCount {
		return nil, fmt.Errorf(ErrorDimensions)
	}

	res := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		res[i] = make([]float64, m.ColumnCount)

		for j := 0; j < m.ColumnCount; j++ {
			res[i][j] = m.Value[i][j] + matrix.Value[i][j]
		}
	}

	return NewMatrix(res)
}

func (m *Matrix) Sub(matrix *Matrix) (*Matrix, error) {
	return m.Sum(matrix.MulScalar(-1))
}

func (m *Matrix) Inverse2() (*Matrix, error) {
	if m.ColumnCount != 2 || m.RowCount != 2 {
		return nil, fmt.Errorf(ErrorMatrix2by2)
	}

	a, b, c, d := m.Value[0][0], m.Value[0][1], m.Value[1][0], m.Value[1][1]
	factor := a*d - b*c

	value := [][]float64{
		{d / factor, -b / factor},
		{-c / factor, a / factor},
	}

	return NewMatrix(value)
}

func (m *Matrix) ConditionNumber2() (float64, error) {
	if m.ColumnCount != 2 || m.RowCount != 2 {
		return 0, fmt.Errorf(ErrorMatrix2by2)
	}
	i, _ := m.Inverse2()
	return i.UniformNorm() * m.UniformNorm(), nil
}

func (m *Matrix) UniformNorm() float64 {
	norm := math.Abs(m.Value[0][0])

	for i := 0; i < m.RowCount; i++ {
		sum := float64(0)

		for j := 0; j < m.ColumnCount; j++ {
			sum += math.Abs(m.Value[i][j])
		}

		if sum > norm {
			norm = sum
		}
	}

	return norm
}

func (m *Matrix) Norm() float64 {
	res := float64(0)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			res += m.Value[i][j] * m.Value[i][j]
		}
	}

	return math.Sqrt(res)
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

func (m *Matrix) IsDominance() bool {
	for i := 0; i < m.RowCount; i++ {
		diag := math.Abs(m.Value[i][i])
		sum := float64(0)

		for j := 0; j < m.ColumnCount; j++ {
			if i != j {
				sum += math.Abs(m.Value[i][j])
			}
		}

		if diag <= sum {
			return false
		}
	}

	return true
}

func (m *Matrix) IsSymmetric() (bool, error) {
	if m.RowCount != m.ColumnCount {
		return false, fmt.Errorf(ErrorSquareMatrix)
	}

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Value[i][j] != m.Value[j][i] {
				return false, nil
			}
		}
	}

	return true, nil
}

func (m *Matrix) IsPositive() bool {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Value[i][j] <= 0 {
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

func (m *Matrix) Flat() []float64 {
	res := make([]float64, m.RowCount*m.ColumnCount)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			res[i*m.ColumnCount+j] = m.Value[i][j]
		}
	}

	return res
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
