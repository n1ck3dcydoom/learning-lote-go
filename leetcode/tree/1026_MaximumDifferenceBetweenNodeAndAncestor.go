package tree

import (
	"learning-lote-go/codetop/common"
	"math"
)

func MaxAncestorDiff(root *common.TreeNode) int {
	// 一般情况对于节点 node 要得到以 node 为祖先的所有子树中差值最大的结果
	// 计算 node 的孩子能够得到的最大和最小的差值
	// |node.val - min| 和 |node.val - max| 得到 node 的 min 和 max
	// 很明显需要先遍历子树，然后对子树的遍历结果做处理，采用后序遍历的模板
	_, _, ans := dfs(root)
	return ans
}

func dfs(root *common.TreeNode) (min, max, ans int) {
	if root == nil {
		// 空节点，min 返回最大值，max 返回最小值
		// 这样空节点的父节点做 minInt 和 maxInt 计算的时候会忽略掉空节点的最值
		return math.MaxInt, math.MinInt, ans
	}
	// 叶子结点返回自己的 val 值
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val, ans
	}
	lmin, lmax, lans := dfs(root.Left)
	rmin, rmax, rans := dfs(root.Right)

	// 计算 node 节点的 min 和 max
	min = minInt(lmin, rmin)
	max = maxInt(lmax, rmax)

	// 保存中间结果的最大差值

	ans = maxInt(absInt(root.Val-min), absInt(root.Val-max))
	ans = maxInt(ans, maxInt(lans, rans))
	return minInt(min, root.Val), maxInt(max, root.Val), ans
}

func absInt(ans int) int {
	if ans < 0 {
		return -ans
	}
	return ans
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
