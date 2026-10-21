package distancek

import (
	"slices"

	"leetcode-go/template"
)

// 863. 二叉树中所有距离为 K 的结点
// https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/
// 给一棵二叉树（值互不相同）、一个目标节点 target 和整数 k，
// 返回所有与 target 距离恰好为 k 的节点值（顺序任意）。
//
// 核心思路：树里只能往下走，但距离可以「往上再往下」。
// 先 DFS 一遍记录每个节点的父节点，就把树变成了一张无向图（左、右、父三个邻居），
// 再从 target 出发做 BFS，走 k 步时队列里的节点就是答案。

type TreeNode = template.TreeNode

// distanceK 解法一：父指针哈希 + BFS ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(n)。
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := map[*TreeNode]*TreeNode{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		for _, child := range []*TreeNode{node.Left, node.Right} {
			if child != nil {
				parent[child] = node
				dfs(child)
			}
		}
	}
	if root == nil || target == nil {
		return []int{}
	}
	dfs(root)

	visited := map[*TreeNode]bool{target: true}
	queue := []*TreeNode{target}
	for step := 0; step < k && len(queue) > 0; step++ {
		var next []*TreeNode
		for _, node := range queue {
			for _, nb := range []*TreeNode{node.Left, node.Right, parent[node]} {
				if nb != nil && !visited[nb] {
					visited[nb] = true
					next = append(next, nb)
				}
			}
		}
		queue = next
	}
	res := make([]int, 0, len(queue))
	for _, node := range queue {
		res = append(res, node.Val)
	}
	return res
}

// distanceKEnumeration 解法二：对每个节点求它到 target 的距离（对照用）
// 先求 root 到每个节点的路径，两节点距离 = depth(a) + depth(b) - 2*depth(LCA)，
// LCA 的深度就是两条根路径的最长公共前缀长度 - 1。逐个节点判断是否等于 k。
// 时间复杂度 O(n*h)，空间复杂度 O(n*h)。
func distanceKEnumeration(root *TreeNode, target *TreeNode, k int) []int {
	paths := map[*TreeNode][]*TreeNode{}
	var dfs func(node *TreeNode, path []*TreeNode)
	dfs = func(node *TreeNode, path []*TreeNode) {
		if node == nil {
			return
		}
		path = append(slices.Clone(path), node)
		paths[node] = path
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, nil)
	tp, ok := paths[target]
	if !ok {
		return []int{}
	}
	res := []int{}
	for node, p := range paths {
		common := 0
		for common < len(p) && common < len(tp) && p[common] == tp[common] {
			common++
		}
		if (len(p)-common)+(len(tp)-common) == k {
			res = append(res, node.Val)
		}
	}
	return res
}
