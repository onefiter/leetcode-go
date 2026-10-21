package zigzaglevelorder

import (
	"slices"

	"leetcode-go/template"
)

// 103. 二叉树的锯齿形层序遍历
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
// 给定二叉树的根节点 root，返回其节点值的锯齿形层序遍历：
// 第一层从左到右，第二层从右到左，第三层再从左到右，以此类推。
//
// 核心思路：照常做层序遍历（每层从左到右），只是在奇数层（从 0 计）把这一层的结果反过来；
// 或者直接按下标把值写到正确位置，省掉反转。

type TreeNode = template.TreeNode

// zigzagLevelOrderBFS 解法一：BFS + 按方向直接写入下标 ⭐ 最优方案
// 每层先开一个长度为层宽的切片，从左到右处理节点时：
// 正向层写到下标 i，反向层写到下标 size-1-i，无需额外反转。
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大层宽）。
func zigzagLevelOrderBFS(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	cur := []*TreeNode{root}
	for leftToRight := true; len(cur) > 0; leftToRight = !leftToRight {
		size := len(cur)
		vals := make([]int, size)
		next := []*TreeNode{}
		for i, node := range cur {
			if leftToRight {
				vals[i] = node.Val
			} else {
				vals[size-1-i] = node.Val
			}
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
	return res
}

// zigzagLevelOrderDFS 解法二：DFS 带层号 + 奇数层反转
// 先用前序 DFS 得到普通层序结果（每层从左到右），再把奇数层反转。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，不计结果）。
func zigzagLevelOrderDFS(root *TreeNode) [][]int {
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
	for i := 1; i < len(res); i += 2 {
		slices.Reverse(res[i])
	}
	return res
}
