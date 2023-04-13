package hash

import "math"

func MostFrequentEven(nums []int) int {
	// 出现最多的数，统计计数，哈希计数
	res := math.MaxInt
	freq := 0
	hash := make(map[int]int, 0)
	for _, n := range nums {
		if n%2 != 0 {
			continue
		}
		hash[n]++
		if freq < hash[n] {
			freq = hash[n]
			res = n
		} else if freq == hash[n] && n < res {
			freq = hash[n]
			res = n
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
