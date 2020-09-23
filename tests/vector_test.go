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
