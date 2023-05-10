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
// 	// 直接暴力计算
// 	a := make([]string, 1)
// 	a[0] = "1"
// 	for i := 1; i <= N; i++ {
// 		rpre := []rune(a[i-1])
// 		rcur := make([]rune, 0)
// 		// 解释前一项
// 		c := rpre[0]
// 		cCount := 1
// 		for j := 1; j < len(rpre); j++ {
// 			// 出现相同字符
// 			if rpre[j-1] == rpre[j] {
// 				cCount++
// 			} else {
// 				// 出现其他字符，翻译当前字符
// 				rcur = append(rcur, rune(cCount)+'0', c)
// 				c = rpre[j]
// 				cCount = 1
// 			}
// 		}
// 		rcur = append(rcur, rune(cCount)+'0', c)
// 		a = append(a, string(rcur))
// 	}
// 	fmt.Println(a[N])
// }
