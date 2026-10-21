package trimbst

import "leetcode-go/template"

// 669. 修剪二叉搜索树
// https://leetcode.cn/problems/trim-a-binary-search-tree/
// 给一棵二叉搜索树和区间 [low, high]，删掉所有值不在区间内的节点，
// 保持剩余节点的相对结构（原来的祖孙关系不变），返回新的根。
//
// 核心思路：利用 BST 性质——
//   - 当前节点值 < low：它和它的整棵左子树都太小，答案只可能在右子树里；
//   - 当前节点值 > high：它和整棵右子树都太大，答案只可能在左子树里；
//   - 否则保留当前节点，分别修剪左右子树。

type TreeNode = template.TreeNode

// trimBST 解法一：递归 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// trimBSTIterative 解法二：迭代
// 1. 先让根落进区间：根太小就往右走，太大就往左走。
// 2. 修剪左侧：沿着左链往下，遇到左孩子 < low 就用它的右子树替换它（循环直到满足）。
// 3. 修剪右侧：同理沿右链，遇到右孩子 > high 就用它的左子树替换。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func trimBSTIterative(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	for node := root; node != nil; node = node.Left {
		for node.Left != nil && node.Left.Val < low {
			node.Left = node.Left.Right
		}
	}
	for node := root; node != nil; node = node.Right {
		for node.Right != nil && node.Right.Val > high {
			node.Right = node.Right.Left
		}
	}
	return root
}
