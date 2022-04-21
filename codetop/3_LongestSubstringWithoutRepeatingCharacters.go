package codetop

func LengthOfLongestSubstring(s string) int {
	if "" == s {
		return 0
	}
	// 记录字符最后一次出现的位置
	var hash = make(map[byte]int)
	var l, res = 0, 0
	for r := 0; r < len(s); r++ {
		_, contains := hash[s[r]]
		if contains {
			// 如果出现过，则更新左端点
			// 判断当前左端点 l 和最后一次出现的位置 pos，去二者的较大值
			l = max(l, hash[s[r]]+1)
		}
		// 更新字符出现的位置
		hash[s[r]] = r
		// 更新最长不重复子区间长度
		res = max(res, r-l+1)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
