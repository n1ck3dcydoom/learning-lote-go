package dp

import (
	"learning-lote-go/leetcode/base"
)

func LongestArithSeqLength(nums []int) int {
	// 2 <= nums.length <= 1000，0 <= nums[i] <= 500
	// 定义 dp[i][j] 为以 nums[i] 结尾，公差为 j 的等差数组长度
	dp := make([][]int, len(nums))
	for i := range dp {
		// 公差范围在 [-500,500] 对应有 1001 个数
		dp[i] = make([]int, 1001)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}
	// 状态转移方程
	// 考虑以 nums[i] 结尾，公差为 j 的情况
	// 遍历 0~i-1，如果 nums[i] 能够加入到 nums[k] 形成新的等差数列
	// 则 dp[i][nums[i]-nums[k]] = dp[k][num[i]-num[k]]+1
	ans := 1
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			// 将公差 [-500,500] 映射为 [0,1000]
			diff := nums[i] - nums[j] + 500
			dp[i][diff] = base.MaxInt(dp[i][diff], dp[j][diff]+1)
			ans = base.MaxInt(ans, dp[i][nums[i]-nums[j]+500])
		}
	}
	return ans
}
