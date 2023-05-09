package main

// import (
// 	"fmt"
// )

// func main() {
// 	word := ""
// 	n, _ := fmt.Scan(&word)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	fmt.Println(longestBeautifulSubstring(word))
// }

// func longestBeautifulSubstring(word string) int {
// 	// 必须以 a 开头，且按照 aeiou 排序
// 	// 滑动窗口，检查每个以 a 开头的区间，判断区间的最大值
// 	l := 0
// 	ans := 0
// 	var array = []rune{'a', 'e', 'i', 'o', 'u'}
// 	index := 0
// 	rword := []rune(word)

// 	for l < len(rword) {
// 		if word[l] == 'a' {
// 			r := l + 1
// 			for r < len(rword) {
// 				// 什么情况下右端点 r 可以扩大?
// 				if array[index] == rword[r] {
// 					// 1. r 指向的字符和当前元音字符相等
// 					r++
// 				} else if index <= 3 && array[index+1] == rword[r] {
// 					// 2. r 指向的字符是下一个元音字符，最后一个元音字符不存在下一个
// 					r++
// 					index++
// 				} else {
// 					// 停止寻找 r
// 					break
// 				}
// 			}
// 			if index == 4 {
// 				var maxInt = func(a, b int) int {
// 					if a < b {
// 						return b
// 					}
// 					return a
// 				}
// 				ans = maxInt(ans, r-l)
// 			}
// 			l = r
// 		} else {
// 			l++
// 		}
// 		index = 0
// 	}
// 	return ans
// }
