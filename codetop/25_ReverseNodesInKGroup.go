package codetop

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 每次循环 k 个节点，找到新的 t
	var t *ListNode = head
	for i := 0; i < k; i++ {
		// 如果剩下的不够 k 个节点，不需要翻转直接返回 head
		if t == nil {
			return head
		}
		// 找到新的 t 节点，注意这里是左闭右开区间，即翻转的是[h,t)
		t = t.Next
	}
	// 翻转当前 k 个节点，得到新的头结点nh
	var nh *ListNode = reverse25(head, t)
	// 让当前 head 节点指向下一次递归翻转返回后的新头结点
	head.Next = ReverseKGroup(t, k)
	// 返回翻转后的新头结点
	return nh
}

func reverse25(h *ListNode, t *ListNode) *ListNode {
	// 翻转 h 和 t 之间的节点，左闭右开区间 [h,t)
	var p *ListNode = nil
	var n *ListNode = nil

	for h != t {
		// 保存 h.next
		n = h.Next
		// 翻转 h
		h.Next = p
		// 移动 p 和 h
		p = h
		h = n
	}
	return p
}
