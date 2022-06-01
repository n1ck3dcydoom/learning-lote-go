package codetop

import (
	"learning-lote-go/codetop/common"
)

func ZigzagLevelOrder(root *common.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	// 双端队列
	queue := make([]*common.TreeNode, 0)
	// 头节点入队
	queue = append(queue, root)

	odd := 1
	for len(queue) > 0 {
		len := len(queue)
		list := make([]int, len)
		for i := 0; i < len; i++ {
			// 头节点出队
			node := queue[0]
			queue = queue[1:]
			if odd == 1 {
				list[i] = node.Val
			} else {
				list[len-i-1] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		odd *= -1
		res = append(res, list)
	}
	return res
}
