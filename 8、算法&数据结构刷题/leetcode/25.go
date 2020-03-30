/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    
    sign := true
    newHead := head
    var preEnd *ListNode
    startP := head

    for true {
        endKp := startP
        for i:=1; i<k && endKp != nil; i++ {
            endKp = endKp.Next
        }
        if endKp == nil {
            break
        }

        nextKp := endKp.Next
        endKp.Next = nil

        np := reverse(startP)

        if preEnd != nil {
            preEnd.Next = np
        }
        preEnd = startP
        
        if (sign) {
            sign = false
            newHead = np
        }

        startP = nextKp
    }
    if startP != nil {
        preEnd.Next = startP
    }

    return newHead
}

// 返回队列首
func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}