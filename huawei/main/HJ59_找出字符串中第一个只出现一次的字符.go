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
// 	hash := make([]rune, 26)

// 	for i := range rs {
// 		hash[rs[i]-'a']++
// 	}
// 	found := false
// 	for i := 0; i < len(rs); i++ {
// 		if hash[rs[i]-'a'] == 1 {
// 			fmt.Println(string(rs[i]))
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		fmt.Println("-1")
// 	}
// }
