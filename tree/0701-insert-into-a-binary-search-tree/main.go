package insertbst

import "leetcode-go/template"

// 701. 二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/
// 把值 val 插入 BST，返回插入后的根节点。保证 val 与树中所有值都不同，答案可能不唯一。
//
// 核心思路：新节点总能作为叶子插入。像查找一样从根往下走（小往左、大往右），
// 走到某个空位置时，把新节点挂在那里即可，不需要调整树的其他部分。

type TreeNode = template.TreeNode

// insertIntoBST 解法一：迭代找空位插入 ⭐ 最优方案
// 时间复杂度 O(h)，空间复杂度 O(1)。
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	cur := root
	for {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return root
			}
			cur = cur.Right
		}
	}
}

// insertIntoBSTRecursive 解法二：递归插入
// 递归到空子树时返回新节点，由上一层接到对应孩子指针上。
// 时间复杂度 O(h)，空间复杂度 O(h)（递归栈）。
func insertIntoBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBSTRecursive(root.Left, val)
	} else {
		root.Right = insertIntoBSTRecursive(root.Right, val)
	}
	return root
}
