package bfs

// import (
// 	"container/list"
// 	"learning-lote-go/codetop/common"
// )

// func LargestValues(root *common.TreeNode) []int {
// 	ans := make([]int, 0)
// 	// 使用队列维护 bfs 遍历的中间数据
// 	queue := list.New()
// 	// 头结点入队
// 	queue.PushBack(root)
// 	// bfs 遍历
// 	for queue.Len() != 0 {
// 		// 当前层遍历次数
// 		n := queue.Len()
// 		max := queue.Front().Value.(*common.TreeNode).Val
// 		// 遍历当前层所有节点
// 		for i := 0; i < n; i++ {
// 			node := queue.Front()
// 			// 已经遍历过得节点从队列中移除掉
// 			queue.Remove(node)
// 			max = maxInt(max, node.Value.(*common.TreeNode).Val)
// 			// 左右孩子进入队列
// 			if node.Value.(*common.TreeNode).Left != nil {
// 				queue.PushBack(node.Value.(*common.TreeNode).Left)
// 			}
// 			if node.Value.(*common.TreeNode).Right != nil {
// 				queue.PushBack(node.Value.(*common.TreeNode).Right)
// 			}
// 		}
// 		ans = append(ans, max)
// 	}
// 	return ans
// }

// func maxInt(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }
