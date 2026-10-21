package delnodes

import (
	"slices"

	"leetcode-go/template"
)

// 1110. 删点成林
// https://leetcode.cn/problems/delete-nodes-and-return-forest/
// 给一棵值互不相同的二叉树和要删除的值列表 to_delete，删掉这些节点后，
// 原树会断成若干棵树（森林），返回每棵树的根（顺序任意）。
//
// 核心思路：一个节点会成为新树的根 ⇔ 它自己不被删，并且它的父节点不存在或被删了。
// 用后序 DFS：先处理孩子（孩子被删时把父指针置空），再处理自己；
// 自己被删时，把仍存活的孩子加入答案。最后根节点没被删的话也加入答案。

type TreeNode = template.TreeNode

// delNodes 解法一：后序 DFS + 哈希集合 ⭐ 最优方案
// 时间复杂度 O(n + m)（m 为待删除个数），空间复杂度 O(n + m)。
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	del := make(map[int]bool, len(toDelete))
	for _, v := range toDelete {
		del[v] = true
	}
	var res []*TreeNode
	var dfs func(node *TreeNode) *TreeNode // 返回处理后的节点（被删则返回 nil）
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		if !del[node.Val] {
			return node
		}
		if node.Left != nil {
			res = append(res, node.Left)
		}
		if node.Right != nil {
			res = append(res, node.Right)
		}
		return nil
	}
	if dfs(root) != nil {
		res = append(res, root)
	}
	return res
}

// delNodesEnumeration 解法二：逐个删除（对照用）
// 维护当前森林；每删除一个值，就在所有树里找到它（同时记录父节点），
// 断开父子连接，再把它的左右孩子作为新树加入森林。
// 时间复杂度 O(n*m)，空间复杂度 O(n)。
func delNodesEnumeration(root *TreeNode, toDelete []int) []*TreeNode {
	var forest []*TreeNode
	if root != nil {
		forest = append(forest, root)
	}
	for _, v := range toDelete {
		for i, tree := range forest {
			parent, node := findWithParent(nil, tree, v)
			if node == nil {
				continue
			}
			if parent == nil {
				forest = slices.Delete(forest, i, i+1)
			} else if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			if node.Left != nil {
				forest = append(forest, node.Left)
			}
			if node.Right != nil {
				forest = append(forest, node.Right)
			}
			break
		}
	}
	return forest
}

func findWithParent(parent, node *TreeNode, v int) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}
	if node.Val == v {
		return parent, node
	}
	if p, n := findWithParent(node, node.Left, v); n != nil {
		return p, n
	}
	return findWithParent(node, node.Right, v)
}
