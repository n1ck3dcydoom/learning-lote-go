package main

import (
	"fmt"
	"learning-lote-go/codetop/common"
	"learning-lote-go/leetcode/greedy"
)

func main() {
	root := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val: 4,
				Left: &common.TreeNode{
					Val: 8,
				},
				Right: &common.TreeNode{
					Val: 9,
				},
			},
			Right: &common.TreeNode{
				Val: 5,
				Left: &common.TreeNode{
					Val: 10,
				},
				Right: &common.TreeNode{
					Val: 11,
				},
			},
		},
		Right: &common.TreeNode{
			Val: 3,
			Left: &common.TreeNode{
				Val: 6,
			},
			Right: &common.TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(greedy.BtreeGameWinningMove(root, 11, 3))
}
