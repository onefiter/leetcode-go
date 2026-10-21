package mergetrees

import "leetcode-go/template"

// 617. 合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/
// 把两棵二叉树「叠」在一起：同一位置两个节点都存在就把值相加，
// 只有一个存在就直接用那个节点。返回合并后的树。
//
// 核心思路：同步遍历两棵树。某一侧为空时，另一侧的整棵子树可以直接复用，不必再往下走。

type TreeNode = template.TreeNode

// mergeTrees 解法一：递归 DFS ⭐ 最优方案
// 有一侧为空就返回另一侧；否则把 root2 的值加到 root1 上，再递归合并左右子树。
// 直接复用 root1 的节点，不额外分配。
// 时间复杂度 O(min(m,n))（只访问两树都有节点的位置），空间复杂度 O(min(h1,h2))。
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// mergeTreesBFS 解法二：迭代 BFS（成对入队）
// 队列里放「两边都非空」的节点对，把值加到 root1 上；
// 对左右孩子：两边都有就成对入队，只有 root2 有就直接挂到 root1 上。
// 时间复杂度 O(min(m,n))，空间复杂度 O(min(m,n))。
func mergeTreesBFS(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := [][2]*TreeNode{{root1, root2}}
	for len(queue) > 0 {
		a, b := queue[0][0], queue[0][1]
		queue = queue[1:]
		a.Val += b.Val
		if a.Left != nil && b.Left != nil {
			queue = append(queue, [2]*TreeNode{a.Left, b.Left})
		} else if a.Left == nil {
			a.Left = b.Left
		}
		if a.Right != nil && b.Right != nil {
			queue = append(queue, [2]*TreeNode{a.Right, b.Right})
		} else if a.Right == nil {
			a.Right = b.Right
		}
	}
	return root1
}
