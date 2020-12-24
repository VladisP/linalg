package tests

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"linalg/matrix"
	"math"
	"testing"
)

func Test1(t *testing.T) {
	size := 10

	a, _ := fillRandomMatrix(size)
	L, U := a.LU()
	a2, _ := L.MulMatrix(U)

	fmt.Printf("Delta: %.15f\n", a.Sub(a2).Norm())
}

func Test2(t *testing.T) {
	size := 3

	a, _ := fillOPMatrix(size)
	L, U := a.LU()
	a2, _ := L.MulMatrix(U)

	fmt.Println(L)
	fmt.Println(U)
	fmt.Printf("Delta: %.15f\n", a.Sub(a2).Norm())
}

func Test3(t *testing.T) {
	size := 5

	a, _ := fillRandomMatrix(size)
	f := fillRandomVector(size)
	x := matrix.SolveSystemLU(a, f)
	x2 := matrix.SolveSystemGonum(a, f)
	delta, _ := x.Sub(x2)

	fmt.Printf("LU: \t%s\n\n", x)
	fmt.Printf("Gonum: \t%s\n\n", x2)
	fmt.Printf("Delta: \t%.15f\n\n", delta.Norm())
}

func Test4(t *testing.T) {
	size := 5

	a, _ := fillRandomMatrix(size)
	detLU := a.DetLU()
	detGonum := mat.Det(*a)
	delta := math.Abs(detLU - detGonum)

	fmt.Printf("LU: \t%.15f\n\n", detLU)
	fmt.Printf("Gonum: \t%.15f\n\n", detGonum)
	fmt.Printf("Delta: \t%.15f\n\n", delta)
}

func Test5(t *testing.T) {
	size := 10

	a, _ := fillRandomMatrix(size)
	invA1 := a.InverseLU()
	invA2, _ := matrix.InverseGonum(a)

	fmt.Printf("Delta: \t%.20f\n\n", invA1.Sub(invA2).Norm())
}
