package tests

import (
	"fmt"
	"linalg/jacobi"
	"linalg/seidel"
	"testing"
	"time"
)

func TestSeidel(t *testing.T) {
	size := 10000

	fmt.Println("---------------SEIDEL VS JACOBI---------------")
	testMatrix, _ := fillDominanceMatrix(size)
	testVector := fillRandomVector(size)

	fmt.Println("Jacobi")

	start := time.Now()
	solution, iterCount, _ := jacobi.Jacobi(testMatrix, testVector)
	duration := time.Since(start).Milliseconds()

	fmt.Printf("Solution: %s\n", solution)
	res, _ := testMatrix.MulVector(solution)
	delta, _ := res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.UniformNorm())
	fmt.Printf("Duration: %d ms\n", duration)
	fmt.Printf("Iter count: %d\n", iterCount)

	fmt.Println("Seidel")

	start = time.Now()
	solution, iterCount, _ = seidel.Seidel(testMatrix, testVector)
	duration = time.Since(start).Milliseconds()

	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.UniformNorm())
	fmt.Printf("Duration: %d ms\n", duration)
	fmt.Printf("Iter count: %d\n", iterCount)
	fmt.Println("---------------END TEST---------------")
}
