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

// 	// 动态规划
// 	dp := make([]int, N+1)
// 	dp[0] = 1
// 	dp[1] = 1
// 	for i := 2; i <= N; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	fmt.Println(dp[N])

// 	// 递归
// 	// var dfs func(n int) int
// 	// dfs = func(n int) int {
// 	// 	if n == 0 || n == 1 {
// 	// 		return 1
// 	// 	}
// 	// 	return dfs(n-1) + dfs(n-2)
// 	// }
// 	// return dfs(number)
// }
