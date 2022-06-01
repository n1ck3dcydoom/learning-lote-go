package codetop

import "learning-lote-go/codetop/common"

func HasCycle(head *common.ListNode) bool {
	// 快慢指针
	// 快指针每次走 2 步,慢指针每次走 1 步
	// 如果有环,则快慢指针一定会相遇
	// 如果无环,快指针最终会指向 nil

	f := head
	s := head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if s == f {
			return true
		}
	}
	return false
}
