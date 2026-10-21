package deletebstnode

import "leetcode-go/template"

// 450. 删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/
// 给定 BST 的根和一个值 key，删除值为 key 的节点并返回（可能变化的）根。
//
// 核心思路：先按 BST 性质找到目标节点，再分三种情况删除：
//   - 没有孩子：直接删掉（返回 nil）；
//   - 只有一个孩子：用这个孩子顶替它；
//   - 两个孩子都有：找右子树的最小节点（中序后继）来顶替它。

type TreeNode = template.TreeNode

// deleteNode 解法一：递归 ⭐ 最优方案
// key 比根小去左子树删，比根大去右子树删，相等时按三种情况处理。
// 两个孩子都有时：把左子树整体挂到右子树最左节点的左边，返回右子树作为新根。
// 时间复杂度 O(h)，空间复杂度 O(h)（递归栈），h 为树高。
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	default:
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 找右子树最小节点，把左子树挂到它左边。
		succ := root.Right
		for succ.Left != nil {
			succ = succ.Left
		}
		succ.Left = root.Left
		return root.Right
	}
	return root
}

// deleteNodeIterative 解法二：迭代
// 用循环找到目标节点 cur 及其父节点 parent；
// 若 cur 有两个孩子，找到中序后继 succ，把 succ 从原位置摘下并替换 cur；
// 最后把替换后的子树接回 parent。
// 时间复杂度 O(h)，空间复杂度 O(1)。
func deleteNodeIterative(root *TreeNode, key int) *TreeNode {
	var parent *TreeNode
	cur := root
	for cur != nil && cur.Val != key {
		parent = cur
		if key < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return root // 没找到
	}
	var replace *TreeNode
	switch {
	case cur.Left == nil:
		replace = cur.Right
	case cur.Right == nil:
		replace = cur.Left
	default:
		// 找中序后继 succ 及其父节点 sp。
		sp, succ := cur, cur.Right
		for succ.Left != nil {
			sp, succ = succ, succ.Left
		}
		if sp != cur {
			sp.Left = succ.Right // 摘下 succ
			succ.Right = cur.Right
		}
		succ.Left = cur.Left
		replace = succ
	}
	if parent == nil {
		return replace
	}
	if parent.Left == cur {
		parent.Left = replace
	} else {
		parent.Right = replace
	}
	return root
}
