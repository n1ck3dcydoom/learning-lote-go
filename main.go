package main

import (
	"fmt"
	"learning-lote-go/codetop"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	// sout 5
	fmt.Println(codetop.FindKthLargest(nums, 2))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	// sout 4
	fmt.Println(codetop.FindKthLargest(nums, 4))
}
