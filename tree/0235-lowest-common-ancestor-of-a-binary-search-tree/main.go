package lcabst

import "leetcode-go/template"

// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 给定一棵 BST 和其中两个节点 p、q，找出它们的最近公共祖先（节点也可以是自己的祖先）。
//
// 核心思路：利用 BST 的有序性，从根往下走：
//   - p、q 都比当前节点小：LCA 一定在左子树；
//   - p、q 都比当前节点大：LCA 一定在右子树；
//   - 否则（一大一小，或其中一个就是当前节点）：当前节点就是分叉点，即 LCA。

type TreeNode = template.TreeNode

// lowestCommonAncestor 解法一：利用 BST 性质迭代 ⭐ 最优方案
// 只走一条从根向下的路径，不需要栈。
// 时间复杂度 O(h)，空间复杂度 O(1)。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}

// lowestCommonAncestorRecursive 解法二：利用 BST 性质递归
// 与解法一判断逻辑相同，用递归表达「往哪边走」。
// 时间复杂度 O(h)，空间复杂度 O(h)（递归栈）。
func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}
	return root
}
