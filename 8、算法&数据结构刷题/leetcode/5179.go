/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // 排序后，重新生成下树就可以了
 func balanceBST(root *TreeNode) *TreeNode {
    var intList []int
    getNodeVal(root, &intList)
    
    sort.Ints(intList)
    
    return buildBST(intList, 0, len(intList))
}

func getNodeVal(root *TreeNode,intList *[]int) {
    if root == nil {
        return 
    }
    *intList = append(*intList, root.Val)
    getNodeVal(root.Left, intList)
    getNodeVal(root.Right, intList)
}

func buildBST(intList []int, l int, r int) *TreeNode {
    if (l >= r) {
        return nil
    }
    mid := (l+r)>>1;
    curNode := TreeNode{}
    curNode.Val = intList[mid]
    curNode.Left = buildBST(intList, l, mid)
    curNode.Right =  buildBST(intList, mid+1, r)
    return &curNode
}


