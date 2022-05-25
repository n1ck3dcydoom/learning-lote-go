package codetop

import "learning-lote-go/codetop/common"

func MergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &common.ListNode{}
	t := head

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			t.Next = list2
			list2 = list2.Next
		} else {
			t.Next = list1
			list1 = list1.Next
		}
		t = t.Next
	}
	if list1 != nil {
		t.Next = list1
	} else {
		t.Next = list2
	}
	return head.Next
}
