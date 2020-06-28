package graph

import "fmt"

/*
  1 0 0 0
  1 1 0 0
  0 1 0 1
  0 0 1 0

ans = 3
*/

/*
	We can create a graph OR
	Better simpler way is to user 4 directions and
*/

func CountServers(grid [][]int) int {

	// grid -> graph map[Vertex][]*Vertex
	// count all that have len(value) > 0

	//graph := ServerGraph{
	//	Edges: map[ServerVertex][]*ServerVertex{},
	//}
	//
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	totalServers := 0
	notConnected := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				totalServers++

				isConnected := false
				for k := 0; k < len(dirs); k++ {
					row := i + dirs[k][0]
					col := j + dirs[k][1]

					fmt.Printf("%v, %v, %v, %v, %v\n", i, j, k, row, col)
					if row >= 0 && row < len(grid) &&
						col >= 0 && col < len(grid[i]) &&
						grid[row][col] == 1 {

						isConnected = true
						fmt.Printf("%v, %v, %v, %v, %v %v\n", i, j, k, row, col, isConnected)
						break
					}
				}
				if !isConnected {
					notConnected++
				}
			}
		}
	}

	fmt.Printf("%v, %v \n", totalServers, notConnected)
	return totalServers - notConnected

}

/*
fmt.Printf("result %v\n", graph.CountServers([][]int {
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
		{0, 0, 1, 0},
	}))

*/
