package codetop

import "learning-lote-go/codetop/common"

// ReverseList 迭代法
func ReverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	// 遍历节点
	var p = head
	// 前驱结点
	var prep *common.ListNode = nil
	for p != nil {
		// 临时节点保存 p.next
		var tmp = p.Next
		// 反转p
		p.Next = prep
		// 更新 p 和 prep
		prep = p
		p = tmp
	}
	return prep
}

// ReverseList1 递归法
// 递归函数返回反转后链表的新节点
func ReverseList1(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	return reverse(head)
}

func reverse(head *common.ListNode) *common.ListNode {
	if head.Next == nil {
		return head
	}
	var rhead = reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rhead
}
