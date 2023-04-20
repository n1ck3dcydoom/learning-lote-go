package main

import (
	"fmt"
	"learning-lote-go/leetcode/dp"
)

func main() {
	// Output: 4
	fmt.Println(dp.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// Output: 4
	fmt.Println(dp.LengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	// Output: 1
	fmt.Println(dp.LengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
