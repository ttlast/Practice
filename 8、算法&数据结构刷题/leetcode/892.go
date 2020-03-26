var dirx []int
var diry []int

func surfaceArea(grid [][]int) int {
    // top right bottom left
    dirx = []int{-1, 0, 1, 0}
    diry = []int{0, 1, 0, -1}

    nlen := len(grid)
    var ret int
    
    for i:=0; i<nlen; i++ {
        for j:=0; j<nlen; j++ {
            if grid[i][j] <= 0 {
                continue
            }
            ret += getSurface(grid, nlen, i, j)
        }
    }
    return ret
}

func getSurface(grid [][]int, nlen int, x int, y int) int {
    ret := 2
    for i:=0; i<4; i++ {
        nx := x + dirx[i]
        ny := y + diry[i]
        if nx < 0 || nx >= nlen || ny < 0 || ny >= nlen {
            ret += grid[x][y]
            continue
        }
        if grid[x][y] > grid[nx][ny] {
            ret += grid[x][y] - grid[nx][ny]
        }
    }
    return ret
}