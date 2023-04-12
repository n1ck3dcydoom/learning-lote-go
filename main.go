package main

import (
	"fmt"
	"learning-lote-go/leetcode/greedy"
)

func main() {
	fmt.Println(greedy.LongestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
	fmt.Println(greedy.LongestDecomposition("aaa"))
	fmt.Println(greedy.LongestDecomposition("elvtoelvto"))
}
