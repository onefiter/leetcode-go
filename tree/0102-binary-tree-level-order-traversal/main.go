package levelorder

import "leetcode-go/template"

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 给定二叉树的根节点 root，按层从上到下、每层从左到右返回节点值，每层一个切片。
//
// 核心思路：
//   - BFS：队列里始终放「当前这一层」的节点，一次处理一整层，同时收集下一层；
//   - DFS：前序遍历时带上层号 depth，把节点值追加到 res[depth]；
//     因为前序先左后右，同一层内仍是从左到右的顺序。

type TreeNode = template.TreeNode

// levelOrderBFS 解法一：BFS 按层处理 ⭐ 最优方案
// 每轮把当前层 cur 的所有节点值收集起来，并把它们的孩子放进 next，然后 cur = next。
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大层宽，最坏 O(n)）。
func levelOrderBFS(root *TreeNode) [][]int {
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
	return res
}

// levelOrderDFS 解法二：DFS 带层号
// 前序遍历（根 -> 左 -> 右），第一次到达某层时新开一个切片，之后把值追加到对应层。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func levelOrderDFS(root *TreeNode) [][]int {
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
	return res
}
