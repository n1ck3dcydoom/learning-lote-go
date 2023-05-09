package main

// import (
// 	"fmt"
// 	"strconv"
// 	"unicode"
// )

// func main() {
// 	s := ""
// 	n, _ := fmt.Scan(&s)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	// 使用两个栈模拟运算
// 	s1 := make([]int, 100)
// 	st1 := 0
// 	s2 := make([]rune, 100)
// 	st2 := 0

// 	// 遍历算数
// 	rs := []rune(s)
// 	for i := 0; i < len(rs); i++ {
// 		// * 和 / 高级运算，和栈顶元素运算后再将结果入栈
// 		if rs[i] == '*' || rs[i] == '/' {
// 			a := s1[st1]
// 			st1--
// 			j:=i+1
// 			for ;j<len(rs);j++{
// 				if unicode.IsDigit()
// 			}
// 		}
// 	}
// }
