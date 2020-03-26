func numRookCaptures(board [][]byte) int {
    var ret int
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[0]); j++ {
            if board[i][j] == 'R' {
                ret = check(board, i, j)
            }
        }
    }
    return ret
}

func check(board[][]byte, x int, y int) int {
    var ret = 0
    // top. right bottom left
    for nx:=x-1; nx >= 0; nx-- {
        if checkPoint(board, nx, y, &ret) {
            break
        }
    }
    for ny:=y+1; ny < 8; ny++ {
        if checkPoint(board, x, ny, &ret) {
            break
        }
    }
    for nx:=x+1; nx < 8; nx++ {
        if checkPoint(board, nx, y, &ret) {
            break
        }
    }
    for ny:=y-1; ny >= 0; ny-- {
        if checkPoint(board, x, ny, &ret) {
            break
        }
    }

    return ret
}

func checkPoint(board [][]byte, nx int, ny int, result *int) bool {
    if board[nx][ny] == 'p' {
        (*result) ++
        return true
    }
    if board[nx][ny] != '.' {
        return true
    }
    return false
}