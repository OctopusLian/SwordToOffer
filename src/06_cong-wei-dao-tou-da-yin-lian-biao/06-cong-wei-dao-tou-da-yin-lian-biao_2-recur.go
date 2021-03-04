/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	res := reversePrint(head.Next)
	return append(res, head.Val)
}