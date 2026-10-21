package validatebst

import (
	"math"

	"leetcode-go/template"
)

// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/
// 判断一棵二叉树是否是有效的二叉搜索树（BST）：
// 左子树所有节点 < 根 < 右子树所有节点，且左右子树本身也是 BST。
//
// 易错点：只比较「父子」不够，必须保证整棵左子树都小于根。
// 思路一：递归时给每个节点传入允许的取值区间 (lo, hi)，往左走收紧上界、往右走收紧下界。
// 思路二：BST 的中序遍历严格递增，边遍历边和前一个值比较即可。
// 节点值可能取到 math.MinInt32 / math.MaxInt32，边界用 int64 的 MinInt64/MaxInt64 表示「无穷」。

type TreeNode = template.TreeNode

// isValidBST 解法一：递归上下界 ⭐ 最优方案
// 每个节点必须落在开区间 (lo, hi) 内；递归左子树时上界变为当前值，右子树时下界变为当前值。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，h 为树高）。
func isValidBST(root *TreeNode) bool {
	var check func(node *TreeNode, lo, hi int64) bool
	check = func(node *TreeNode, lo, hi int64) bool {
		if node == nil {
			return true
		}
		v := int64(node.Val)
		if v <= lo || v >= hi {
			return false
		}
		return check(node.Left, lo, v) && check(node.Right, v, hi)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}

// isValidBSTInorder 解法二：中序遍历（迭代）判断严格递增
// 用栈模拟中序遍历，记录上一个访问的值 prev，一旦当前值 <= prev 就不是 BST。
// 用 hasPrev 标记是否已有前驱，避免和 MinInt32 等边界值混淆。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func isValidBSTInorder(root *TreeNode) bool {
	stack := []*TreeNode{}
	prev, hasPrev := 0, false
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if hasPrev && cur.Val <= prev {
			return false
		}
		prev, hasPrev = cur.Val, true
		cur = cur.Right
	}
	return true
}
