package duplicatesubtrees

import (
	"strconv"

	"leetcode-go/template"
)

// 652. 寻找重复的子树
// https://leetcode.cn/problems/find-duplicate-subtrees/
// 找出所有「重复出现」的子树（结构和值都相同），每类只返回其中任意一棵的根。
//
// 核心思路：给每棵子树一个能唯一标识它的「指纹」，用哈希表统计出现次数，
// 次数恰好变成 2 时记录一次答案（避免同一类被记录多次）。
//   - 序列化指纹：后序序列化成字符串 "(左)值(右)"，简单直观，但字符串可能很长。
//   - 三元组编号：把 (左子树编号, 值, 右子树编号) 映射成一个整数编号，
//     每个节点只做 O(1) 次哈希，整体 O(n)。

type TreeNode = template.TreeNode

// findDuplicateSubtrees 解法一：三元组编号 ⭐ 最优方案
// 空树编号为 0；非空子树的编号由三元组 (左编号, 值, 右编号) 决定，相同三元组复用同一个编号。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type triple struct{ l, v, r int }
	ids := map[triple]int{}
	count := map[int]int{}
	var res []*TreeNode
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		key := triple{dfs(node.Left), node.Val, dfs(node.Right)}
		id, ok := ids[key]
		if !ok {
			id = len(ids) + 1
			ids[key] = id
		}
		count[id]++
		if count[id] == 2 {
			res = append(res, node)
		}
		return id
	}
	dfs(root)
	return res
}

// findDuplicateSubtreesSerialize 解法二：后序序列化 + 哈希
// 每棵子树序列化为 "(" + 左 + ")" + 值 + "(" + 右 + ")"，括号保证结构和值的边界。
// 时间复杂度 O(n^2)（最坏情况下每层都要拼接长度 O(n) 的字符串），空间复杂度 O(n^2)。
func findDuplicateSubtreesSerialize(root *TreeNode) []*TreeNode {
	count := map[string]int{}
	var res []*TreeNode
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		s := "(" + dfs(node.Left) + ")" + strconv.Itoa(node.Val) + "(" + dfs(node.Right) + ")"
		count[s]++
		if count[s] == 2 {
			res = append(res, node)
		}
		return s
	}
	dfs(root)
	return res
}
