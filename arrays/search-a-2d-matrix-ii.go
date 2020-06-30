package arrays

func SearchMatrix(matrix [][]int, target int) bool {
	// validate
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix) - 1
	n := len(matrix[0]) - 1

	i := 0
	j := n
	for i >= 0 && j >= 0 && i <= m && j <= n {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			j--
		} else {
			i++
		}
	}
	return false
}
