package binarytreediameter

import "leetcode-go/template"

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/
// 求二叉树中任意两节点之间最长路径的「边数」，路径不一定经过根节点。
//
// 核心思路：任何一条路径都有一个「最高点」。以节点 n 为最高点的最长路径
// = 左子树深度 + 右子树深度。后序遍历求每个节点的深度，顺手更新全局最大值。

type TreeNode = template.TreeNode

// diameterOfBinaryTree 解法一：后序 DFS 求深度并更新直径 ⭐ 最优方案
// depth(n) 返回以 n 为根的子树深度（节点数），经过 n 的最长路径边数 = 左深 + 右深。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var depth func(*TreeNode) int
	depth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := depth(n.Left), depth(n.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	depth(root)
	return ans
}

// diameterOfBinaryTreeIterative 解法二：迭代后序遍历 + 哈希表记录深度
// 用栈模拟后序遍历，节点出栈时其左右子树深度均已算好，存入 map。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func diameterOfBinaryTreeIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	depth := map[*TreeNode]int{nil: 0}
	stack := []*TreeNode{}
	var prev *TreeNode
	for cur := root; cur != nil || len(stack) > 0; {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		if top.Right != nil && top.Right != prev {
			cur = top.Right // 右子树还没处理
			continue
		}
		stack = stack[:len(stack)-1]
		l, r := depth[top.Left], depth[top.Right]
		ans = max(ans, l+r)
		depth[top] = max(l, r) + 1
		prev = top
	}
	return ans
}

// diameterOfBinaryTreeEnumeration 解法三：枚举每个节点作为最高点（对照用）
// 对每个节点单独计算左右子树高度，高度会被重复计算。
// 时间复杂度 O(n^2)（链状树最坏），空间复杂度 O(h)。
func diameterOfBinaryTreeEnumeration(root *TreeNode) int {
	if root == nil {
		return 0
	}
	through := height(root.Left) + height(root.Right)
	return max(through, diameterOfBinaryTreeEnumeration(root.Left), diameterOfBinaryTreeEnumeration(root.Right))
}

func height(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return max(height(n.Left), height(n.Right)) + 1
}
