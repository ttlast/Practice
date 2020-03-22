/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func maxSumBST(root *TreeNode) int {
	var ret int
	isSearchTree(root, &ret)
	return ret
}

// int : is bst tree
// int , int, int : sum, min, max
func isSearchTree(root *TreeNode, result *int) (int, int, int, int) {
	sign := 1
	ret, min, max := 0, 0, 0
	if nil == root {
		return sign, ret, min, max
	}

	min, max, ret = root.Val, root.Val, root.Val
	if nil != root.Left {
		lsign, lsum, lmin, lmax := isSearchTree(root.Left, result)
        if root.Val <= lmax {
            sign = 0
        }
		sign &= lsign
		if sign > 0 && root.Val > lmax {
			ret += lsum
			min = lmin
		}
	}
	if nil != root.Right {
		rsign, rsum, rmin, rmax := isSearchTree(root.Right, result)
        if root.Val >= rmin {
            sign = 0
        }
		sign &= rsign
		if sign > 0 && root.Val < rmin {
			ret += rsum
			max = rmax
		}
	}
    if sign > 0 && *result < ret {
        *result = ret
    }
    
	return sign, ret, min, max
}