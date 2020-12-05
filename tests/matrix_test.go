package tests

import (
	"linalg/matrix"
	"linalg/vector"
	"testing"
)

func TestMulMatrix(t *testing.T) {
	m1, _ := matrix.NewMatrix([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	m2, _ := matrix.NewMatrix([][]float64{
		{1, 3, 5},
		{2, 4, 6},
	})
	ans, _ := matrix.NewMatrix([][]float64{
		{5, 11, 17},
		{11, 25, 39},
		{17, 39, 61},
	})

	m, err := m1.MulMatrix(m2)

	if err != nil {
		t.Error(err)
	}

	if m.String() != ans.String() {
		t.Errorf("Expected %s, received %s", ans.String(), m.String())
	}
}

func TestMulVector(t *testing.T) {
	m, _ := matrix.NewMatrix([][]float64{
		{1, 3, 5},
		{2, 4, 6},
	})
	v := vector.NewVector([]float64{4, 5, 6})
	ans := vector.NewVector([]float64{49, 64})

	res, err := m.MulVector(v)

	if err != nil {
		t.Error(err)
	}

	if ans.String() != res.String() {
		t.Errorf("Expected %s, received %s", ans.String(), res.String())
	}
}

func TestNearestPower2(t *testing.T) {
	testValues := []int{1, 2, 3, 7, 8, 9, 22}
	expectedValues := []int{1, 2, 4, 8, 8, 16, 32}

	for i, v := range testValues {
		if np := matrix.NearestPower2(v); np != expectedValues[i] {
			t.Errorf("Error: expected %d, received %d\n", expectedValues[i], np)
		}
	}
}

func TestSlice(t *testing.T) {
	m, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	})
	u := matrix.SliceRange{Start: 0, End: 3}
	v := matrix.SliceRange{Start: 3, End: 6}

	slice1 := m.Slice(u, u)
	slice2 := m.Slice(u, v)
	slice3 := m.Slice(v, u)
	slice4 := m.Slice(v, v)

	res1, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3},
		{7, 8, 9},
		{13, 14, 15},
	})
	res2, _ := matrix.NewMatrix([][]float64{
		{4, 5, 6},
		{10, 11, 12},
		{16, 17, 18},
	})
	res3, _ := matrix.NewMatrix([][]float64{
		{19, 20, 21},
		{25, 26, 27},
		{31, 32, 33},
	})
	res4, _ := matrix.NewMatrix([][]float64{
		{22, 23, 24},
		{28, 29, 30},
		{34, 35, 36},
	})

	if slice1.String() != res1.String() {
		t.Errorf("Error slice1:\n %s\n", slice1.String())
	}
	if slice2.String() != res2.String() {
		t.Errorf("Error slice2:\n %s\n", slice2.String())
	}
	if slice3.String() != res3.String() {
		t.Errorf("Error slice3:\n %s\n", slice3.String())
	}
	if slice4.String() != res4.String() {
		t.Errorf("Error slice4:\n %s\n", slice4.String())
	}
}

func TestCompleteToSquare(t *testing.T) {
	m, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	c := m.CompleteToSquare(matrix.NearestPower2(m.ColumnCount))

	res, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3, 0},
		{4, 5, 6, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	if res.String() != c.String() {
		t.Errorf("Error:\n %s\n", c.String())
	}
}

func TestCopyToSlice(t *testing.T) {
	m := matrix.IdentityMatrix(6, 6)
	u := matrix.SliceRange{Start: 0, End: 3}
	v := matrix.SliceRange{Start: 3, End: 6}
	res, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	})

	m1, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3},
		{7, 8, 9},
		{13, 14, 15},
	})
	m2, _ := matrix.NewMatrix([][]float64{
		{4, 5, 6},
		{10, 11, 12},
		{16, 17, 18},
	})
	m3, _ := matrix.NewMatrix([][]float64{
		{19, 20, 21},
		{25, 26, 27},
		{31, 32, 33},
	})
	m4, _ := matrix.NewMatrix([][]float64{
		{22, 23, 24},
		{28, 29, 30},
		{34, 35, 36},
	})

	matrix.CopyToSlice(m.Slice(u, u), m1)
	matrix.CopyToSlice(m.Slice(u, v), m2)
	matrix.CopyToSlice(m.Slice(v, u), m3)
	matrix.CopyToSlice(m.Slice(v, v), m4)

	if m.String() != res.String() {
		t.Errorf("Error:\n %s\n", m.String())
	}
}
