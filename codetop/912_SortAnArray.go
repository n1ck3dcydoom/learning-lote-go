package codetop

func SortArray(nums []int) []int {

	// 快速排序
	quickSort1(nums, 0, len(nums)-1)
	return nums
}

func bubbleSort(nums []int) {
	n := len(nums)
	// n 个数需要 n 趟排序
	for i := 0; i < n; i++ {
		// 每次从第一个
	}
}

func quickSort1(nums []int, l, r int) {

	if l > r {
		return
	}
	i := l
	j := r
	sentinel := nums[l]
	for i < j {
		// 先 从右往左寻找第一个比哨兵小的数
		for i < j && nums[j] >= sentinel {
			j--
		}
		// 再 从左往右寻找第一个比哨兵大的数
		for i < j && nums[i] <= sentinel {
			i++
		}
		// 还没相遇就交换
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// ij 相遇之后，说明 i 左边的数都小于等于哨兵，j 右边的数都大于等于哨兵
	// 此事哨兵的正确位置就找到了
	nums[l], nums[i] = nums[i], nums[l]

	// 从哨兵开始对左右两侧递归排序
	quickSort1(nums, l, i-1)
	quickSort1(nums, i+1, r)
}
