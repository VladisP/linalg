package vector

import (
	"fmt"
	"strconv"
)

type Vector struct {
	Size  int
	Value []float64
}

func NewVector(value []float64) *Vector {
	return &Vector{
		Size:  len(value),
		Value: value,
	}
}

func (v *Vector) ScalarProd(vector *Vector) (float64, error) {
	if !equalSize(v, vector) {
		return 0, fmt.Errorf(ErrorDimensions)
	}

	res := float64(0)

	for i := range v.Value {
		res += v.Value[i] * vector.Value[i]
	}

	return res, nil
}

func (v *Vector) String() string {
	s := "("

	for i, c := range v.Value {
		s += strconv.FormatFloat(c, 'f', -1, 64)

		if i != len(v.Value)-1 {
			s += ", "
		}
	}

	return s + ")"
}
