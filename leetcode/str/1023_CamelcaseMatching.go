package str

import "unicode"

func CamelMatch(queries []string, pattern string) []bool {
	// 驼峰命名法匹配
	ans := make([]bool, 0)
	for _, query := range queries {
		ans = append(ans, match(query, pattern))
	}
	return ans
}

func match(q, p string) bool {
	index := 0
	n := len(p)
	for _, c := range q {
		// index 完成匹配，检查 c 是否是大写
		if index == n {
			if unicode.IsUpper(c) {
				return false
			}
			continue
		}
		if c == rune(p[index]) {
			// 模式串当前字符完成匹配，index++ 继续检查下一位
			index++
		} else {
			// 如果模式串 p[index] != c 考虑 c 是否是大写字母
			if unicode.IsUpper(c) {
				// c 是大写字母，返回 false
				return false
			}
			// c 是小写字母，index 不动，c 后移继续匹配
		}
	}
	// 当 c 完成遍历后，如果 index 还没遍历完 p，说明没完成匹配 false
	if index < n {
		return false
	}
	return true
}
