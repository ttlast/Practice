/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func kthToLast(head *ListNode, k int) int {
    var pre,cur *ListNode
    cur = head
    pre = head
    for i:=1; i<k; i++ {
        pre = pre.Next
    }
    for nil != pre.Next {
        cur = cur.Next
        pre = pre.Next
    }

    return cur.Val
}