package searchbst

import "leetcode-go/template"

// 700. 二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/
// 在 BST 中找到值等于 val 的节点，返回以它为根的子树；不存在返回 nil。
//
// 核心思路：BST 左小右大，比当前值小就往左，比当前值大就往右，相等就找到了。
// 每一步都排除一半子树，只走一条从根向下的路径。

type TreeNode = template.TreeNode

// searchBST 解法一：迭代查找 ⭐ 最优方案
// 时间复杂度 O(h)，空间复杂度 O(1)。
func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil && cur.Val != val {
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur
}

// searchBSTRecursive 解法二：递归查找
// 时间复杂度 O(h)，空间复杂度 O(h)（递归栈）。
func searchBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBSTRecursive(root.Left, val)
	}
	return searchBSTRecursive(root.Right, val)
}
