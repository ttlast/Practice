/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{val, nil,nil}
    }

    if (val > root.Val) {
        newRoot := TreeNode{val, nil,nil}
        newRoot.Left = root
        return &newRoot
    }

    root.Right = insertIntoMaxTree(root.Right, val)
    return root
}