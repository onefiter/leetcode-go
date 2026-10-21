package recoverbst

import "leetcode-go/template"

// 99. 恢复二叉搜索树
// https://leetcode.cn/problems/recover-binary-search-tree/
// 一棵 BST 中恰好有两个节点的值被错误地交换了，要求在不改变结构的前提下恢复它。
//
// 核心思路：BST 的中序序列应严格递增。交换两个值后，中序序列会出现逆序对：
//   - 相邻两个被交换：只有一处 prev > cur，交换这两个即可；
//   - 不相邻：出现两处逆序，第一个错误节点是第一处的 prev，第二个是第二处的 cur。
// 统一做法：第一次遇到逆序记 first=prev、second=cur；之后再遇到逆序只更新 second=cur。
// 最后交换 first 和 second 的值。

type TreeNode = template.TreeNode

// recoverTree 解法一：Morris 中序遍历 ⭐ 最优方案
// 利用叶子节点空闲的右指针临时指向中序后继，免去栈；遍历完会把临时指针还原。
// 访问顺序与普通中序一致，边访问边找逆序点。
// 时间复杂度 O(n)（每条边最多走两次），空间复杂度 O(1)。
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	visit := func(cur *TreeNode) {
		if prev != nil && prev.Val > cur.Val {
			if first == nil {
				first = prev
			}
			second = cur
		}
		prev = cur
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			visit(cur)
			cur = cur.Right
			continue
		}
		// 找左子树最右节点（中序前驱）。
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur // 建立线索，稍后回到 cur
			cur = cur.Left
		} else {
			pre.Right = nil // 左子树已访问完，拆除线索
			visit(cur)
			cur = cur.Right
		}
	}
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeRecursive 解法二：递归中序遍历
// 与解法一的找逆序逻辑相同，只是用递归完成中序遍历。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func recoverTreeRecursive(root *TreeNode) {
	var first, second, prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
