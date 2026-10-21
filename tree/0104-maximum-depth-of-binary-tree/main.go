package maxdepthtree

import "leetcode-go/template"

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 给定二叉树 root，返回其最大深度：从根到最远叶子节点的最长路径上的节点数。
//
// 核心思路：
//   - DFS：一棵树的深度 = max(左子树深度, 右子树深度) + 1，空树深度为 0；
//   - BFS：按层遍历，数一共有多少层。

type TreeNode = template.TreeNode

// maxDepthDFS 解法一：递归 DFS ⭐ 最优方案
// 后序计算左右子树深度，取较大者加 1。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，最坏 O(n)）。
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthDFS(root.Left), maxDepthDFS(root.Right)) + 1
}

// maxDepthBFS 解法二：BFS 层序计数
// 每处理完一整层，深度加 1。
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大层宽）。
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		depth++
		next := []*TreeNode{}
		for _, node := range cur {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
	}
	return depth
}
