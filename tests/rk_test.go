package tests

import (
	"fmt"
	"linalg/gauss"
	"linalg/matrix"
	"linalg/vector"
	"testing"
)

func TestRK(t *testing.T) {
	a, _ := matrix.NewMatrix([][]float64{
		{100, 99},
		{99, 98},
	})
	f := vector.NewVector([]float64{199, 197})

	sol1, _ := gauss.ClassicGauss(a, f)
	fmt.Printf("Solution: %s\n", sol1)

	deltaA, _ := matrix.NewMatrix([][]float64{
		{0, 0},
		{0, 0},
	})
	deltaF := vector.NewVector([]float64{-0.01, 0.01})

	condNum, _ := a.ConditionNumber2()
	fmt.Printf("Cond Num: %.15f\n", condNum)

	expect := ((deltaF.UniformNorm() / f.UniformNorm()) + (deltaA.UniformNorm() / a.UniformNorm())) * condNum
	fmt.Printf("Expect: %.15f\n", expect)

	aPlusDelta := a.Sum(deltaA)
	fPlusDelta, _ := f.Sum(deltaF)

	sol2, _ := gauss.ClassicGauss(aPlusDelta, fPlusDelta)
	fmt.Printf("Solution 2: %s\n", sol2)

	deltaSol, _ := sol2.Sub(sol1)
	fmt.Printf("Delta x: %s\n", deltaSol)

	e := deltaSol.UniformNorm() / sol1.UniformNorm()
	fmt.Printf("Ans: %.15f\n", e)

	fmt.Printf("Is ok: %v\n", e <= expect)
}
