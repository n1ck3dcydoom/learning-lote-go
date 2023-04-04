package main

import (
	"fmt"
	"learning-lote-go/leetcode/dp"
)

func main() {
	stones := []int{3, 2, 4, 1}
	fmt.Println(dp.MergeStones(stones, 2))
}
