package array

import "sort"

func Search(nums []int, target int) int {
	// 二分查找要求原数列有序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 定义左右区间为 [0, n-1] 两端闭区间
	l, r := 0, len(nums)-1
	// 二分查找的难点，其一在于，r 右端点等于 n-1 还是 n
	// 其二在于结束查找条件 l < r 还是 l <= r
	// 其三在于 mid 移动时，l 和 r 端点的变化
	// 结合起来看，就是查找结束时，查找区间内是否还有元素，以及每次查找过滤掉已经查找过的元素
	// 1. [l,r] = [0,n-1]，即查找区间为闭区间
	// 1.1 结束条件 l < r，结束时 l == r，查找区间 [l,l] 还剩下最后一个元素
	// 1.2 结束条件 l <= r，结束时 l == r+1，查找区间 [r+1,r] 为空
	// 2. [l,r) = [0, n)，即查找区间为开区间
	// 2.1 结束条件 l < r，结束时 l == r，查找区间 [l,l) 为空
	// 2.2 结束条件 l <= r，结束时 l == r+1，查找区间 [r+1, r) 为空

	// 对于查找区间 [l,r] 每次计算 mid 之后相当于已经访问过 l 和 r 了
	// 这样 l = mid+1 或者 r = mid-1
	// 对于查找区间 [l,r) 每次计算 mid 之后，l 已经访问过，但是 r 还没有
	// 所以 l = mid+1 或者 r = mid
	for l <= r {
		// 简单防止 l+r 溢出
		mid := l + (r-l)/2
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return -1
}
