/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func longestZigZag(root *TreeNode) int {
    var retVal int
    dfsLRTree(root, &retVal)
    return retVal
}

// left, right
func dfsLRTree(root *TreeNode, retVal *int) (int, int) {
    var left, right int
    if nil != root.Left {
        _,rr := dfsLRTree(root.Left, retVal)
        left = rr+1
    }
    if nil != root.Right {
        ll,_ := dfsLRTree(root.Right, retVal)
        right = ll+1
    }
    *retVal = MaxInt(*retVal, left)
    *retVal = MaxInt(*retVal, right)
    return left, right
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}