package dfs

import "strconv"

func DiffWaysToCompute(expression string) []int {
	// 根据运算符直接 dfs 分治搜索运算符左右两侧的计算结果
	res := make([]int, 0)

	runes := []rune(expression)
	n := len(runes)
	for i := 0; i < n; i++ {
		c := runes[i]

		// 运算符将表达式分割为左右两个部分
		if '-' == c || '+' == c || '*' == c {
			exp1 := runes[0:i]
			exp2 := runes[i+1 : n]

			res1 := DiffWaysToCompute(string(exp1))
			res2 := DiffWaysToCompute(string(exp2))

			for _, r1 := range res1 {
				for _, r2 := range res2 {
					if '-' == c {
						res = append(res, r1-r2)
					} else if '+' == c {
						res = append(res, r1+r2)
					} else if '*' == c {
						res = append(res, r1*r2)
					}
				}
			}
		}
	}
	// 纯数字直接加入结果集
	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}
	return res
}
