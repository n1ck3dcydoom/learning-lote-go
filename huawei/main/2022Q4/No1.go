package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	inputs := scanner.Text()
// 	pwds := strings.Split(inputs, " ")
// 	// 使用 hash 保存每个密码
// 	hash := make(map[string]int)
// 	for i := 0; i < len(pwds); i++ {
// 		hash[pwds[i]] = len(pwds[i])
// 	}

// 	ans := ""
// 	// 遍历所有密码
// 	for i := 0; i < len(pwds); i++ {
// 		pwd := pwds[i]
// 		// 尝试从末尾截取 pwd，检查 npwd 是否在 hash 当中
// 		found := true
// 		for j := 1; j < len(pwd); j++ {
// 			npwd := pwd[:j]
// 			if _, ok := hash[npwd]; !ok {
// 				// 不存在就停止继续查找
// 				found = false
// 				break
// 			}
// 		}
// 		if found {
// 			curLen := len(pwd)
// 			// 找到更长的密码，覆盖以前的密码
// 			if curLen > len(ans) {
// 				ans = pwd
// 			} else if curLen == len(ans) {
// 				// 长度相同的密码比较字典顺序
// 				slice := []string{ans, pwd}
// 				sort.Strings(slice)
// 				ans = slice[1]
// 			}
// 		}
// 	}
// 	fmt.Print(ans)
// }
