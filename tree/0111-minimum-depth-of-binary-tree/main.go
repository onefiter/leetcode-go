package mindepthtree

import "leetcode-go/template"

// 111. 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 给定二叉树 root，返回其最小深度：从根到「最近叶子节点」的最短路径上的节点数。
// 叶子节点指没有左右孩子的节点。
//
// 核心思路：
//   - BFS 一层一层往下找，遇到的第一个叶子所在的层就是答案，可以提前结束；
//   - DFS 要注意边界：如果某节点只有一侧子树，不能取 min(左, 右)（空的一侧是 0，
//     但空子树不是叶子），必须走非空的那一侧。例如 [1,2] 的答案是 2 而不是 1。

type TreeNode = template.TreeNode

// minDepthBFS 解法一：BFS 层序，遇到第一个叶子即返回 ⭐ 最优方案
// 按层处理，某层中第一次出现叶子节点时，当前层数就是最小深度。
// 树很不平衡时（如一侧很浅一侧很深）能提前结束，不必遍历整棵树。
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大层宽）。
func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := []*TreeNode{root}
	for depth := 1; ; depth++ {
		next := []*TreeNode{}
		for _, node := range cur {
			if node.Left == nil && node.Right == nil {
				return depth // 第一个叶子，提前结束
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
	}
}

// minDepthDFS 解法二：递归 DFS
// 空树为 0；叶子为 1；只有一侧子树时取那一侧深度 + 1；两侧都有时取较小者 + 1。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func minDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepthDFS(root.Right) + 1 // 只有右子树（或是叶子）
	}
	if root.Right == nil {
		return minDepthDFS(root.Left) + 1 // 只有左子树
	}
	return min(minDepthDFS(root.Left), minDepthDFS(root.Right)) + 1
}
