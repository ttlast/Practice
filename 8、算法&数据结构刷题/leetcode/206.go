/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    
    p := head

    var pre *ListNode
    for p != nil {
        next := p.Next
        p.Next = pre
        pre = p
        p = next
    }

    return pre
    
}