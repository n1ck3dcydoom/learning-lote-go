package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

}
