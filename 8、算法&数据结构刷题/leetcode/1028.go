 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
    mm := make(map[int]*TreeNode)
    var root *TreeNode
    var val int
    var level int
    for i:=0; i<len(S); i++ {
        if S[i] == '-' {
            insertVal(&root, mm, &val, &level)
            level ++
        } else {
            val = val*10+int(S[i]-'0')
        }
    }
    insertVal(&root, mm, &val, &level)
    
    return root
}

func insertVal(root **TreeNode, mm map[int]*TreeNode, val *int, level *int) {
    if *val != 0 {       //val in [1, 10^9]
        node := &TreeNode{*val,nil,nil}
        mm[*level] = node
        if *level == 0 {
            *root = node
        } else {
            pre := mm[*level-1]
            if pre.Left == nil {
                    pre.Left = node
            } else {
                    pre.Right = node
            }
        }
        *val = 0
        *level = 0
    }
}