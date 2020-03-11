/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameterRsult int

	diameterRsult, _ = depthOfBinaryTree(root)

	return diameterRsult
}

// diameterResult, depthResult
func depthOfBinaryTree(root *TreeNode) (int, int) {
	if nil == root {
		return 0, 0
	}

	var diameterResult int
	var depthResult int

	if nil != root.Left {
		diameterLeft, depthLeft := depthOfBinaryTree(root.Left)
		depthResult = depthLeft + 1
		diameterResult = diameterLeft
		diameterResult = MaxInt(diameterResult, depthLeft + 1)
	}

	if nil != root.Right {
		diameterRight, depthRight := depthOfBinaryTree(root.Right)

		diameterResult = MaxInt(diameterResult, diameterRight)
		diameterResult = MaxInt(diameterResult, depthResult + depthRight+1)
		depthResult = MaxInt(depthResult, depthRight + 1)
	}

	return diameterResult, depthResult
}