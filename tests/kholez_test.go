package tests

import (
	"fmt"
	"linalg/gauss"
	"linalg/kholez"
	"strings"
	"testing"
	"time"
)

func TestKholezBenchmark(t *testing.T) {
	sizes := []int{3, 10, 20, 50, 100, 200, 500, 1000}
	res := ""

	k := benchmarkKholez(sizes)
	for _, v := range k {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	fmt.Println(strings.ReplaceAll(res, ".", ","))
}

func benchmarkKholez(sizes []int) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		testMatrix, _ := fillOPMatrix(size)
		testVector := fillRandomVector(size)

		start := time.Now()
		k, _ := kholez.Kholez(testMatrix)
		fv, _ := gauss.ClassicGauss(k, testVector)
		_, _ = gauss.ClassicGauss(k.Transpose(), fv)
		duration := float64(time.Since(start).Nanoseconds()) * 1e-6

		res[i] = duration
	}

	return res
}
