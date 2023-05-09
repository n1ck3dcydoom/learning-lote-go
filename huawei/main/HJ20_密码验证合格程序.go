package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func main() {
// 	code := ""
// 	n, _ := fmt.Scan(&code)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	// 长度必须超过 8 位
// 	if len(code) <= 8 {
// 		fmt.Println("NG")
// 	} else {
// 		kinds := make([]bool, 4)
// 		rcode := []rune(code)
// 		for i := 0; i < len(rcode); i++ {
// 			if unicode.IsLower(rcode[i]) {
// 				kinds[0] = true
// 			} else if unicode.IsUpper(rcode[i]) {
// 				kinds[1] = true
// 			} else if unicode.IsDigit(rcode[i]) {
// 				kinds[2] = true
// 			} else {
// 				kinds[3] = true
// 			}
// 		}
// 		count := 0
// 		for i := range kinds {
// 			if kinds[i] {
// 				count++
// 			}
// 		}
// 		if count < 3 {
// 			fmt.Println("NG")
// 		} else {
// 			// 检查是否有长度大于 2 的重复子串
// 			// 因为更长的子串一定重复，所以只需要检查是否存在长度为 3 的重复子串
// 			// 直接暴力枚举
// 			flag := true
// 			for i := 0; i <= len(rcode)-3; i++ {
// 				for j := i + 3; j <= len(rcode)-3; j++ {
// 					if rcode[i] == rcode[j] && rcode[i+1] == rcode[j+1] && rcode[i+2] == rcode[j+2] {
// 						flag = false
// 						break
// 					}
// 				}
// 				if !flag {
// 					break
// 				}
// 			}
// 			if flag {
// 				fmt.Println("OK")
// 			} else {
// 				fmt.Println("NG")
// 			}
// 		}
// 	}
// }
