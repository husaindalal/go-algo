package hash

func FlipColumsEqualRows(matrix [][]int) int {

	maxRows := 0
	for i := 0; i < len(matrix); i++ {
		equalRows := 0
		for j := i; j < len(matrix); j++ {
			eq := areRowsEqual(matrix, i, j)
			feq := areFlipEqual(matrix, i, j)
			if eq || feq {
				equalRows++
			}
		}

		if equalRows > maxRows {
			maxRows = equalRows
		}
	}

	return maxRows
}

func areRowsEqual(matrix [][]int, i, j int) bool {
	for k := 0; k < len(matrix[0]); k++ {
		if matrix[i][k] != matrix[j][k] {
			return false
		}
	}
	return true
}

func areFlipEqual(matrix [][]int, i, j int) bool {
	for k := 0; k < len(matrix[0]); k++ {
		if matrix[i][k] == matrix[j][k] {
			return false
		}
	}
	return true
}

/*
result := hash.FlipColumsEqualRows([][]int {
		{1,0,0,1,0},
		{1,0,0,1,0},
		{1,0,1,1,1},
		{0,1,1,0,1},
		{1,0,0,1,1},
		{0,1,1,0,1},
	})


	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

*/
