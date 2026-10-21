package balancebst

import (
	"slices"

	"leetcode-go/template"
)

// 1382. 将二叉搜索树变平衡
// https://leetcode.cn/problems/balance-a-binary-search-tree/
// 给一棵二叉搜索树，返回一棵节点值相同的「平衡」二叉搜索树
// （每个节点左右子树高度差不超过 1）。答案可能不唯一。
//
// 核心思路：BST 的中序遍历是有序序列。
// 先中序遍历得到有序节点列表，再每次取中间元素作为根，
// 左半边递归建左子树、右半边递归建右子树——左右规模最多差 1，自然平衡。

type TreeNode = template.TreeNode

// balanceBST 解法一：中序遍历收集节点 + 取中点重建（复用原节点）⭐ 最优方案
// 中序遍历把节点指针按顺序存进数组；build(l, r) 取 mid 作为根，
// 递归挂上左右子树。直接复用原节点，不新建节点。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func balanceBST(root *TreeNode) *TreeNode {
	var nodes []*TreeNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nodes = append(nodes, node)
		inorder(node.Right)
	}
	inorder(root)

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		node := nodes[mid]
		node.Left = build(l, mid-1)
		node.Right = build(mid+1, r)
		return node
	}
	return build(0, len(nodes)-1)
}

// balanceBSTEnumeration 解法二：任意遍历收集值 + 排序 + 新建节点（对照用）
// 不利用 BST 中序有序的性质：先前序遍历收集所有值，排序后按中点新建一棵树。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func balanceBSTEnumeration(root *TreeNode) *TreeNode {
	var vals []int
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	slices.Sort(vals)

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		return &TreeNode{Val: vals[mid], Left: build(l, mid-1), Right: build(mid+1, r)}
	}
	return build(0, len(vals)-1)
}
