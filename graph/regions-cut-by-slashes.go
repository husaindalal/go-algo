package graph

import (
	"fmt"
)

/*
/ = i, j+1 <> i+1, j
| = i, j <> i+1, j+1

expand each cell 3 times. and each / is marked with 1's
eg / ==
0 0 1
0 1 0
1 0 0

find islands

*/

func RegionsBySlashes(grid []string) int {
	// validate len, characters

	graph := make([][]int, len(grid)*3)
	for i := range grid {
		for k := 0; k < 3; k++ {
			graph[i*3+k] = make([]int, len(grid)*3)
		}
	}

	fmt.Printf("graph1 %v\n", graph)

	for i, word := range grid {
		for j, char := range word {
			if char == '/' {
				populateForward(graph, i, j)
			} else if char == '\\' {
				populateBackward(graph, i, j)
			}
		}
	}

	fmt.Printf("graph2 %v\n", graph)

	regions := 0

	for r := range graph {
		for c := range graph[r] {
			if graph[r][c] == 0 {
				markAdjacent(graph, r, c)
				fmt.Printf("graph2 %v %v %v\n", r, c, graph)
				regions++
			}
		}
	}

	return regions
}

func markAdjacent(graph [][]int, r int, c int) {
	//if graph[r][c] == 1 {
	//	return
	//}
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, dir := range dirs {
		row := r + dir[0]
		col := c + dir[1]

		if row >= 0 && col >= 0 &&
			row < len(graph) && col < len(graph[0]) &&
			graph[row][col] == 0 {

			graph[row][col] = 1
			markAdjacent(graph, row, col)

		}
	}
}

func populateForward(graph [][]int, i int, j int) {

	graph[i*3][j*3+2] = 1
	graph[i*3+1][j*3+1] = 1
	graph[i*3+2][j*3] = 1
}

func populateBackward(graph [][]int, i int, j int) {

	graph[i*3][j*3] = 1
	graph[i*3+1][j*3+1] = 1
	graph[i*3+2][j*3+2] = 1
}

/*
[0 0 0 0 0 1]
[0 0 0 0 1 0]
[0 0 0 1 0 0]
[0 0 1 0 0 0]
[0 1 0 0 0 0]
[1 0 0 0 0 0]]

fmt.Printf("result %v\n", graph.RegionsBySlashes([]string{" /", "/ "}))

*/
