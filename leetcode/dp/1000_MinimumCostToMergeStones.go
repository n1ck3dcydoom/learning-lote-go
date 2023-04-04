package dp

import "math"

func MergeStones(stones []int, k int) int {
	// 首先判断 k 和 n 是否存在解
	// 从 n 堆变为 1 堆减少了 n-1 堆，每次合并会减少 k-1 堆石头
	// 则 (n-1) % (k-1) ==0 则说明有解，其他情况无法保证最终只剩下 1 堆石头
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	// 令 dp[i][j][k] 表示把第 i 堆到第 j 堆石头合并为 k 堆石头所需要的最小花费
	// dp 下标从 0 开始计算，则题目所求的结果为，dp[0][n-1][1]

	// 假设在有解的情况下 n 堆石头，每次合并 k 堆
	// 考虑最后一步：即 最后剩下 k 堆合并成 1 堆，此时的花费是 dp[0][n-1][k] + sum(s[0 ~ n-1])
	// 即拆分为子问题，将[0, n-1] 堆石头合并为 k 堆的最小花费加上合并最后 k 堆的花费
	// 继续考虑最后 k 堆里面的第 1 堆石头，它是怎么合并得到的
	// 1. 如果它本身就是单独的 1 堆，无需合并，即 dp[0][0][1] + dp[1][n-1][k-1]
	// 2. 如果需要通过 1 次合并得到，即 dp[0][k-1][1] + dp[k][n-1][k-1]
	// 3. 如果需要通过 2 次合并得到，即 dp[0][2(k-1)][1] + dp[2(k-1)+1][n-1][k-1]

	// 考虑一般的情况，将 i 到 j 的石头合并为 k 堆
	// 有 dp[i][j][k] = min(dp[i][m][1]+dp[m+1][j][k-1])
	// 其中 m=i+(k-1)x，且 j-(m+1)+1>=k-1 即 m<=j-k+1 即 i+(k-1)x <= j-k+1
	// 最终可得 x <= (j-k)/(k-1)

	// 第一步、定义 dp 数组
	dp := make([][][]int, n)

	// 第二步，初始化
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for p := range dp[i][j] {
				dp[i][j][p] = -1
			}
		}
	}

	// 计算前缀和
	s := make([]int, n+1)
	s[0] = 0
	for i, stone := range stones {
		s[i+1] = s[i] + stone
	}

	// 第三步，状态转移
	var dfs func(i, j, p int) (res int)
	dfs = func(i, j, p int) (res int) {
		// 备忘录剪枝
		memo := &dp[i][j][p]
		if *memo != -1 {
			return *memo
		}
		defer func() {
			*memo = res
		}()
		if p == 1 {
			// 只剩 1 堆石头，无需合并
			if i == j {
				return 0
			}
			// 当 p == 1 时，则无需合并第一堆，直接合并剩下的所有堆
			return dfs(i, j, k) + s[j+1] - s[i]
		}
		// 枚举第 1 堆的所有合并方法
		res = math.MaxInt
		for m := i; m < j; m += k - 1 {
			res = min(res, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}
		return res
	}
	return dfs(0, n-1, 1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
