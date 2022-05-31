package codetop

import "learning-lote-go/codetop/common"

func MaxSubArray(nums []int) int {
	// 动态规划
	// 第一步、定义 dp[i] 表示以 i 结尾的子序列的最大和
	var dp []int = make([]int, len(nums)+1)

	// 第二步、初始状态
	dp[0] = 0
	dp[1] = nums[0]

	// 第三步、状态转移
	// 考虑以 i 结尾的时候
	// 如果 dp[i-1]>0，则 dp[i-1]+nums[i] > nums[i]，dp[i] = dp[i-1] + nums[i]
	// 如果 dp[i-1]<0，则 dp[i-1]+nums[i] < nums[i]，dp[i] = nums[i]
	res := 0
	for i := 1; i < len(nums); i++ {
		if dp[i-1] >= 0 {
			dp[i] = dp[i-1] + nums[i-1]
		} else {
			dp[i] = nums[i-1]
		}
		res = common.Max(res, dp[i])
	}
	return res
}
