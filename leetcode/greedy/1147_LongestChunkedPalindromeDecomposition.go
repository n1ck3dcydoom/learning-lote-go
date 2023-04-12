package greedy

func LongestDecomposition(text string) int {
	// 最长的段式回文
	// 双指针从两端往中间查找，找到最短的字串后，分割两端子串，直到双指针相遇
	// (abc)dabc xxxxx abcd(abc)
	// 已经找到当前最短的子串 abc 之后，就可以截断原串没必要继续查找更长的子串了
	// 因为 abcdabc xxxxx abcdabc 可以继续划分为 (abc)(d)(abc) xxxxx (abc)(d)(abc)

	k := 0
	// 索引下标从 1 开始遍历到 n，为了保证切割字符串 [:i] 右区间可以被排除在外
	for text != "" {
		i, n := 1, len(text)
		// 枚举左右子串
		for i <= n/2 && text[:i] != text[n-i:] {
			i++
		}
		// 长度超过一半的子串都枚举完了，自然不可能分割为两个子串，结束枚举s
		if i > n/2 {
			// 此时只能以自己本身分段
			k++
			break
		}
		// 分割左右子串
		text = text[i : n-i]
		// 分段长度 +2
		k += 2
	}
	return k
}
