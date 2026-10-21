package minbstdiff

import (
	"math"
	"slices"

	"leetcode-go/template"
)

// 530. 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
// 给一棵 BST（至少两个节点），返回任意两个不同节点值之差的最小绝对值。
//
// 核心思路：BST 中序遍历是升序序列，最小差值一定出现在「相邻」两个元素之间，
// 所以中序遍历时记住上一个值，和当前值做差取最小即可。

type TreeNode = template.TreeNode

// getMinimumDifference 解法一：中序遍历相邻差 ⭐ 最优方案
// 递归中序遍历，prev 记录上一个访问的值，每访问一个节点更新 ans = min(ans, val-prev)。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt
	prev, hasPrev := 0, false
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if hasPrev {
			ans = min(ans, node.Val-prev)
		}
		prev, hasPrev = node.Val, true
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// getMinimumDifferenceIterative 解法二：迭代中序遍历相邻差
// 用显式栈模拟中序遍历，逻辑与解法一相同。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func getMinimumDifferenceIterative(root *TreeNode) int {
	ans := math.MaxInt
	prev, hasPrev := 0, false
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if hasPrev {
			ans = min(ans, cur.Val-prev)
		}
		prev, hasPrev = cur.Val, true
		cur = cur.Right
	}
	return ans
}

// getMinimumDifferenceEnumeration 解法三：收集所有值后排序（对照用）
// 任意遍历收集值，排序后比较相邻元素。没有利用 BST 中序有序的性质。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func getMinimumDifferenceEnumeration(root *TreeNode) int {
	var vals []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	slices.Sort(vals)
	ans := math.MaxInt
	for i := 1; i < len(vals); i++ {
		ans = min(ans, vals[i]-vals[i-1])
	}
	return ans
}
