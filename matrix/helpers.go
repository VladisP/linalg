package matrix

type SliceRange struct {
	Start int
	End   int
}

func (m *Matrix) Slice(rowRange, colRange SliceRange) *Matrix {
	rows := m.Value[rowRange.Start:rowRange.End]
	s := make([][]float64, len(rows))

	for i, r := range rows {
		s[i] = r[colRange.Start:colRange.End]
	}

	newMatrix, _ := NewMatrix(s)
	return newMatrix
}

func NearestPower2(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return x + 1
}

func (m *Matrix) CompleteToSquare(size int) *Matrix {
	value := make([][]float64, size)

	for i := 0; i < size; i++ {
		value[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			if i < m.RowCount && j < m.ColumnCount {
				value[i][j] = m.Value[i][j]
			} else {
				value[i][j] = 0
			}
		}
	}

	newMatrix, _ := NewMatrix(value)
	return newMatrix
}

func CopyToSlice(slice *Matrix, src *Matrix) {
	for i := 0; i < slice.RowCount; i++ {
		for j := 0; j < slice.ColumnCount; j++ {
			slice.Value[i][j] = src.Value[i][j]
		}
	}
}

func equalColumnsCount(rows [][]float64) bool {
	for i := 0; i < len(rows)-1; i++ {
		if len(rows[i]) != len(rows[i+1]) {
			return false
		}
	}

	return true
}
