package main

// import (
// 	"fmt"
// )

// func main() {
// 	// 哈希计数
// 	s := ""
// 	n, _ := fmt.Scan(&s)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	hash := make(map[rune]int)
// 	rs := []rune(s)
// 	for i := 0; i < len(rs); i++ {
// 		if v, ok := hash[rs[i]]; ok {
// 			hash[rs[i]] = v + 1
// 		} else {
// 			hash[rs[i]] = 1
// 		}
// 	}
// 	fmt.Println(len(hash))
// }
