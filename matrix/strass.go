package matrix

import (
	"fmt"
	"sync"
)

const nMin = 128

func (m *Matrix) MulStrass(secondMatrix *Matrix, concurrency bool) (*Matrix, error) {
	if m.ColumnCount != secondMatrix.RowCount {
		return nil, fmt.Errorf(ErrorMulMatricesSize)
	}

	n := m.RowCount
	if secondMatrix.RowCount > n {
		n = secondMatrix.RowCount
	}
	if secondMatrix.ColumnCount > n {
		n = secondMatrix.ColumnCount
	}
	n = NearestPower2(n)

	mSquare1 := m.CompleteToSquare(n)
	mSquare2 := secondMatrix.CompleteToSquare(n)
	mul := mulStrass(mSquare1, mSquare2, concurrency)

	return mul.Slice(
		SliceRange{Start: 0, End: m.RowCount},
		SliceRange{Start: 0, End: secondMatrix.ColumnCount},
	), nil
}

func mulStrass(a *Matrix, b *Matrix, concurrency bool) *Matrix {
	n := a.RowCount

	if n <= nMin {
		mul, _ := a.MulMatrix(b)
		return mul
	}

	middle := n / 2
	u := SliceRange{Start: 0, End: middle}
	v := SliceRange{Start: middle, End: n}
	c := IdentityMatrix(a.RowCount, a.ColumnCount)

	var p1, p2, p3, p4, p5, p6, p7 *Matrix
	wg := sync.WaitGroup{}

	if concurrency {
		wg.Add(7)
		go func() {
			p1 = mulStrass(a.Slice(u, u).Sum(a.Slice(v, v)), b.Slice(u, u).Sum(b.Slice(v, v)), concurrency)
			wg.Done()
		}()
		go func() {
			p2 = mulStrass(a.Slice(v, u).Sum(a.Slice(v, v)), b.Slice(u, u), concurrency)
			wg.Done()
		}()
		go func() {
			p3 = mulStrass(a.Slice(u, u), b.Slice(u, v).Sub(b.Slice(v, v)), concurrency)
			wg.Done()
		}()
		go func() {
			p4 = mulStrass(a.Slice(v, v), b.Slice(v, u).Sub(b.Slice(u, u)), concurrency)
			wg.Done()
		}()
		go func() {
			p5 = mulStrass(a.Slice(u, u).Sum(a.Slice(u, v)), b.Slice(v, v), concurrency)
			wg.Done()
		}()
		go func() {
			p6 = mulStrass(a.Slice(v, u).Sub(a.Slice(u, u)), b.Slice(u, u).Sum(b.Slice(u, v)), concurrency)
			wg.Done()
		}()
		go func() {
			p7 = mulStrass(a.Slice(u, v).Sub(a.Slice(v, v)), b.Slice(v, u).Sum(b.Slice(v, v)), concurrency)
			wg.Done()
		}()
	} else {
		p1 = mulStrass(a.Slice(u, u).Sum(a.Slice(v, v)), b.Slice(u, u).Sum(b.Slice(v, v)), concurrency)
		p2 = mulStrass(a.Slice(v, u).Sum(a.Slice(v, v)), b.Slice(u, u), concurrency)
		p3 = mulStrass(a.Slice(u, u), b.Slice(u, v).Sub(b.Slice(v, v)), concurrency)
		p4 = mulStrass(a.Slice(v, v), b.Slice(v, u).Sub(b.Slice(u, u)), concurrency)
		p5 = mulStrass(a.Slice(u, u).Sum(a.Slice(u, v)), b.Slice(v, v), concurrency)
		p6 = mulStrass(a.Slice(v, u).Sub(a.Slice(u, u)), b.Slice(u, u).Sum(b.Slice(u, v)), concurrency)
		p7 = mulStrass(a.Slice(u, v).Sub(a.Slice(v, v)), b.Slice(v, u).Sum(b.Slice(v, v)), concurrency)
	}

	wg.Wait()

	CopyToSlice(c.Slice(u, u), p1.Sum(p4).Sub(p5).Sum(p7))
	CopyToSlice(c.Slice(u, v), p3.Sum(p5))
	CopyToSlice(c.Slice(v, u), p2.Sum(p4))
	CopyToSlice(c.Slice(v, v), p1.Sum(p3).Sub(p2).Sum(p6))

	return c
}
