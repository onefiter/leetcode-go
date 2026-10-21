package postordertraversal

import (
	"slices"

	"leetcode-go/template"
)

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 给定二叉树的根节点 root，返回它节点值的「后序」遍历：左子树 -> 右子树 -> 根。
//
// 核心思路：
//   - 递归：先递归左、右子树，最后记录根；
//   - 迭代：一路向左压栈，栈顶节点只有在「右子树为空或右子树刚访问完」时才能访问，
//     用 prev 记录上一个访问的节点来判断；
//   - 反转前序：按「根 -> 右 -> 左」做前序，结果反转后正好是「左 -> 右 -> 根」。

type TreeNode = template.TreeNode

// postorderTraversalIterative 解法一：迭代 + 显式栈 + prev 标记 ⭐ 最优方案
// 一路向左压栈；看栈顶 node：
//   - 若 node.Right 为空或 node.Right == prev（右子树刚处理完），弹出并访问，prev = node；
//   - 否则转向右子树继续压栈。
//
// 不需要额外的结果反转，一次遍历直接得到后序。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func postorderTraversalIterative(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == prev {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			prev = node
		} else {
			cur = node.Right
		}
	}
	return res
}

// postorderTraversalRecursive 解法二：递归
// 按「左 -> 右 -> 根」的顺序递归。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，最坏 O(n)）。
func postorderTraversalRecursive(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

// postorderTraversalReversePreorder 解法三：反转「根-右-左」前序
// 用栈做一次「根 -> 右 -> 左」的前序（先压左再压右），最后整体反转即得后序。
// 写法最简单，但需要最后反转一次结果。
// 时间复杂度 O(n)，空间复杂度 O(h)（不计结果切片）。
func postorderTraversalReversePreorder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	slices.Reverse(res)
	return res
}
