package tests

import (
	"linalg/vector"
	"testing"
)

func TestScalar(t *testing.T) {
	v1 := vector.NewVector([]float64{1, 2, 3})
	v2 := vector.NewVector([]float64{4, 5, 6})
	ans := float64(32)

	p, err := v1.ScalarProd(v2)

	if err != nil {
		t.Error(err)
	}

	if p != ans {
		t.Errorf("Expected %f, but received %f", ans, p)
	}
}

func TestSub(t *testing.T) {
	v1 := vector.NewVector([]float64{1, 2, 3})
	v2 := vector.NewVector([]float64{4, 5, 6})
	ans := vector.NewVector([]float64{-3, -3, -3})
	res, _ := v1.Sub(v2)

	if res.String() != ans.String() {
		t.Errorf("Expected %s, but received %s", ans, res)
	}
}

func TestNorm(t *testing.T) {
	v := vector.NewVector([]float64{3, 4})
	norm := v.Norm()
	ans := float64(5)

	if norm != ans {
		t.Errorf("Expected %f, but received %f", ans, norm)
	}
}
