package arrays

import "fmt"

func CountSquares(matrix [][]int) int {
	//dirs := [][]int {
	//	{0,1},
	//	{1,0},
	//	{1,1},
	//}
	result := 0 //make([]int, len(matrix)) // TODO min of m & n
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				fmt.Printf("%v %v %v\n", i, j, findOnes(matrix, i, j))
				result = result + findOnes(matrix, i, j)
			}
		}
	}

	return result
}

func findOnes(matrix [][]int, i, j int) int {
	result := 0
	for k := 0; k < len(matrix); k++ {
		if i+k >= len(matrix) || j+k >= len(matrix[0]) {
			break
		}
		allOnes := true
		q := j + k
		for p := i; p < i+k; p++ {
			if matrix[p][q] != 1 {
				allOnes = false
				break
			}
		}
		p := i + k
		for q := j; q < j+k; q++ {
			if matrix[p][q] != 1 {
				allOnes = false
				break
			}
		}
		if allOnes {
			result++
		}

	}

	return result
}

/*

	fmt.Printf("result %v\n", arrays.CountSquares([][]int{
		{0,0,1},
		{1,1,1},
		{0,1,1},
	}))

*/
