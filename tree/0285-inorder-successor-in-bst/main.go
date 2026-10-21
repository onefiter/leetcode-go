package inordersuccessorbst

import "leetcode-go/template"

// 285. 二叉搜索树中的中序后继
// https://leetcode.cn/problems/inorder-successor-in-bst/
// 给定 BST 的根和其中一个节点 p，返回 p 的中序后继（比 p.Val 大的最小节点）；没有则返回 nil。
//
// 核心思路：中序后继就是「大于 p.Val 的最小值」。从根往下走：
//   - 当前值 > p.Val：当前节点可能是答案，先记下，再往左找更小的候选；
//   - 当前值 <= p.Val：答案只可能在右边。
// 走到空为止，最后记下的候选就是答案。

type TreeNode = template.TreeNode

// inorderSuccessor 解法一：利用 BST 性质自顶向下查找 ⭐ 最优方案
// 时间复杂度 O(h)，空间复杂度 O(1)。
func inorderSuccessor(root, p *TreeNode) *TreeNode {
	var succ *TreeNode
	cur := root
	for cur != nil {
		if cur.Val > p.Val {
			succ = cur // 候选，继续往左找更小的
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return succ
}

// inorderSuccessorEnumeration 解法二：迭代中序遍历（对照用）
// 不利用 BST 性质，按中序遍历，访问到 p 之后的下一个节点就是后继。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func inorderSuccessorEnumeration(root, p *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == p {
			return cur
		}
		prev = cur
		cur = cur.Right
	}
	return nil
}
