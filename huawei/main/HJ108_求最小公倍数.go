package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 0
// 	b := 0
// 	n, _ := fmt.Scan(&a, &b)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	if a%b == 0 || b%a == 0 {
// 		fmt.Println(maxInt(a, b))
// 	} else {
// 		max := maxInt(a, b)
// 		min := minInt(a, b)
// 		k := 2
// 		for (max*k)%min != 0 {
// 			k++
// 		}
// 		fmt.Println(max * k)
// 	}
// }

// func maxInt(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func minInt(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
