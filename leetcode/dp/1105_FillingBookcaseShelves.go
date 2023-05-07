package dp

func MinHeightShelves(books [][]int, shelfWidth int) int {
	// 定义 dp[i] 表示前 i 本书放置到书架后能够得到的最小高度
	dp := make([]int, len(books)+1)

	// 初始化，没有书放到书架的时候，其高度为 0
	dp[0] = 0

	// 状态转移方程
	// 考虑放置第 i 本书时，前面 i-1 本书已经放置完成且得到最小高度 dp[i-1]
	// 第 i 本书的厚度为 w，高度为 h
	// 1. 如果第 i 本书单独放新的一层书架，则 dp[i] = dp[i-1]+h
	// 2. 如果第 i 本书能够和前面若干本书放到同一层，则当前层的高度应该是当前层里面最高的书
	// 第 i 本书和前 j 本书放到同一层，dp[i] = max(books[k]) 0<j<=k<i，且 sum(books[k]) < shelfWidth

	for i := 1; i <= len(books); i++ {

	}
	return -1
}
