package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestStrass(t *testing.T) {
	sizes := []int{1, 100, 200, 400, 600, 800, 1000, 1600, 2000}
	res := ""

	classic := benchmarkClassic(sizes)
	for _, v := range classic {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	str1 := benchmarkStrass(sizes, false)
	for _, v := range str1 {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	str2 := benchmarkStrass(sizes, true)
	for _, v := range str2 {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	fmt.Println(strings.ReplaceAll(res, ".", ","))
}

func benchmarkClassic(sizes []int) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		m1, _ := fillRandomMatrix(size)
		m2, _ := fillRandomMatrix(size)

		start := time.Now()
		_, _ = m1.MulMatrix(m2)
		duration := float64(time.Since(start).Nanoseconds()) * 1e-6

		res[i] = duration
	}

	return res
}

func benchmarkStrass(sizes []int, concurrency bool) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		m1, _ := fillRandomMatrix(size)
		m2, _ := fillRandomMatrix(size)

		start := time.Now()
		_, _ = m1.MulStrass(m2, concurrency)
		duration := float64(time.Since(start).Nanoseconds()) * 1e-6

		res[i] = duration
	}

	return res
}
