package tests

import (
	"linalg/permutation"
	"testing"
)

func TestPermutation(t *testing.T) {
	p := permutation.Id(3)
	_ = p.Swap(0, 1)
	res1 := permutation.Of([]int{1, 0, 2})

	if ok, _ := p.Equals(res1); !ok {
		t.Errorf("Expected %s, but received %s", res1, p)
	}

	_ = p.Swap(0, 2)
	res2 := permutation.Of([]int{2, 0, 1})

	if ok, _ := p.Equals(res2); !ok {
		t.Errorf("Expected %s, but received %s", res2, p)
	}
}
