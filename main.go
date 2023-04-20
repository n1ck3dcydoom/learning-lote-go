package main

import (
	"fmt"
	"learning-lote-go/leetcode/dp"
)

func main() {
	// Output: 1
	fmt.Println(dp.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}))
	// Output: 2
	fmt.Println(dp.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
	// Output: -1
	fmt.Println(dp.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}))
	// Output: 8
	fmt.Println(dp.MakeArrayIncreasing([]int{5, 16, 19, 2, 1, 12, 7, 14, 5, 16}, []int{6, 17, 4, 3, 6, 13, 4, 3, 18, 17, 16, 7, 14, 1, 16}))
}
