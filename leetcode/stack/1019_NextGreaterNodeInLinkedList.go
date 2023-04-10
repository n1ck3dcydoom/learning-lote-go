package stack

import "learning-lote-go/codetop/common"

func NextLargerNodes(head *common.ListNode) []int {
	// 下一个更大/小的数，一眼单调栈
	var res []int
	// 链表下标从 0 开始计算
	var index = -1
	// 遍历链表时，不像数组那样方便的拿到元素下标，需要额外字段保存遍历的下标索引
	// stack[i][0] 保存元素值, stack[i][1] 保存链表下标
	var stack [][]int
	for head != nil {
		// 修改遍历下标索引
		index++
		res = append(res, 0)
		// 栈顶指针 len(stack)-1
		// 判断单调栈栈顶元素和当前元素大小关系
		for len(stack) > 0 && stack[len(stack)-1][0] < head.Val {
			// 弹出的栈顶元素
			top := stack[len(stack)-1]
			// 截取栈，左闭右开区间，相当于去掉最后一个元素即去掉栈顶元素
			stack = stack[:len(stack)-1]
			res[top[1]] = head.Val
		}
		// 入栈，保存元素值和元素下标
		node := []int{head.Val, index}
		stack = append(stack, node)
		head = head.Next
	}
	return res
}
