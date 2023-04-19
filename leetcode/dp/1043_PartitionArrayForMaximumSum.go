package dp

import "learning-lote-go/leetcode/base"

func maxSumAfterPartitioning(arr []int, k int) int {
	// 定义 dp[i] 表示以 arr[i] 结尾的最大和
	dp := make([]int, len(arr))

	// 初始值
	dp[0] = arr[0]

	// 状态转移
	// 考虑以 arr[i] 结尾的场景
	// 1. arr[i] 和前 0 个构成子数组，dp[i] = dp[i-1]+arr[i]
	// 2. arr[i] 和前 1 个构成子数组，dp[i] = dp[i-2]+2*max(arr[i],arr[i-1])
	// 3. arr[i] 和前 k-1 个构成子数组，dp[i] = dp[i-k-1]+(k+1)*max(arr[i],arr[i-1]...arr[i-k])
	// 综上所述，dp[i] 取 max

	for i := 1; i < len(arr); i++ {
		max := arr[i]
		// 往前枚举 k 种情况
		// i = 5, k = 4, j = 5,4,3,2
		for j := i; j > i-k; j-- {
			max = base.MaxInt(max, arr[j])
			if j == 0 {
				dp[i] = base.MaxInt(dp[i], (i-j+1)*max)
				break
			}
			dp[i] = base.MaxInt(dp[i], dp[j-1]+(i-j+1)*max)
		}
	}

	return dp[len(arr)-1]
}
