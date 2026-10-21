package balancedtree

import "leetcode-go/template"

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/
// 判断一棵二叉树是否是高度平衡的：每个节点的左右子树高度差的绝对值不超过 1。
//
// 核心思路：平衡需要知道每个节点左右子树的高度。
// 自底向上（后序）计算高度时顺便判断是否平衡，一旦发现不平衡就用 -1 向上传递，
// 这样每个节点只访问一次。

type TreeNode = template.TreeNode

// isBalanced 解法一：自底向上（后序遍历 + 剪枝）⭐ 最优方案
// height(node) 返回子树高度；若子树不平衡返回 -1，上层看到 -1 直接返回 -1。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func isBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := height(node.Left)
		if l == -1 {
			return -1
		}
		r := height(node.Right)
		if r == -1 || l-r > 1 || r-l > 1 {
			return -1
		}
		return max(l, r) + 1
	}
	return height(root) != -1
}

// isBalancedEnumeration 解法二：自顶向下（对照用）
// 对每个节点分别计算左右子树高度判断高度差，再递归判断左右子树。
// 高度会被重复计算：最坏（链状树）每个节点都要重新算一遍下方所有节点的高度。
// 时间复杂度 O(n^2)，空间复杂度 O(h)。
func isBalancedEnumeration(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := depth(root.Left), depth(root.Right)
	if l-r > 1 || r-l > 1 {
		return false
	}
	return isBalancedEnumeration(root.Left) && isBalancedEnumeration(root.Right)
}

// depth 返回子树高度。
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}
