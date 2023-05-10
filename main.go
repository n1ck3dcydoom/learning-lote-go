package main

import (
	"fmt"
	"learning-lote-go/leetcode/bfs"
)

func main() {
	ans := bfs.UpdateBoard([][]byte{{'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'M', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}}, []int{3, 0})
	for i := range ans {
		for j := range ans[i] {
			fmt.Printf("%c ", ans[i][j])
		}
		if i < len(ans)-1 {
			fmt.Printf("\n")
		}
	}
}
