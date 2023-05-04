package main

// import (
// 	"fmt"
// )

// func main() {
// 	input := make([]int, 50000)
// 	done := false
// 	index := 0

// 	sum := 0
// 	pcount := 0
// 	ncount := 0
// 	for !done {
// 		n, _ := fmt.Scan(&input[index])
// 		if input[index] > 0 {
// 			sum += input[index]
// 			pcount++
// 		} else if input[index] < 0 {
// 			ncount++
// 		}
// 		index++
// 		if n == 0 {
// 			done = true
// 		}
// 	}
// 	fmt.Println(ncount)
// 	res := 0.0
// 	if pcount != 0 {
// 		res = float64(sum) / float64(pcount)
// 	}
// 	fmt.Printf("%.1f", res)
// }
