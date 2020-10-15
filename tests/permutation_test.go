package tests

import (
	"linalg/permutation"
	"testing"
)

func TestPermutation(t *testing.T) {
	p := permutation.Id(4)
	_ = p.Swap(0, 1)
	res1 := permutation.Of([]int{1, 0, 2, 3})

	if ok, _ := p.Equals(res1); !ok {
		t.Errorf("Expected %s, but received %s", res1, p)
	}

	p2 := permutation.Id(4)
	_ = p2.Swap(1, 3)
	res2 := permutation.Of([]int{0, 3, 2, 1})

	if ok, _ := p2.Equals(res2); !ok {
		t.Errorf("Expected %s, but received %s", res2, p2)
	}

	_ = p.Compose(p2)
	res3 := permutation.Of([]int{3, 0, 2, 1})

	if ok, _ := p.Equals(res3); !ok {
		t.Errorf("Expected %s, but received %s", res3, p)
	}
}
