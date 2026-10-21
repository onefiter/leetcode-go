package levelorderbottom

import (
	"slices"

	"leetcode-go/template"
)

// 107. 二叉树的层序遍历 II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
// 给定二叉树的根节点 root，返回其节点值「自底向上」的层序遍历：
// 从叶子所在的最底层到根所在的层，每层内部仍从左到右。
//
// 核心思路：先按普通层序（自顶向下）收集每一层，最后把层的顺序整体反转即可。
// 不要每层都插到切片头部，那样是 O(层数^2) 的搬移。

type TreeNode = template.TreeNode

// levelOrderBottomBFS 解法一：BFS 按层收集 + 最后反转 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大层宽，不计结果）。
func levelOrderBottomBFS(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		vals := make([]int, 0, len(cur))
		next := []*TreeNode{}
		for _, node := range cur {
			vals = append(vals, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, vals)
		cur = next
	}
	slices.Reverse(res)
	return res
}

// levelOrderBottomDFS 解法二：DFS 带层号 + 最后反转
// 前序 DFS 把值追加到 res[depth]，遍历完后反转层顺序。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，不计结果）。
func levelOrderBottomDFS(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	slices.Reverse(res)
	return res
}
