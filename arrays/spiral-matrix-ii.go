package arrays

import "fmt"

func GenerateMatrixII(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	fmt.Printf("res  %v \n", result)

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	counts := []int{n - 1}
	for i := n - 1; i > 0; i-- {
		counts = append(counts, i, i)
	}

	x := 0
	y := 0
	val := 1
	result[0][0] = 1
	for i := 0; i < len(counts); i++ {
		dir := dirs[i%4]
		for j := 0; j < counts[i]; j++ {
			val = val + 1
			x = x + dir[0]
			y = y + dir[1]

			result[x][y] = val
		}
	}

	return result
}
