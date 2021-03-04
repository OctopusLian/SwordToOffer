/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	var re []int //正序
	var er []int //逆序
	for head != nil {
		re = append(re, head.Val) //从头到尾记录链表的每个节点的值
		head = head.Next
	}
	for i := len(re) - 1; i >= 0; i-- {
		er = append(er, re[i]) //逆序记录链表的每个节点的值
	}
	return er //返回
}