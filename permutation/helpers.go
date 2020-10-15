package permutation

import "fmt"

func checkSizes(p1 Permutation, p2 Permutation) error {
	if len(p1) != len(p2) {
		return fmt.Errorf(ErrorDimensions)
	}
	return nil
}
