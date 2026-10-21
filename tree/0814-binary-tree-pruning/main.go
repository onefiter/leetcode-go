package pruning

import "leetcode-go/template"

// 814. 二叉树剪枝
// https://leetcode.cn/problems/binary-tree-pruning/
// 二叉树的节点值只有 0 和 1。删掉所有「不包含 1」的子树，返回剪枝后的根。
//
// 核心思路：一个节点要不要删，取决于它的整棵子树里有没有 1，
// 所以必须先处理完左右子树再决定自己——这是典型的后序遍历。
// 剪完左右子树后，如果当前节点是 0 且已经没有孩子，那它整棵子树都不含 1，删掉。

type TreeNode = template.TreeNode

// pruneTree 解法一：后序递归 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

// pruneTreeEnumeration 解法二：自顶向下，每个节点单独判断子树是否含 1（对照用）
// 对每个节点，先 O(子树大小) 地检查左/右子树是否含 1，不含就直接剪掉，再递归处理保留的孩子。
// 时间复杂度 O(n*h)（最坏 O(n^2)），空间复杂度 O(h)。
func pruneTreeEnumeration(root *TreeNode) *TreeNode {
	if !containsOne(root) {
		return nil
	}
	if !containsOne(root.Left) {
		root.Left = nil
	} else {
		pruneTreeEnumeration(root.Left)
	}
	if !containsOne(root.Right) {
		root.Right = nil
	} else {
		pruneTreeEnumeration(root.Right)
	}
	return root
}

func containsOne(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Val == 1 || containsOne(node.Left) || containsOne(node.Right)
}
