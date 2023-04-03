package greedy

func prevPermOpt1(arr []int) []int {
	// 只交换一次，且交换后的序列小于交换前，还要满足最大字典序
	// 贪心思路：交换的两个数，pre 一定大于 back
	// 为了保证交换后的数尽可能大，那么交换的两个数就要尽可能的小（靠近右边）
	// 逆序扫描，left = n - 2, right = n - 1
	// 如果 left < right, 那么交换后的数反而更大
	// 所以需要找到第一个 left > right 递减的位置，这样交换 left 后对面前的数影响最小
	// 此时 left 右侧是一个递增序列，还需要逆序找到第一个比 left 大的数 right
	// 为了避免重复 right，需要找到下标最小的 right

	// 枚举 left
	for i := len(arr) - 2; i >= 0; i-- {
		// 找到第一个 left > left + 1 的位置
		// 此时 left 和 left+1 是满足上述条件最小的两个数，这样交换后产生的数也更接近原来的数
		// 且交换后新数比原来的更小（因为较大的高位换成了较小的低位）
		if arr[i] > arr[i+1] {
			j := len(arr) - 1
			// 此时 left 右侧是一个递增序列，可能比 left 更大也可能比 left 更小
			// 为了保证交换后的数字典排序尽可能大且比原来要小
			// 所以需要从末尾找第一个比 left 更小的数 right
			// 为了避免 right 重复出现，还需要找到下标最小的 right
			// 综上所述，要在 [0, n-2] 找最小的 left
			// 要在 [left, n-1] 找最小的 right
			for arr[j] >= arr[i] || arr[j] == arr[j-1] {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}
	return arr
}
