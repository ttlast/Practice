func maxAreaOfIsland(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	var retArea int
	for i := range grid {
		for j:= range grid[0] {
			if grid[i][j] == 1 {
				retArea = MaxInt(retArea, dfsIsland(grid, i, j, h, w))
			}
		}
	}
	return retArea
}

func dfsIsland(grid [][]int, i int, j int, h int, w int) int {
	var retVal int
	if grid[i][j] != 1 {
		return retVal
	}

	grid[i][j] = 0
	retVal ++
	
	if i-1 >= 0 {
		retVal += dfsIsland(grid, i-1, j, h, w)
	}
	if i+1 < h {
		retVal += dfsIsland(grid, i+1, j, h, w)
	}
	if j-1 >= 0 {
		retVal += dfsIsland(grid, i, j-1, h, w)
	}
	if j+1 < w {
		retVal += dfsIsland(grid, i, j+1, h, w)
	}
	return retVal
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}