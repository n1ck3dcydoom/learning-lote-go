package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := ""
// 	n, _ := fmt.Scan(&s)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	rs := []rune(s)
// 	hash := make(map[rune]int)
// 	less := 1000
// 	for i := range rs {
// 		if v, ok := hash[rs[i]]; ok {
// 			hash[rs[i]] = v + 1
// 		} else {
// 			hash[rs[i]] = 1
// 		}
// 	}
// 	for _, v := range hash {
// 		if v < less {
// 			less = v
// 		}
// 	}
// 	ans := make([]rune, 0)
// 	for i := range rs {
// 		if hash[rs[i]] == less {
// 			continue
// 		} else {
// 			ans = append(ans, rs[i])
// 		}
// 	}
// 	fmt.Println(string(ans))
// }
