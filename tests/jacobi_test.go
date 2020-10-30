package tests

import (
	"fmt"
	"linalg/gauss"
	"linalg/jacobi"
	"testing"
	"time"
)

func TestJacobi(t *testing.T) {
	size := 1000

	fmt.Println("---------------JACOBI VS GAUSS---------------")
	testMatrix, _ := fillDominanceMatrix(size)
	testVector := fillRandomVector(size)

	fmt.Println("Gauss")

	start := time.Now()
	solution, _ := gauss.ClassicGauss(testMatrix, testVector)
	duration := time.Since(start).Milliseconds()

	fmt.Printf("Solution: %s\n", solution)
	res, _ := testMatrix.MulVector(solution)
	delta, _ := res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.UniformNorm())
	fmt.Printf("Duration: %d ms\n", duration)

	fmt.Println("Jacobi")

	start = time.Now()
	solution, iterCount, _ := jacobi.Jacobi(testMatrix, testVector)
	duration = time.Since(start).Milliseconds()

	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.UniformNorm())
	fmt.Printf("Duration: %d ms\n", duration)
	fmt.Printf("Iter count: %d\n", iterCount)
	fmt.Println("---------------END TEST---------------")
}
