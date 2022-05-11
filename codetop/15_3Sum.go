package codetop

import "sort"

// threeSum 哈希算法会导致重复结果出现，去重非常复杂
func threeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	nmap := make(map[int]int, n)
	for _, num := range nums {
		nmap[num] = num
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := 0 - nums[i] - nums[j]
			if _, ok := nmap[k]; ok {
				res = append(res, []int{nums[i], nums[j], k})
			}
		}
	}
	return res
}

// ThreeSum 排序+双指针查找
func ThreeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	// 排序
	sort.Ints(nums)
	// 指针 i 用于遍历数组
	// 指针 j,k 双指针查找满足条件的三元组，同时需要两端去重
	// i 只需要遍历到倒数第三个就行，n-1 和 n-2 都不需要遍历，
	for i := 0; i < n-2; i++ {
		// 对 i 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 找到满足条件的三元组之后需要对 jk 双端去重
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				// 找到满足的 i,j,k，且 j 和 k 完成去重后，还需要继续移动 jk
				// 因为去重后的 j 和 k 是最后一次重复出现的位置下一趟查找需要从下一个位置开始
				j++
				k--
			} else if sum > 0 {
				// 没有满足条件的三元组，则利用数组排序的性质，大于 0 就移动右端点，小于 0 就移动左端点 l
				k--
			} else {
				j++
			}
		}
	}
	return res
}
