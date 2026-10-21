package kthsmallestbst

import "leetcode-go/template"

// 230. 二叉搜索树中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
// 给定 BST 的根和整数 k（1 <= k <= n），返回树中第 k 小的值。
//
// 核心思路：BST 的中序遍历就是从小到大的序列，中序访问到的第 k 个节点就是答案。
// 迭代写法可以在数到第 k 个时立刻返回，不必遍历整棵树。

type TreeNode = template.TreeNode

// kthSmallest 解法一：迭代中序遍历 + 提前停止 ⭐ 最优方案
// 用栈模拟中序，每弹出一个节点 k--，减到 0 时即为答案。
// 时间复杂度 O(h + k)，空间复杂度 O(h)。
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1 // k 超出范围（题目保证不会发生）
}

// kthSmallestRecursive 解法二：递归中序遍历
// 递归中序计数，找到后设置标记，后续递归直接返回（剪枝）。
// 时间复杂度 O(h + k)，空间复杂度 O(h)（递归栈）。
func kthSmallestRecursive(root *TreeNode, k int) int {
	res, found := -1, false
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || found {
			return
		}
		dfs(node.Left)
		if found {
			return
		}
		k--
		if k == 0 {
			res, found = node.Val, true
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
