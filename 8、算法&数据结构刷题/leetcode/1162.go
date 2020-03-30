type Node struct {
    X int
    Y int
}
var dirX []int
var dirY []int

func init() {
    // up, right, down, left
    dirX = []int{-1,0,1,0}
    dirY = []int{0,1,0,-1}
}

func maxDistance(grid [][]int) int {
    n := len(grid)
    ret := bfsDistance(grid, n)
    return ret
}

func bfsDistance(grid [][]int, n int) int {
    ret := -1
    dst := make([][]int, n)
    for i:=0; i<n; i++ {
        dst[i] = make([]int, n)
    }
    var top, bottom int
    queue := make([]*Node, n*n)

    //source
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            if grid[i][j] == 1 {
                queue[top] = &Node{i,j}
                top ++
            }
        }
    }

    for bottom < top {
        curNode := queue[bottom]
        //fmt.Println(*curNode)
        bottom ++
        for i:=0; i<4; i++ {
            newX := curNode.X + dirX[i]
            newY := curNode.Y + dirY[i]
            if newX < 0 || newX >= n || newY < 0 || newY >= n {
                continue
            }
            if grid[newX][newY] != 0 {
                continue
            }
            if dst[newX][newY] == 0 {
                dst[newX][newY] = dst[curNode.X][curNode.Y]+1
                if dst[newX][newY] > ret {
                    ret = dst[newX][newY]
                }
                queue[top] = &Node{newX,newY}
                top ++
            }
        }
        //fmt.Println(dst)
    }
    //fmt.Println(dst)

    return ret
}