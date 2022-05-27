package codetop

func TwoSum(nums []int, target int) []int {
	// hash 保存另一个数
	var hash = make(map[int]int)
	for i := range nums {
		if j, ok := hash[target-nums[i]]; ok {
			return []int{i, j}
		} else {
			hash[nums[i]] = i
		}
	}
	return []int{}
}
