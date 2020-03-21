/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 const INT_MAX = int(^uint(0) >> 1)
 func getMinimumDifference(root *TreeNode) int {
	retInt := INT_MAX

	pLeft := root.Left
	for nil != pLeft {
		if nil == pLeft.Right {
			break
		}
        pLeft = pLeft.Right
	}
	if nil != pLeft {
		retInt = root.Val - pLeft.Val
	}

	pRight := root.Right
	for nil != pRight {
		if nil == pRight.Left {
			break
		}
        pRight = pRight.Left
	}
	if nil != pRight {
		retInt = MinInt(retInt, pRight.Val - root.Val)
	}

	if nil != root.Left {
		retInt = MinInt(retInt, getMinimumDifference(root.Left))
	}

	if nil != root.Right {
		retInt = MinInt(retInt, getMinimumDifference(root.Right))
	}

	return retInt
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}