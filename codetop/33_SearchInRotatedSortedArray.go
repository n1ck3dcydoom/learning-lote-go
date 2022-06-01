package codetop

func Search(nums []int, target int) int {
	// 二分查找
	len := len(nums)

	// 查找区间[l,r]
	l := 0
	r := len - 1

	// 结束时 l=r+1,查找区间为[r+1,r],区间为空
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		// 旋转后原有序数组形成左区间和右区间两个部分,左区间的左端点比右区间的所有值都大
		// 判断 mid 落在左区间还是右区间
		// 然后在判断 target 值和 mid 的关系得出 target 在左区间还是右区间

		if nums[mid] >= nums[l] {
			// mid落在左区间
			if target >= nums[l] && nums[mid] > target {
				// target 比左区间最小值大,且 mid 也比 target 大
				// 则 target 一定在左区间,直接收缩右端点为 mid-1
				r = mid - 1
			} else {
				// target 比左区间最小值更小,则说明 target 落到了右区间,直接收缩左端点为 mid+1
				// target 大于左区间最小值,但是大于 mid,则说明 target 落到了左区间,直接收缩左端点为 mid+1 (下一步 mid 更大)
				l = mid + 1
			}
		} else {
			// mid落在右区间
			if target <= nums[r] && target > nums[mid] {
				// target比右区间最大值还要小,则说明 target 一定在右区间
				// 且 target 大于 mid,直接收缩左端点为 mid+1
				l = mid + 1
			} else {
				// target 比右区间最大值更大,则说明 target 落到了左区间,直接收缩右端点为 mid-1
				// target 小于右区间最大值,但是大于 mid,则说明 target 落到了右区间,直接收缩右端点为 mid-1 (下一步 mid 更小)
				r = mid - 1
			}
		}
	}
	return -1
}
