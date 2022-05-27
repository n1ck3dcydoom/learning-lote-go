package codetop

import (
	"learning-lote-go/codetop/common"
)

func LevelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// 使用队列层序遍历二叉树
	var p []*common.TreeNode

	// 头节点入队
	p = append(p, root)

	var res [][]int
	// 遍历队列
	for i := 0; len(p) > 0; i++ {
		// 当前层的队列
		var q []*common.TreeNode
		res = append(res, []int{})
		for j := 0; j < len(p); j++ {
			node := p[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
	}
	return res
}
