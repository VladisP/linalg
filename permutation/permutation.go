package permutation

import (
	"fmt"
	"strconv"
)

type Permutation []int

func Id(n int) Permutation {
	p := make(Permutation, n)

	for i := 0; i < n; i++ {
		p[i] = i
	}

	return p
}

func Of(v []int) Permutation {
	p := make(Permutation, len(v))
	copy(p, v)
	return p
}

func (p Permutation) Swap(i, j int) error {
	if i >= len(p) || j >= len(p) {
		return fmt.Errorf(ErrorIndexOutOfRange)
	}
	p[i], p[j] = p[j], p[i]
	return nil
}

func (p Permutation) Equals(permutation Permutation) (bool, error) {
	if err := checkSizes(p, permutation); err != nil {
		return false, err
	}

	for i := range p {
		if p[i] != permutation[i] {
			return false, nil
		}
	}

	return true, nil
}

func (p Permutation) String() string {
	s := "("

	for i, c := range p {
		s += strconv.Itoa(c)

		if i != len(p)-1 {
			s += ", "
		}
	}

	return s + ")"
}
