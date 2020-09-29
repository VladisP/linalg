package tests

import (
	"linalg/gauss"
	"linalg/matrix"
	"linalg/vector"
	"testing"
)

const Eps = 1e-6

func TestGauss(t *testing.T) {
	testMatrix, _ := matrix.NewMatrix([][]float64{
		{2, 4, 1},
		{5, 2, 1},
		{2, 3, 4},
	})
	testVector := vector.NewVector([]float64{36, 47, 37})
	ans := vector.NewVector([]float64{7, 5, 2})

	res, err := gauss.Gauss(testMatrix, testVector)
	if err != nil {
		t.Error(err)
	}

	if delta, _ := ans.Sub(res); delta.Norm() >= Eps {
		t.Errorf("Answer: expected %s, received %s", ans.String(), res.String())
	}
}

func TestGauss2(t *testing.T) {
	testMatrix, _ := matrix.NewMatrix(fillMatrix(1000))
	testVector := vector.NewVector(fillVector(1000))

	solution, err := gauss.Gauss(testMatrix, testVector)
	if err != nil {
		t.Error(err)
	}

	ans, _ := testMatrix.MulVector(solution)

	if delta, _ := ans.Sub(testVector); delta.Norm() >= Eps {
		t.Errorf("Expected %s, received %s", testVector.String(), ans.String())
		t.Errorf("Solution vector is %s", solution.String())
	}
}
