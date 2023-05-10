package main

// import (
// 	"fmt"
// )

// func main() {
// 	s, t := "", ""
// 	n, _ := fmt.Scan(&s, &t)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	// 双指针遍历查找
// 	rs := []rune(s)
// 	rt := []rune(t)
// 	i, j := 0, 0
// 	for i < len(rs) && j < len(rt) {
// 		if rs[i] == rt[j] {
// 			i++
// 			j++
// 		} else {
// 			j++
// 		}
// 	}
// 	if i == len(rs) {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }
