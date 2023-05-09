package main

// import (
// 	"fmt"
// )

// func main() {
// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	// 遍历每个数，检查另一个差数是不是素数
// 	res := 9999
// 	l := 0
// 	r := 0
// 	for i := 1; i <= N/2; i++ {
// 		a := i
// 		b := N - i
// 		if check(a) && check(b) {
// 			if b-a < res {
// 				res = b - a
// 				l = a
// 				r = b
// 			}
// 		}
// 	}
// 	fmt.Printf("%d\n%d", l, r)
// }

// func check(n int) bool {
// 	if n != 2 && n%2 == 0 {
// 		return false
// 	}
// 	for i := 2; i <= n/2; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
