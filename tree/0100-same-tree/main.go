package sametree

import "leetcode-go/template"

// 100. 相同的树
// https://leetcode.cn/problems/same-tree/
// 给两棵二叉树的根 p 和 q，判断它们是否「结构相同且对应节点值相同」。
//
// 核心思路：两棵树相同 ⇔ 根相同 && 左子树相同 && 右子树相同。
// 同步遍历两棵树，任一位置一个为空一个不为空、或值不等，就不相同。

type TreeNode = template.TreeNode

// isSameTree 解法一：递归 DFS ⭐ 最优方案
// 两个都空 → 相同；只有一个空 → 不同；值不同 → 不同；否则递归比较左右子树。
// 时间复杂度 O(min(m,n))，空间复杂度 O(min(h1,h2))（递归栈）。
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// isSameTreeBFS 解法二：迭代 BFS（成对入队）
// 队列里每次放一对「应当相同」的节点，出队后按解法一的规则比较，
// 再把 (左,左)、(右,右) 两对入队。
// 时间复杂度 O(min(m,n))，空间复杂度 O(min(m,n))。
func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	queue := [][2]*TreeNode{{p, q}}
	for len(queue) > 0 {
		a, b := queue[0][0], queue[0][1]
		queue = queue[1:]
		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		queue = append(queue, [2]*TreeNode{a.Left, b.Left}, [2]*TreeNode{a.Right, b.Right})
	}
	return true
}
