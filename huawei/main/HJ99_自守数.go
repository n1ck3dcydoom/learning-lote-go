package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	ans := 0
// 	for i := 0; i <= N; i++ {
// 		if checkMath(i) {
// 			fmt.Printf("i=%d and i*i=%d\n", i, i*i)
// 			ans++
// 		}
// 	}
// 	fmt.Println(ans)
// }

// func checkString(n int) bool {
// 	// 使用字符串函数判断是否存在后缀
// 	s := strconv.FormatInt(int64(n), 10)
// 	ss := strconv.FormatInt(int64(n*n), 10)
// 	return strings.HasSuffix(ss, s)
// }

// func checkMath(n int) bool {
// 	// 使用数学方式判断，平方数减去本身后是否能够被 10,100,1000,10000,100000... 整除
// 	if n == 0 || n == 1 {
// 		return true
// 	}
// 	tmp := n*n - n
// 	// n 若是 2 位数，则平方数至少得是 3 位数及以上才能有 2 位数的后缀
// 	if n >= 10000 {
// 		return tmp%100000 == 0
// 	} else if n >= 1000 {
// 		return tmp%10000 == 0
// 	} else if n >= 100 {
// 		return tmp%1000 == 0
// 	} else if n >= 10 {
// 		return tmp%100 == 0
// 	} else {
// 		return tmp%10 == 0
// 	}
// 	return false
// }
