package main

// import (
// 	"fmt"
// )

// func main() {
// 	// m
// 	s1 := ""
// 	// n
// 	s2 := ""
// 	n, _ := fmt.Scan(&s1, &s2)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	// 我他妈直接动态规划
// 	// dp[i][j] 表示，以 s1[i] s2[j] 结尾的最长公共子串长度
// 	dp := make([][]int, 0)
// 	M := len(s1)
// 	N := len(s2)
// 	// 多初始化 1 行 1 列，用于表示 s1 或者 s2 为空串的情况
// 	// s1 或者 s2 为空串的情况都初始化为 0
// 	for i := 0; i <= M; i++ {
// 		row := make([]int, N+1)
// 		dp = append(dp, row)
// 	}
// 	// 状态转移方程
// 	// 考虑一般情况以 s1[i] s2[j] 结尾的时候
// 	// 如果，s1[i] == s2[j] 则 dp[i][j] 可以通过 dp[i-1][j-1] 转移，即 dp[i][j] = dp[i-1][j-1]+1
// 	ans := 0
// 	for i := 1; i <= M; i++ {
// 		for j := 1; j <= N; j++ {
// 			if s1[i-1] == s2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 				maxInt := func(a, b int) int {
// 					if a < b {
// 						return b
// 					}
// 					return a
// 				}
// 				ans = maxInt(ans, dp[i][j])
// 			}
// 		}
// 	}
// 	fmt.Println(ans)
// }
