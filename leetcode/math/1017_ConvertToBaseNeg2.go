package math

func BaseNeg2(n int) string {
	// 特殊情况
	if n == 0 {
		return "0"
	}
	res := ""
	// 将十进制转换为任意进制，都可以使用短除法
	for n != 0 {
		// 余数
		mod := n % (-2)
		if mod == 0 {
			res = "0" + res
		} else {
			// 余数等于 1 或者 -1 都视为 1
			// 其中余数等于 -1 需要额外处理
			res = "1" + res
		}
		// 当余数等于 -1 时，通过 +2 让余数变为 1
		// 由于余数多了 2，而除数等于 -2 是小于 0 的，所以商需要 +1
		// 这样在还原被除数 = 商*除数+余数时，由于商 +1，相当于多了一个为负数除数，刚好抵消余数多出来的 2
		if mod < 0 {
			n = n/(-2) + 1
		} else {
			n = n / (-2)
		}
	}

	return res
}
