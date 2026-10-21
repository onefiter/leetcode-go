package preordertraversal

import "leetcode-go/template"

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// 给定二叉树的根节点 root，返回它节点值的「前序」遍历：根 -> 左子树 -> 右子树。
//
// 核心思路：
//   - 递归：先记录根，再递归左、右子树；
//   - 迭代：栈里先压右孩子再压左孩子，这样左孩子先出栈；
//   - Morris：和中序 Morris 一样建立/断开回连，只是在「第一次到达」节点时就访问它。

type TreeNode = template.TreeNode

// preorderTraversalMorris 解法一：Morris 遍历 ⭐ 最优方案（空间 O(1)）
// 对当前节点 cur：
//   - 没有左子树：访问 cur，走向右子树；
//   - 有左子树：找左子树最右节点 pre。
//     若 pre.Right 为空，说明第一次到达 cur：先访问 cur，再建立回连 pre.Right = cur，进入左子树；
//     若 pre.Right == cur，说明左子树已遍历完：断开回连，走向右子树。
//
// 时间复杂度 O(n)，空间复杂度 O(1)（不计结果切片）。
func preorderTraversalMorris(root *TreeNode) []int {
	res := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			res = append(res, cur.Val) // 第一次到达就访问
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil // 恢复原树
			cur = cur.Right
		}
	}
	return res
}

// preorderTraversalRecursive 解法二：递归
// 按「根 -> 左 -> 右」的顺序递归。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，最坏 O(n)）。
func preorderTraversalRecursive(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// preorderTraversalIterative 解法三：迭代 + 显式栈
// 弹出栈顶访问，然后先压右孩子、再压左孩子（栈后进先出，保证左子树先处理）。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func preorderTraversalIterative(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
