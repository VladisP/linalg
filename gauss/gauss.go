package gauss

import (
	"linalg/matrix"
	"linalg/permutation"
	"linalg/vector"
)

func ClassicGauss(a *matrix.Matrix, f *vector.Vector) (*vector.Vector, error) {
	if err := checkConstraints(a, f); err != nil {
		return nil, err
	}

	aCopy := a.Copy()
	fCopy := f.Copy()

	for i := 0; i < aCopy.RowCount; i++ {
		if err := eliminate(i, aCopy, fCopy); err != nil {
			return nil, err
		}
	}

	return substitution(aCopy, fCopy)
}

func PivotingGauss(a *matrix.Matrix, f *vector.Vector) (*vector.Vector, error) {
	if err := checkConstraints(a, f); err != nil {
		return nil, err
	}

	aCopy := a.Copy()
	fCopy := f.Copy()
	p := permutation.Id(aCopy.ColumnCount)

	for i := 0; i < aCopy.RowCount; i++ {
		if err := columnPivoting(i, aCopy, fCopy); err != nil {
			return nil, err
		}
		if err := eliminate(i, aCopy, fCopy); err != nil {
			return nil, err
		}
		if err := rowPivoting(i, aCopy, p); err != nil {
			return nil, err
		}
		if err := eliminate(i, aCopy, fCopy); err != nil {
			return nil, err
		}
	}

	permutationSolution, err := substitution(aCopy, fCopy)
	if err != nil {
		return nil, err
	}

	solution := make([]float64, len(permutationSolution.Value))
	for i, v := range p {
		solution[i] = permutationSolution.Value[v]
	}

	return vector.NewVector(solution), err
}
