package dp

import (
	"math"
	"sort"
)

func MakeArrayIncreasing(arr1 []int, arr2 []int) int {
	// 整体思路，扫描 arr1 ，对于已经扫描过的元素 [0,i] 已经通过操作使其满足严格递增的条件
	// 当扫描完 arr1 之后即可得到整个数组的严格递增
	// 考虑元素 arr1[i] 有两种选择，替换或者不替换 arr1[i]
	// 1. 不替换，此时 arr1[i] > arr1[i-1] 已经满足 [0,i] 元素严格递增，可以不用替换 arr1[i]
	// 2. 替换，此时 arr1[i] <= arr1[i-1] 此时不满足严格递增，需要从 arr2 里面选一个元素替换 arr1[i]
	// 为了保证尽可能少的替换，要从 arr2 里面选择 第一个严格大于 arr1[i-1] 的元素
	// 所以可以对 arr2 做二分查找，找第一个严格大于 target 的元素
	// 如果找不到就返回 -1 说明无法通过替换满足 arr1 的严格递增

	// 注意：对于已经形成严格递增的子序列，选择替换也有可能达成最优解
	// 例如 a1=[1,5,3,4,6] a2=[1,2,3] 对于 a1[1]=5 此时 a1[1]>a1[0] 已经严格递增
	// 但是替换 a1[1]=2 可以达成更优解
	// 对于 [1,5] 在 a2 里面已经找不到大于 5 的元素，会误判为 a1 无法满足严格递增

	// arr1 没有元素或者只有一个元素，不需要任何替换就满足严格递增
	if len(arr1) < 2 {
		return 0
	}

	// arr2 需要二分查找，先排序处理，升序排序
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	// 定义 dp[i][j] 表示前 i 个元素在经过 j 次替换后组成的严格递增子数组的末尾元素的最小值
	// 而且替换只能从 arr2 里面选取元素，最多的替换次数不超过 len(arr2)
	// 对于上述的 [1,5] 这里的最小值显然为 2
	dp := make([][]int, len(arr1)+1)
	for i := range dp {
		// 枚举替换次数的数组为什么要+1
		// 因为要枚举 0 次到 len(arr2)次，总共需要 len(arr2)+1 个元素
		dp[i] = make([]int, len(arr2)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	// 1. 不替换 arr1[i]，则 arr1[i] 一定大于前 i-1 个元素替换若干次后得到的最小末尾元素
	// 即 arr1[i] > dp[i-1][j]，此时 arr1[i] 可以直接加入当前序列，而且替换次数不变
	// 此时 dp[i][j] = arr1[i] (暂时的)

	// 2. 替换 arr1[i]，则替换后的 arr1[i] 一定大于前 i-1 个元素替换 j-1 次后得到的最小末尾元素
	// 此时从 arr2 里面查找满足 arr2[k]>dp[i-1][j-1] 的最小元素即可
	// dp[i][j] = arr2[k] (暂时的)

	// 综上所述，dp[i][j] = min(dp[i][j], arr2[k])

	// 扫描 arr1 需要计算 dp[i-1] 的情况，如果 arr1 从下标 0 开始遍历，会导致数组越界
	// 初始化 dp[0][0] 的情况，让 arr2 从下标 1 开始遍历，避免数组越界
	// 对于 arr1[0] 不进行交换能够得到的严格递增子数组末尾最小元素就是它本身 arr1[0]
	dp[0][0] = -1

	// 扫描 arr1
	for i := 1; i <= len(arr1); i++ {
		// 扫描替换次数，从 0 次到 len(arr2) 次
		// 由于可以从 arr1 和 arr2 里各选一个索引进行替换
		// 扫描次数取 arr1 的遍历索引 i 和 len(arr2) 较小值
		for j := 0; j <= minInt(i, len(arr2)); j++ {
			// 不进行替换
			if arr1[i-1] > dp[i-1][j] {
				dp[i][j] = arr1[i-1]
			}
			// 如果前 i-1 个数能够通过 j-1 次替换满足 arr1 是严格递增子数组
			// 则进行替换，尝试在 arr2 里面找到第一个大于 dp[i-1][j-1] 的元素，贪心策略
			if j-1 >= 0 && dp[i-1][j-1] != math.MaxInt {
				// 如果 t == math.MaxInt 表示在 arr2 里面找不到能够替换的数
				k := search(arr2, dp[i-1][j-1])
				if k != math.MaxInt {
					dp[i][j] = minInt(dp[i][j], arr2[k])
				}
			}

			// 如果已经扫描到最后一个元素，且已经能够满足 arr1 是严格递增数组
			// 返回第一满足 arr1 为严格递增数组的最小 j 值
			if i == len(arr1) && dp[i][j] != math.MaxInt {
				return j
			}
		}
	}
	return -1
}

func search(arr []int, t int) int {
	// 二分查找第一个大于 t 的元素
	// 这里采用蓝红染色二分法
	// 建模第一个大于 t 的元素，左边为蓝色区间 <=，右边为红色区间 >
	// 目标元素划分为包含 = 的区间，即蓝色区间最后一个元素

	// 初始区间，所有元素均为染色
	l, r := -1, len(arr)
	// 循环条件统一使用 l+1 < r
	for l+1 < r {
		mid := l + (r-l)/2
		// 确定划分函数 IsBlue()，如果 mid < t，根据建模可知，此时 mid 应当划分到蓝色
		// 如果 mid == t，根据建模可知 mid 应当划分到蓝色，红色是大于 t 的元素
		// 如果 mid > t，根据建模可知 mid 应当划分到红色
		if arr[mid] > t {
			r = mid
		} else {
			l = mid
		}
	}
	// 处理边界情况，整个区间全为蓝色或者红色
	// 如果整个区间全部为红色，说明每个元素都大于 t，根据建模返回红色区间第一个元素即可
	// 如果整个区间全部为蓝色，说明每个元素都不大于 t，即 <= t，此时不存在大于 t 的元素
	if l == len(arr)-1 {
		return math.MaxInt
	}
	return r
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
