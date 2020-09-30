package vector

import (
	"fmt"
	"math"
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

func (v *Vector) Sub(vector *Vector) (*Vector, error) {
	if !equalSize(v, vector) {
		return nil, fmt.Errorf(ErrorDimensions)
	}

	value := make([]float64, len(v.Value))

	for i := range v.Value {
		value[i] = v.Value[i] - vector.Value[i]
	}

	return NewVector(value), nil
}

func (v *Vector) Norm() float64 {
	sc, _ := v.ScalarProd(v)

	return math.Sqrt(sc)
}

func (v *Vector) SwapComponents(i, j int) error {
	if i < 0 || j < 0 || i >= v.Size || j >= v.Size {
		return fmt.Errorf(ErrorIndexOutOfRange)
	}
	v.Value[i], v.Value[j] = v.Value[j], v.Value[i]
	return nil
}

func (v *Vector) Copy() *Vector {
	value := make([]float64, v.Size)
	copy(value, v.Value)

	return NewVector(value)
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
