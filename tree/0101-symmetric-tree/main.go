package symmetrictree

import "leetcode-go/template"

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/
// 给一棵二叉树的根，判断它是否轴对称（左右镜像）。
//
// 核心思路：树对称 ⇔ 左子树和右子树互为镜像。
// 两棵树 a、b 互为镜像 ⇔ 根值相等 && a.Left 与 b.Right 镜像 && a.Right 与 b.Left 镜像。

type TreeNode = template.TreeNode

// isSymmetric 解法一：递归 ⭐ 最优方案
// 用 mirror(a, b) 判断两棵子树是否镜像：外侧对外侧、内侧对内侧递归比较。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var mirror func(a, b *TreeNode) bool
	mirror = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		return mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
	}
	return mirror(root.Left, root.Right)
}

// isSymmetricQueue 解法二：迭代队列
// 队列里成对存放「应当互为镜像」的节点；出队比较后，
// 按 (a.Left, b.Right)、(a.Right, b.Left) 成对入队。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func isSymmetricQueue(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		a, b := queue[0], queue[1]
		queue = queue[2:]
		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		queue = append(queue, a.Left, b.Right, a.Right, b.Left)
	}
	return true
}
