package fb

import "fmt"

// https://www.interviewbit.com/problems/rotate-matrix

/*
 Value in the matrix is not important
 Can we use additional space? Is the structure immutable?

*/

func RotateMatrix(matrix [][]string) [][]string {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	n := len(matrix) - 1
	m := len(matrix[0]) - 1
	for i := 0; i <= n; i++ {
		for j := i; j <= m; j++ {
			fmt.Printf("i %d j %d  %s  %s \n", i, j, matrix[i][j], matrix[n-i][m-j])
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i <= n/2; i++ {
		for j := 0; j <= m; j++ {
			fmt.Printf("i %d j %d  %s  %s \n", i, j, matrix[i][j], matrix[n-i][m-j])
			matrix[i][j], matrix[n-i][j] = matrix[n-i][j], matrix[i][j]
		}
	}

	return matrix
}

//result := fb.RotateMatrix([][]string{
//{
//"A", "B", "C",
//},
//{
//"D", "E", "F",
//},
//{
//"G", "H", "I",
//},
//})
//
//fmt.Printf("Result %v \n", result)
