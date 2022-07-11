package codetop

import "learning-lote-go/codetop/common"

func LowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	// 当一个节点 root 的左子树和右子树分别出现了 p 或者 q,那么当前节点 root 就是 p 和 q 的 LCA
	// 当节点 p 或者 q 的子树中出现了 q 或者 p,则 p 或者 q 是 p 和 q 的 LCA

	// 后序遍历从叶子节点往上查找
	// 找到 p 或者 q 直接返回 p 或者 q
	// 遍历到空节点也直接返回
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 判断 left 和 right 的返回结果
	if left == nil && right == nil {
		// 左子树或者右子树都没有找到 p 或者 q,则说明当前节点的子树中没有 p 和 q
		return nil
	} else if left != nil && right == nil {
		// 左子树里面找到了 p 或者 q,右子树没有找到,返回找到的这个节点
		return left
	} else if left == nil && right != nil {
		// 右子树里面找到了 p 或者 q,左子树没有找到,返回找到的这个节点
		return right
	} else {
		// 左右子树分别找到了 p 和 q,此时当前节点 root 就是 p 和 q 的 LCA
		return root
	}
}
