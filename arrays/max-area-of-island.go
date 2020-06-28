package arrays

func MaxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	//dirs := [][]int {
	//	{-1, 0},
	//	{1, 0},
	//	{0, -1},
	//	{0, 1},
	//}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				area := findArea(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func findArea(grid [][]int, i, j int) int {
	//area := 0
	if i < 0 || i >= len(grid) ||
		j < 0 || j >= len(grid[0]) ||
		grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0

	area := 1 + findArea(grid, i-1, j) +
		findArea(grid, i+1, j) +
		findArea(grid, i, j-1) +
		findArea(grid, i, j+1)

	return area
}

/*
	fmt.Printf("result %v\n", arrays.MaxAreaOfIsland([][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}))

*/
