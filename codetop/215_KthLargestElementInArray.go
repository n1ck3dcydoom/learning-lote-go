package codetop

func FindKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

// 手写快排
func quickSort(nums []int, l int, r int) {
	// 左端点超过了右端点，结束这一趟快排
	if l > r {
		return
	}
	// 以最左侧的元素作为哨兵
	base := nums[l]

	i := l
	j := r
	for i < j {
		// 从右往左查找第一个小于哨兵的元素
		for i < j && nums[j] >= base {
			j--
		}
		// 从左往右查找第一个大于哨兵的元素
		for i < j && nums[i] <= base {
			i++
		}
		// 如果 i 和 j 还没有相遇，交换彼此，直到相遇停止这一趟快排
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 相遇后，和哨兵交换，此时哨兵所处的元素已经排好序
	base, nums[i] = nums[i], base

	// 左右区间分别进行下一趟快排
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}
