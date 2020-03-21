/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var val int
	var root, preNode *ListNode
	for true {
		if l1 == nil && l2 == nil && val <= 0 {
			break
		}
		if l1 != nil {
			val += l1.Val
            l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
            l2 = l2.Next
		}
		curNode := ListNode{val%10, nil}
		val /= 10

		if(preNode == nil) {
			preNode = &curNode
			root = preNode
			continue
		}

		preNode.Next = &curNode
		preNode = &curNode
	}

	return root
}