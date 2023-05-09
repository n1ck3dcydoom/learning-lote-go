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

// 	// 动态规划
// 	// dp[i][j] 表示字符串 s[i,j] 是否是回文子串
// 	// 多初始化一个表示字符串为空的情况
// 	dp := make([][]bool, len(rs)+1)
// 	for i := range dp {
// 		dp[i] = make([]bool, len(rs)+1)
// 	}
// 	// 初始值，对角线上为 true
// 	for i := 0; i <= len(rs); i++ {
// 		dp[i][i] = true
// 		// 考虑到相邻两个元素也是回文子串
// 		if i < len(rs) && rs[i] == rs[i+1] {
// 			dp[i][i+1] = true
// 		}
// 	}
// 	// 状态转移方程
// 	// 如果 s[i] == s[j] 且 s[(i+1),(j-1)] 也是回文子串，则 dp[i][j] = dp[i+1][j-1]
// 	// 这里是 .....******.....
// 	//            ↑l    ↑r，l 和 r 彼此靠近是 l+1 和 r-1
// 	ans := 0
// 	// 遍历子串，0 和 1 已经进行了初始化
// 	for i := 3; i <= len(rs); i++ {
// 	}
// }

// // func main() {
// // 	s := ""
// // 	n, _ := fmt.Scan(&s)
// // 	if n == 0 {
// // 		fmt.Println("error")
// // 	}
// // 	rs := []rune(s)
// // 	ans := 0
// // 	// 双指针查找
// // 	for i := 0; i < len(rs); i++ {
// // 		// 检查奇数长度的子串
// // 		l, r := i-1, i+1
// // 		for l >= 0 && r < len(rs) && rs[l] == rs[r] {
// // 			l--
// // 			r++
// // 		}
// // 		ans = maxInt(ans, r-l-1)
// // 		// 检查偶数长度的子串
// // 		l, r = i, i+1
// // 		for l >= 0 && r < len(rs) && rs[l] == rs[r] {
// // 			l--
// // 			r++
// // 		}
// // 		ans = maxInt(ans, r-l-1)
// // 	}
// // 	fmt.Println(ans)
// // }

// func maxInt(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }
