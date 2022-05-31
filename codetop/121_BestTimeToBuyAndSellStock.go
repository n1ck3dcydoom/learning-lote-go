package codetop

import "learning-lote-go/codetop/common"

func MaxProfit(prices []int) int {
	// 动态规划
	n := len(prices)

	// 第一步,定义 dp 数组
	// 定义 dp[i][2]
	// dp[i][0] 表示在第 i 天不持有股票时,手里的钱
	// dp[i][1] 表示在第 i 天持有股票时,手里的钱
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 第二步,初始状态
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 第三步,状态转移
	// 考虑第 i 天的情况
	// 如果第 i 天不持有股票,即 dp[i][0]
	// 若 i-1 没有股票,则 dp[i][0]=dp[i-1][0]
	// 若 i-1 持有股票,则 dp[i][0]=dp[i-1][1]+prices[i],相当于在第 i 天卖出
	// 如果第 i 天持有股票,即 dp[i][1]
	// 若 i-1 没有股票,则 dp[i][0]=-prices[i],相当于在第 i 天买入,手里的钱就等于负的当天股价
	// 若 i-l 持有股票,则 dp[i][1]=dp[i-1][1]
	for i := 1; i < n; i++ {
		// 没有股票
		dp[i][0] = common.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有股票
		dp[i][1] = common.Max(dp[i-1][1], -prices[i])
	}
	// 最后一天不持有股票一定比持有股票时手里的钱多
	return dp[n-1][0]
}
