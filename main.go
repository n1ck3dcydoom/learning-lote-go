package main

import (
	"fmt"
	"learning-lote-go/leetcode/dp"
)

func main() {
	// Output 6
	fmt.Println(dp.MinHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))

	// Output 4
	fmt.Println(dp.MinHeightShelves([][]int{{1, 3}, {2, 4}, {3, 2}}, 6))
}
