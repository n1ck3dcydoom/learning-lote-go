package greedy

import (
	"learning-lote-go/codetop/common"
)

var ps, ls, rs int

func BtreeGameWinningMove(root *common.TreeNode, n int, x int) bool {
	// 贪心,要尽可能多的染色节点,去除初始的红色节点 x,剩下的分为三个部分
	// x 的父节点以上的其他节点,x 的左孩子,x 的右孩子
	// 蓝色从上面三个部分节点中开始染色,父节点部分能够得到 ps 个,左孩子能够得到 ls 个,右孩子能够得到 rs 个
	// 红色节点总数为 xs,总节点个数为 n,只要有 n-xs > xs 即可获胜
	// 而蓝色节点个数 n-xs 取三个部分的较大值 max(ps,ls,rs)
	// 由于递归遍历过程当中父节点部分的节点个数比较难以计算,采用间接法,先计算 ls 和 rs,用总节点个数 n-ls-rs-1 即可得到 ps

	// 需要先遍历得到左右子树的结果,再对左右子树结果做计算,明显是后序遍历
	var dfs func(root *common.TreeNode) int
	dfs = func(root *common.TreeNode) int {
		// nil 节点终止遍历
		if root == nil {
			return 0
		}
		// 递归遍历左子树
		dls := dfs(root.Left)
		// 递归遍历右子树
		drs := dfs(root.Right)
		// 如果是 x 节点,记录 x 节点的左右子树个数
		if root.Val == x {
			ls = dls
			rs = drs
		}
		return dls + drs + 1
	}
	dfs(root)
	// 判断父节点,左后孩子节点个数,取较大值
	ps = n - ls - rs - 1
	max := max(ps, max(ls, rs))
	// 判断是否能够取胜
	return max > n-max
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
