package common

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 图节点
type Node struct {
	Val       int
	Neighbors []*Node
}
