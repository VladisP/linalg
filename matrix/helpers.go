package matrix

func equalColumnsCount(rows [][]float64) bool {
	for i := 0; i < len(rows)-1; i++ {
		if len(rows[i]) != len(rows[i+1]) {
			return false
		}
	}

	return true
}
