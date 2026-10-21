package lcabinarytree

import "leetcode-go/template"

// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 给定一棵普通二叉树（不是 BST）和其中两个节点 p、q，找出它们的最近公共祖先。
//
// 核心思路（后序递归）：对每个节点问「你的子树里有没有 p 或 q」。
//   - 当前节点就是 p 或 q：直接返回它；
//   - 左右子树各找到一个：当前节点就是分叉点，即 LCA；
//   - 只有一侧找到：把那一侧的结果往上传。
// 另一种思路：先记录每个节点的父亲，把 p 的祖先放进集合，再从 q 往上找第一个在集合里的节点。

type TreeNode = template.TreeNode

// lowestCommonAncestor 解法一：递归后序遍历 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root // p、q 分居两侧
	}
	if left != nil {
		return left
	}
	return right
}

// lowestCommonAncestorParent 解法二：父指针哈希 + 祖先集合
// 先用栈 DFS 记录每个节点的父节点；再把 p 及其所有祖先放入集合；
// 最后从 q 往上走，第一个出现在集合里的节点就是 LCA。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func lowestCommonAncestorParent(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := map[*TreeNode]*TreeNode{root: nil}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, child := range []*TreeNode{node.Left, node.Right} {
			if child != nil {
				parent[child] = node
				stack = append(stack, child)
			}
		}
	}
	ancestors := map[*TreeNode]bool{}
	for cur := p; cur != nil; cur = parent[cur] {
		ancestors[cur] = true
	}
	for cur := q; cur != nil; cur = parent[cur] {
		if ancestors[cur] {
			return cur
		}
	}
	return nil
}
