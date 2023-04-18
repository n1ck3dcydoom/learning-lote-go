package main

import (
	"fmt"
	"learning-lote-go/codetop/common"
	"learning-lote-go/leetcode/tree"
)

func main() {
	root1 := &common.TreeNode{
		Val: 8,
		Left: &common.TreeNode{
			Val: 3,
			Left: &common.TreeNode{
				Val: 1,
			},
			Right: &common.TreeNode{
				Val: 6,
				Left: &common.TreeNode{
					Val: 4,
				},
				Right: &common.TreeNode{
					Val: 7,
				},
			},
		},
		Right: &common.TreeNode{
			Val: 10,
			Right: &common.TreeNode{
				Val: 14,
				Left: &common.TreeNode{
					Val: 13,
				},
			},
		},
	}
	fmt.Println(tree.MaxAncestorDiff(root1))

	root2 := &common.TreeNode{
		Val: 1,
		Right: &common.TreeNode{
			Val: 2,
			Right: &common.TreeNode{
				Val: 0,
				Left: &common.TreeNode{
					Val: 3,
				},
			},
		},
	}
	fmt.Println(tree.MaxAncestorDiff(root2))
}
