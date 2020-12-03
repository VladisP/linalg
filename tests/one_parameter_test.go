package tests

import (
	"fmt"
	"linalg/jacobi"
	"linalg/matrix"
	"linalg/one_parameter"
	"linalg/seidel"
	"linalg/vector"
	"strings"
	"testing"
	"time"
)

func TestOneParameterSpecialCase(t *testing.T) {
	fmt.Println("---------------|ONE PARAMETER VS SEIDEL VS JACOBI| (Special Case)---------------")
	testMatrix, _ := matrix.NewMatrix([][]float64{
		{128, 3, 6},
		{3, 209, 2},
		{6, 2, 105},
	})
	testVector := vector.NewVector([]float64{23, 36, 8})

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

	fmt.Println("One parameter")
	tau := 0.0063964589

	start = time.Now()
	solution, iterCount, _ = one_parameter.OneParameter(testMatrix, testVector, tau)
	duration = time.Since(start).Milliseconds()

	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.UniformNorm())
	fmt.Printf("Duration: %d ms\n", duration)
	fmt.Printf("Iter count: %d\n", iterCount)
	fmt.Println("---------------END TEST---------------")
}

func TestOneParameter(t *testing.T) {
	size := 1000

	fmt.Println("---------------ONE PARAMETER VS SEIDEL VS JACOBI---------------")
	testMatrix, _ := fillOPMatrix(size)
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

	fmt.Println("One parameter")

	startEigen := time.Now()
	_, _, tau, err := one_parameter.GetEigen(testMatrix)
	if err != nil {
		t.Error(err)
	}
	durationEigen := time.Since(startEigen).Milliseconds()
	fmt.Printf("Duration eigen: %d ms\n", durationEigen)
	fmt.Printf("Tau: %.15f\n", tau)

	start = time.Now()
	solution, iterCount, _ = one_parameter.OneParameter(testMatrix, testVector, tau)
	duration = time.Since(start).Milliseconds()

	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.UniformNorm())
	fmt.Printf("Duration: %d ms\n", duration)
	fmt.Printf("Iter count: %d\n", iterCount)
	fmt.Println("---------------END TEST---------------")
}

func TestBenchmarks(t *testing.T) {
	sizes := []int{3, 10, 20, 50, 100, 200, 500, 1000}
	res := ""

	jac := benchmarkJacobi(sizes)
	for _, v := range jac {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	seid := benchmarkSeidel(sizes)
	for _, v := range seid {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	opteigen, opt, max, min := benchmarkOneParameter(sizes)
	for _, v := range opteigen {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"
	for _, v := range opt {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"
	for _, v := range max {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"
	for _, v := range min {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	fmt.Println(strings.ReplaceAll(res, ".", ","))
}

func benchmarkJacobi(sizes []int) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		testMatrix, _ := fillOPMatrix(size)
		testVector := fillRandomVector(size)

		start := time.Now()
		_, _, _ = jacobi.Jacobi(testMatrix, testVector)
		duration := float64(time.Since(start).Nanoseconds()) * 1e-6

		res[i] = duration
	}

	return res
}

func benchmarkSeidel(sizes []int) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		testMatrix, _ := fillOPMatrix(size)
		testVector := fillRandomVector(size)

		start := time.Now()
		_, _, _ = seidel.Seidel(testMatrix, testVector)
		duration := float64(time.Since(start).Nanoseconds()) * 1e-6

		res[i] = duration
	}

	return res
}

func benchmarkOneParameter(sizes []int) (opteigen []float64, opt []float64, max []float64, min []float64) {
	opteigen = make([]float64, len(sizes))
	opt = make([]float64, len(sizes))
	max = make([]float64, len(sizes))
	min = make([]float64, len(sizes))

	for i, size := range sizes {
		testMatrix, _ := fillOPMatrix(size)
		testVector := fillRandomVector(size)

		startEigen := time.Now()
		maxEigen, _, optEigen, _ := one_parameter.GetEigen(testMatrix)
		durationEigen := float64(time.Since(startEigen).Nanoseconds()) * 1e-6

		startOpt := time.Now()
		_, _, _ = one_parameter.OneParameter(testMatrix, testVector, optEigen)
		durationOpt := float64(time.Since(startOpt).Nanoseconds()) * 1e-6

		startMax := time.Now()
		_, _, _ = one_parameter.OneParameter(testMatrix, testVector, 2/maxEigen)
		durationMax := float64(time.Since(startMax).Nanoseconds()) * 1e-6

		startMin := time.Now()
		_, _, _ = one_parameter.OneParameter(testMatrix, testVector, 0)
		durationMin := float64(time.Since(startMin).Nanoseconds()) * 1e-6

		opteigen[i] = durationOpt + durationEigen
		opt[i] = durationOpt
		max[i] = durationMax
		min[i] = durationMin
	}
	return
}
