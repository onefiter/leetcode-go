package invertbinarytree

import "leetcode-go/template"

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/
// 给一棵二叉树的根 root，把每个节点的左右子树互换，返回翻转后的根。
//
// 核心思路：翻转整棵树 = 交换根的左右孩子 + 分别翻转左右子树。
// 每个节点只需要被访问一次并交换一次，用 DFS 或 BFS 遍历都可以。

type TreeNode = template.TreeNode

// invertTree 解法一：递归 DFS ⭐ 最优方案
// 先递归翻转左右子树，再把翻转后的结果交换挂回去。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，h 为树高）。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// invertTreeBFS 解法二：迭代 BFS
// 逐层取出节点，交换它的左右孩子，再把非空孩子入队。
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为树的最大宽度）。
func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
