/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		// l1 + l2 + c 都会走完这么长的距离，走完时，肯定会汇聚于一点
		// l1 + c + l2
		// l2 + c + l1
		// 相交了
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}