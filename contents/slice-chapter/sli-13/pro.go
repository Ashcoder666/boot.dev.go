package sli13

func createMatrix(rows, cols int) [][]int {
	multiplicationTable := make([][]int, rows)

	for i := 0; i < rows; i++ {
		multiplicationTable[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			multiplicationTable[i][j] = i * j
		}
	}

	return multiplicationTable
}
