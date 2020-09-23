package tests

import (
	"linalg/gauss"
	"linalg/matrix"
	"linalg/vector"
	"testing"
)

func TestGauss(t *testing.T) {
	testMatrix, _ := matrix.NewMatrix([][]float64{
		{2, 4, 1},
		{5, 2, 1},
		{2, 3, 4},
	})
	testVector := vector.NewVector([]float64{36, 47, 37})

	triangMatrix, _ := matrix.NewMatrix([][]float64{
		{2, 4, 1},
		{0, -16, -3},
		{0, 0, -102},
	})
	triangVector := vector.NewVector([]float64{36, -86, -204})
	ans := vector.NewVector([]float64{7, 5, 2})

	a, f, err := gauss.TriangularMatrix(testMatrix, testVector)

	if err != nil {
		t.Error(err)
	}
	if a.String() != triangMatrix.String() {
		t.Errorf("TriangMatrix: expected %s, received %s", triangMatrix.String(), a.String())
	}
	if f.String() != triangVector.String() {
		t.Errorf("TriangVector: expected %s, received %s", triangVector.String(), f.String())
	}

	res, err := gauss.Substitution(a, f)

	if err != nil {
		t.Error(err)
	}
	if ans.String() != res.String() {
		t.Errorf("Answer: expected %s, received %s", ans.String(), res.String())
	}
}
