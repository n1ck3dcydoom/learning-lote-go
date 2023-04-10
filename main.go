package main

import (
	"fmt"
	"learning-lote-go/codetop/common"
	"learning-lote-go/leetcode/stack"
)

func main() {
	head := &common.ListNode{
		Val: 2,
		Next: &common.ListNode{
			Val: 7,
			Next: &common.ListNode{
				Val: 4,
				Next: &common.ListNode{
					Val: 3,
					Next: &common.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := stack.NextLargerNodes(head)
	for n := range res {
		fmt.Println(n)
	}
}
