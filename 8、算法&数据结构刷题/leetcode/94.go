/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func inorderTraversal(root *TreeNode) []int {
    if nil == root {
        return {}[]int
    }
    var resultList []int
    resultList = append(resultList, inorderTraversal(root.Left)...)
    resultList = append(resultList, root.Val)
    resultList = append(resultList, inorderTraversal(root.Right)...)
    return resultList
}
