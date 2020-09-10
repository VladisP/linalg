package main

import (
	"fmt"
	"linalg/gauss"
	"linalg/matrix"
	"linalg/vector"
)

func main() {
	// Scalar product test
	fmt.Println(" ---------- Scalar product test ---------- ")

	v1 := vector.NewVector([]float64{1, 2, 3})
	v2 := vector.NewVector([]float64{4, 5, 6})
	p, err := v1.ScalarProd(v2)

	fmt.Printf("v1 = %s\n", v1)
	fmt.Printf("v2 = %s\n", v2)
	fmt.Printf("(v1, v2): %.2f\n", p)
	fmt.Printf("Error: %v\n", err)
	fmt.Println()

	// Matrix mul test
	fmt.Println(" ---------- Matrix mul test ---------- ")

	m1, _ := matrix.NewMatrix([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	m2, _ := matrix.NewMatrix([][]float64{
		{1, 3, 5},
		{2, 4, 6},
	})
	m, err := m1.MulMatrix(m2)

	fmt.Println("m1 =")
	fmt.Println(m1)
	fmt.Println("m2 =")
	fmt.Println(m2)
	fmt.Println("m1 * m2 =")
	fmt.Println(m)
	fmt.Printf("Error: %v\n", err)
	fmt.Println()

	// Matrix * Vector test
	fmt.Println(" ---------- Matrix * Vector test ---------- ")

	v, err := m2.MulVector(v2)
	fmt.Printf("m2 * v2 = %s\n", v)
	fmt.Printf("Error: %v\n", err)
	fmt.Println()

	// Triangular test
	fmt.Println(" ---------- Triangular test ---------- ")

	testMatrix, _ := matrix.NewMatrix([][]float64{
		{1, 2, 3},
		{3, 5, 7},
		{1, 3, 4},
	})
	testVector := vector.NewVector([]float64{3, 0, 1})

	a, f, err := gauss.TriangularMatrix(testMatrix, testVector)
	fmt.Println(testMatrix)
	fmt.Println(testVector)
	fmt.Println(a)
	fmt.Println(f)
	fmt.Println(err)
	fmt.Println()

	// Substitution test
	fmt.Println(" ---------- Substitution test ---------- ")

	res, err := gauss.Substitution(a, f)
	fmt.Println(res)
	fmt.Println(err)
	fmt.Println()

	// Second gauss test
	fmt.Println(" ---------- Second gauss test ---------- ")

	testMatrix2, _ := matrix.NewMatrix([][]float64{
		{2, 4, 1},
		{5, 2, 1},
		{2, 3, 4},
	})
	testVector2 := vector.NewVector([]float64{36, 47, 37})

	a2, f2, err := gauss.TriangularMatrix(testMatrix2, testVector2)
	fmt.Println(testMatrix2)
	fmt.Println(testVector2)
	fmt.Println(a2)
	fmt.Println(f2)
	fmt.Println(err)

	res2, err := gauss.Substitution(a2, f2)
	fmt.Println(res2)
	fmt.Println(err)
}
