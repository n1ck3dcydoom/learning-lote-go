package main

// import (
// 	"fmt"
// 	"learning-lote-go/leetcode/base"
// )

// func main() {
// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	steps := make([]int, N)
// 	for i := 0; i < N; i++ {
// 		n, _ := fmt.Scan(&steps[i])
// 		if n == 0 {
// 			fmt.Println("error")
// 		}
// 	}

// 	// 直接 dp
// 	// 定义 dp[i] 表示到达第 i 个桩能够走得最大步数
// 	dp := make([]int, N)

// 	// dp 数组全部初始化为 1，因为至少能够走 1 步，即自己本身算 1 步
// 	for i := range dp {
// 		dp[i] = 1
// 	}

// 	// 状态转移方程
// 	// 考虑到达第 i 个桩，扫描前面每个 steps[k] < steps[i] 说明能够从 k 转移到 i
// 	// 有 dp[i] = dp[k] + 1，遍历完成后取最大值
// 	ans := 1
// 	for i := 0; i < N; i++ {
// 		for j := i - 1; j >= 0; j-- {
// 			if steps[j] < steps[i] {
// 				dp[i] = base.MaxInt(dp[i], dp[j]+1)
// 			}
// 		}
// 		ans = base.MaxInt(dp[i], ans)
// 	}
// 	fmt.Println(ans)
// }
