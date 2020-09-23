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
