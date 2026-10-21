package completetree

import "leetcode-go/template"

// 958. 二叉树的完全性检验
// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
// 判断一棵二叉树是否是「完全二叉树」：除最后一层外每层都是满的，
// 且最后一层的节点都尽量靠左排列。
//
// 核心思路：按层序遍历（空孩子也入队），完全二叉树的层序序列一定是
// 「一串非空节点 + 一串空节点」。一旦遇到第一个空节点，后面就不能再出现非空节点。

type TreeNode = template.TreeNode

// isCompleteTree 解法一：BFS，遇到空节点后不能再有非空 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大宽度）。
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	seenNil := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			seenNil = true
			continue
		}
		if seenNil {
			return false
		}
		queue = append(queue, node.Left, node.Right)
	}
	return true
}

// isCompleteTreeIndex 解法二：节点编号法
// 给根编号 1，节点 i 的左右孩子编号 2i、2i+1（同堆的存储方式）。
// 完全二叉树 ⇔ 所有节点的最大编号恰好等于节点总数。
// 注意：树很深时编号可能溢出，本题节点数 ≤ 100，不会出现问题；
// 一旦编号超过节点数就可以提前判定为否，也顺带避免了溢出。
// 时间复杂度 O(n)，空间复杂度 O(w)。
func isCompleteTreeIndex(root *TreeNode) bool {
	if root == nil {
		return true
	}
	type item struct {
		node *TreeNode
		idx  int
	}
	queue := []item{{root, 1}}
	count := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		if cur.idx > count { // 编号跳过了某些位置，说明中间有空缺
			return false
		}
		if cur.node.Left != nil {
			queue = append(queue, item{cur.node.Left, 2 * cur.idx})
		}
		if cur.node.Right != nil {
			queue = append(queue, item{cur.node.Right, 2*cur.idx + 1})
		}
	}
	return true
}
