package dp

<<<<<<< HEAD
import "learning-lote-go/leetcode/base"

=======
>>>>>>> c596a6f (Leetcode 300. 最长上升子序列 (LIS 动态规划))
func LengthOfLIS(nums []int) int {
	// LIS 问题，典型的 dp
	n := len(nums)
	// 第一步，定义 dp 数组
	// 定义 dp[i] 表示以 i 结尾的数组最长上升子序列长度
	dp := make([]int, n)

	// 第二步，初始值
	dp[0] = 1

	// 第三步，状态转移方程
	// 考虑一般情况，以 nums[i] 结尾的
	// 1. 如果 nums[i] 可以和前面 nums[k] k<i 组成更长的上升子序列
	// 那么 dp[i] = dp[k] + 1
	// 需要遍历前面所有的子序列 max(dp[k])+1， k<i 且 nums[k] < nums[i]
	// 2. 如果 nums[i] 前面不存在能够组成更长上升子序列的数
	// 那么 dp[i] = 1 (即只有 nums[i] 本身的子序列)

	ans := 1
	// 遍历每个数
	for i := 1; i < n; i++ {
		// 只有当前元素的子序列
		dp[i] = 1
		// 遍历前面的每个数，找到每个能够组成更长上升子序列的位置
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
<<<<<<< HEAD
				dp[i] = base.MaxInt(dp[i], dp[j]+1)
			}
		}
		ans = base.MaxInt(ans, dp[i])
	}
	return ans
}
=======
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		ans = maxInt(ans, dp[i])
	}
	return ans
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
>>>>>>> c596a6f (Leetcode 300. 最长上升子序列 (LIS 动态规划))
