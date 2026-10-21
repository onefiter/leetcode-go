package convertbst

import "leetcode-go/template"

// 538. 把二叉搜索树转换为累加树
// https://leetcode.cn/problems/convert-bst-to-greater-tree/
// 给一棵二叉搜索树（BST），把每个节点的值改成「原树中所有大于等于它的值之和」。
//
// 核心思路：BST 的中序遍历（左-根-右）是升序；反过来「右-根-左」就是降序。
// 按降序访问时维护一个累加和 sum：每到一个节点，sum += node.Val，再把 node.Val 改成 sum。

type TreeNode = template.TreeNode

// convertBST 解法一：反向中序遍历（递归）⭐ 最优方案
// 先递归右子树（更大的值），再处理当前节点，最后递归左子树。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

// convertBSTMorris 解法二：反向 Morris 遍历
// 和普通 Morris 中序左右对调：找当前节点「右子树的最左节点」作为前驱，
// 用它的 Left 指针临时连回当前节点，从而不用栈也能回到父节点；遍历完会把临时指针恢复。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func convertBSTMorris(root *TreeNode) *TreeNode {
	sum := 0
	cur := root
	for cur != nil {
		if cur.Right == nil {
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
			continue
		}
		pre := cur.Right
		for pre.Left != nil && pre.Left != cur {
			pre = pre.Left
		}
		if pre.Left == nil { // 第一次到达：建立临时线索，先去右子树
			pre.Left = cur
			cur = cur.Right
		} else { // 第二次到达：右子树已处理完，拆掉线索，处理当前节点后去左子树
			pre.Left = nil
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
		}
	}
	return root
}

// convertBSTEnumeration 解法三：先收集所有值，再逐个求和（对照用）
// 遍历一遍收集所有原值；对每个节点，暴力累加所有 >= 它的原值。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func convertBSTEnumeration(root *TreeNode) *TreeNode {
	var nodes []*TreeNode
	var collect func(node *TreeNode)
	collect = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		collect(node.Left)
		collect(node.Right)
	}
	collect(root)
	orig := make([]int, len(nodes))
	for i, nd := range nodes {
		orig[i] = nd.Val
	}
	for i, nd := range nodes {
		s := 0
		for _, v := range orig {
			if v >= orig[i] {
				s += v
			}
		}
		nd.Val = s
	}
	return root
}
