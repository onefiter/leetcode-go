package inordertraversal

import "leetcode-go/template"

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 给定二叉树的根节点 root，返回它节点值的「中序」遍历：左子树 -> 根 -> 右子树。
//
// 核心思路：
//   - 递归最直观：先递归左子树，再记录根，再递归右子树；
//   - 迭代用显式栈模拟递归：一路向左压栈，弹出时访问，再转向右子树；
//   - Morris 遍历利用叶子节点空闲的右指针临时「回连」到中序后继，省掉栈，空间 O(1)。

type TreeNode = template.TreeNode

// inorderTraversalMorris 解法一：Morris 遍历 ⭐ 最优方案（空间 O(1)）
// 对当前节点 cur：
//   - 没有左子树：直接访问 cur，走向右子树；
//   - 有左子树：找左子树最右节点 pre（即 cur 的中序前驱）。
//     若 pre.Right 为空，说明第一次来到 cur，令 pre.Right = cur 建立回连，然后进入左子树；
//     若 pre.Right == cur，说明左子树已遍历完，断开回连恢复原树，访问 cur，走向右子树。
//
// 遍历结束后树结构被完全恢复。
// 时间复杂度 O(n)（每条边最多走两次），空间复杂度 O(1)（不计结果切片）。
func inorderTraversalMorris(root *TreeNode) []int {
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
			pre.Right = cur // 建立回连
			cur = cur.Left
		} else {
			pre.Right = nil // 断开回连，恢复原树
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// inorderTraversalRecursive 解法二：递归
// 按「左 -> 根 -> 右」的顺序递归。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，h 为树高，最坏 O(n)）。
func inorderTraversalRecursive(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// inorderTraversalIterative 解法三：迭代 + 显式栈
// 一路向左把节点压栈；走到底后弹出栈顶访问，再转向它的右子树重复该过程。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func inorderTraversalIterative(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}
