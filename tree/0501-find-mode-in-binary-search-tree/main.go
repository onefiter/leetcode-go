package findbstmode

import "leetcode-go/template"

// 501. 二叉搜索树中的众数
// https://leetcode.cn/problems/find-mode-in-binary-search-tree/
// 给一棵含重复值的 BST（左 <= 根 <= 右），找出出现次数最多的值（可能有多个）。
//
// 核心思路：BST 的中序遍历是非递减序列，相同的值一定连在一起。
// 于是只需边遍历边统计「当前值连续出现了几次」，和历史最大次数比较即可，
// 不需要哈希表。

type TreeNode = template.TreeNode

// findMode 解法一：中序遍历 + 连续计数（Morris，O(1) 额外空间）⭐ 最优方案
// 用 Morris 中序遍历按升序访问节点，维护 prev（上一个值）、cnt（当前连续次数）、maxCnt：
// cnt > maxCnt 时清空结果重新记录；cnt == maxCnt 时追加。
// 时间复杂度 O(n)，空间复杂度 O(1)（不计返回结果）。
func findMode(root *TreeNode) []int {
	var res []int
	prev, cnt, maxCnt := 0, 0, 0
	visit := func(v int) {
		if cnt > 0 && v == prev {
			cnt++
		} else {
			prev, cnt = v, 1
		}
		if cnt > maxCnt {
			maxCnt = cnt
			res = []int{v}
		} else if cnt == maxCnt {
			res = append(res, v)
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			visit(cur.Val)
			cur = cur.Right
			continue
		}
		// 找左子树最右节点（前驱）。
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur // 建立线索，稍后回到 cur
			cur = cur.Left
		} else {
			pre.Right = nil // 左子树已访问完，拆除线索
			visit(cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// findModeRecursive 解法二：递归中序遍历 + 连续计数
// 与解法一的计数逻辑相同，只是用递归做中序遍历，更直观。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func findModeRecursive(root *TreeNode) []int {
	var res []int
	prev, cnt, maxCnt := 0, 0, 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if cnt > 0 && node.Val == prev {
			cnt++
		} else {
			prev, cnt = node.Val, 1
		}
		if cnt > maxCnt {
			maxCnt = cnt
			res = []int{node.Val}
		} else if cnt == maxCnt {
			res = append(res, node.Val)
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// findModeEnumeration 解法三：哈希表计数（对照用）
// 任意遍历统计每个值的出现次数，再挑出次数最大的值。没有利用 BST 性质。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func findModeEnumeration(root *TreeNode) []int {
	count := map[int]int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		count[node.Val]++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	maxCnt := 0
	for _, c := range count {
		maxCnt = max(maxCnt, c)
	}
	var res []int
	for v, c := range count {
		if c == maxCnt {
			res = append(res, v)
		}
	}
	return res
}
