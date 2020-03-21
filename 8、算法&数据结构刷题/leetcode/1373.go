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

func maxSumBST(root *TreeNode) int {
	var ret int

	return ret
}

// int : is bst tree
// int , int, int : sum, min, max
func isSearchTree(root *TreeNode) (int, int, int, int) {
	sign := 1
	ret, min, max := 0, 0, 0
	if nil == root {
		return sign, ret, min, max
	}

	min = max = root.Val
	ret = root.Val
	if nil != root.Left {
		lsign, lsum, lmin, lmax := isSearchTree(root.Left)
		sign &= lsign
		if sign > 0 && root.Val > lmax {
			ret += lsum
			min = lmin
		}
	}
	if nil != root.Right {
		rsign, rsum, rmin, rmax := isSearchTree(root.Right)
		sign &= rsign
		if sign > 0 && root.Val < rmin {
			ret += rsum
			max = rmax
		}
	}

	return sign, ret, min, max
}