package main

// import (
// 	"fmt"
// )

// func main() {
// 	s, l := "", ""
// 	n, _ := fmt.Scan(&s, &l)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	rs := []rune(s)
// 	rl := []rune(l)
// 	hash := make(map[rune]int)
// 	for i := 0; i < len(rs); i++ {
// 		hash[rs[i]] = 1
// 	}
// 	for i := 0; i < len(rl); i++ {
// 		delete(hash, rl[i])
// 	}
// 	if len(hash) == 0 {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }
