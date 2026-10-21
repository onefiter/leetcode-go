package pathsum

import "leetcode-go/template"

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/
// 判断二叉树中是否存在一条「根节点到叶子节点」的路径，路径上节点值之和等于 targetSum。
//
// 核心思路：每往下走一步，就从 targetSum 里减掉当前节点值；
// 走到叶子时剩余值恰好等于叶子值，说明找到了。注意必须到叶子才算数。

type TreeNode = template.TreeNode

// hasPathSum 解法一：递归 DFS ⭐ 最优方案
// 把问题转化为：左子树或右子树中是否存在和为 targetSum-root.Val 的路径。
// 时间复杂度 O(n)，空间复杂度 O(h)，h 为树高（递归栈）。
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	rest := targetSum - root.Val
	return hasPathSum(root.Left, rest) || hasPathSum(root.Right, rest)
}

// hasPathSumIterative 解法二：迭代 DFS（显式栈保存节点与剩余和）
// 栈里同时存放节点和「到达该节点前还需要凑的和」，弹出时判断叶子。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	type item struct {
		node *TreeNode
		rest int
	}
	stack := []item{{root, targetSum}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n, rest := cur.node, cur.rest-cur.node.Val
		if n.Left == nil && n.Right == nil && rest == 0 {
			return true
		}
		if n.Right != nil {
			stack = append(stack, item{n.Right, rest})
		}
		if n.Left != nil {
			stack = append(stack, item{n.Left, rest})
		}
	}
	return false
}

// hasPathSumBFS 解法三：BFS 层序遍历，队列中记录根到当前节点的路径和
// 时间复杂度 O(n)，空间复杂度 O(w)，w 为树的最大宽度。
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodes := []*TreeNode{root}
	sums := []int{root.Val}
	for len(nodes) > 0 {
		n, s := nodes[0], sums[0]
		nodes, sums = nodes[1:], sums[1:]
		if n.Left == nil && n.Right == nil && s == targetSum {
			return true
		}
		if n.Left != nil {
			nodes, sums = append(nodes, n.Left), append(sums, s+n.Left.Val)
		}
		if n.Right != nil {
			nodes, sums = append(nodes, n.Right), append(sums, s+n.Right.Val)
		}
	}
	return false
}
