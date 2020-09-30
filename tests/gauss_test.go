package tests

import (
	"fmt"
	"linalg/gauss"
	"linalg/matrix"
	"linalg/vector"
	"testing"
)

const Eps = 1e-6

func TestClassicGauss(t *testing.T) {
	testMatrix, _ := matrix.NewMatrix([][]float64{
		{2, 4, 1},
		{5, 2, 1},
		{2, 3, 4},
	})
	testVector := vector.NewVector([]float64{36, 47, 37})
	ans := vector.NewVector([]float64{7, 5, 2})

	res, err := gauss.ClassicGauss(testMatrix, testVector)
	if err != nil {
		t.Error(err)
	}

	if delta, _ := ans.Sub(res); delta.Norm() >= Eps {
		t.Errorf("Answer: expected %s, received %s", ans.String(), res.String())
	}
}

func TestClassicGauss2(t *testing.T) {
	testMatrix, _ := fillRandomMatrix(1000)
	testVector := fillRandomVector(1000)

	solution, err := gauss.ClassicGauss(testMatrix, testVector)
	if err != nil {
		t.Error(err)
	}

	ans, _ := testMatrix.MulVector(solution)

	if delta, _ := ans.Sub(testVector); delta.Norm() >= Eps {
		t.Errorf("Expected %s, received %s", testVector.String(), ans.String())
		t.Errorf("Solution vector is %s", solution.String())
	}
}

func TestPivotingGauss(t *testing.T) {
	testMatrix, _ := matrix.NewMatrix([][]float64{
		{2, 4, 1},
		{5, 2, 1},
		{2, 3, 4},
	})
	testVector := vector.NewVector([]float64{36, 47, 37})
	ans := vector.NewVector([]float64{7, 5, 2})

	res, err := gauss.PivotingGauss(testMatrix, testVector)
	if err != nil {
		t.Error(err)
	}

	if delta, _ := ans.Sub(res); delta.Norm() >= Eps {
		t.Errorf("Answer: expected %s, received %s", ans.String(), res.String())
	}
}

func TestPivotingGauss2(t *testing.T) {
	testMatrix, _ := fillRandomMatrix(1000)
	testVector := fillRandomVector(1000)

	solution, err := gauss.PivotingGauss(testMatrix, testVector)
	if err != nil {
		t.Error(err)
	}

	ans, _ := testMatrix.MulVector(solution)

	if delta, _ := ans.Sub(testVector); delta.Norm() >= Eps {
		t.Errorf("\nExpected %s \n Received %s", testVector, ans)
		t.Errorf("Solution vector is %s", solution)
	}
}

func TestAllGauss(t *testing.T) {
	size := 10

	fmt.Println("---------------RANDOM---------------")
	testMatrix, _ := fillRandomMatrix(size)
	testVector := fillRandomVector(size)

	fmt.Println("Test Matrix")
	fmt.Println(testMatrix)
	fmt.Println("Test Vector")
	fmt.Println(testVector)

	fmt.Println("Classic Gauss")
	solution, _ := gauss.ClassicGauss(testMatrix, testVector)
	fmt.Printf("Solution: %s\n", solution)
	res, _ := testMatrix.MulVector(solution)
	delta, _ := res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.Norm())

	fmt.Println("Pivoting Gauss")
	solution, _ = gauss.PivotingGauss(testMatrix, testVector)
	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.Norm())
	fmt.Println("---------------END RANDOM---------------")

	fmt.Println("---------------DOMINANCE---------------")
	testMatrix, _ = fillDominanceMatrix(size)
	testVector = fillRandomVector(size)

	fmt.Println("Test Matrix")
	fmt.Println(testMatrix)
	fmt.Println("Test Vector")
	fmt.Println(testVector)

	fmt.Println("Classic Gauss")
	solution, _ = gauss.ClassicGauss(testMatrix, testVector)
	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.Norm())

	fmt.Println("Pivoting Gauss")
	solution, _ = gauss.PivotingGauss(testMatrix, testVector)
	fmt.Printf("Solution: %s\n", solution)
	res, _ = testMatrix.MulVector(solution)
	delta, _ = res.Sub(testVector)
	fmt.Printf("Delta: %.15f\n", delta.Norm())
	fmt.Println("---------------END DOMINANCE---------------")
}
