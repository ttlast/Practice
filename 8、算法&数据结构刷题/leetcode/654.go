/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructMaximumBinaryTree(nums []int) *TreeNode {
    root := dfsTree(nums, 0, len(nums))

    return root
}

func dfsTree(nums []int, l int, r int) *TreeNode {
    if l>=r {
        return nil
    }
    maxpos := l
    for i:=l+1; i<r; i++ {
        if nums[i] > nums[maxpos] {
            maxpos = i
        }
    }

    curNode := TreeNode{nums[maxpos], nil, nil}
    curNode.Left = dfsTree(nums, l, maxpos)
    curNode.Right = dfsTree(nums, maxpos+1, r)

    return &curNode
}