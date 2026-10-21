package rob3

import "leetcode-go/template"

// 337. 打家劫舍 III
// https://leetcode.cn/problems/house-robber-iii/
// 房屋构成一棵二叉树，根是入口；直接相连（父子）的两间房不能同时偷，求能偷到的最高金额。
//
// 核心思路：树形 DP。对每个节点返回一对状态 (不偷它的最大值, 偷它的最大值)，后序遍历自底向上合并：
// 偷当前节点，则两个孩子都不能偷：rob = node.Val + left.skip + right.skip；
// 不偷当前节点，孩子偷不偷随意，各取较大：skip = max(left.skip, left.rob) + max(right.skip, right.rob)。
// 答案为根节点两种状态的较大值。

type TreeNode = template.TreeNode

// rob 解法一：后序遍历 + 返回 (不偷, 偷) 二元状态 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈，h 为树高）。
func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (skip, take int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		ls, lt := dfs(node.Left)
		rs, rt := dfs(node.Right)
		take := node.Val + ls + rs
		skip := max(ls, lt) + max(rs, rt)
		return skip, take
	}
	return max(dfs(root))
}

// robMemo 解法二：记忆化递归（爷爷 + 孙子 vs 两个儿子，对照用）
// f(node) = max(node.Val + 四个孙子的 f, 两个儿子的 f)，用哈希表缓存每个节点的结果避免重复计算。
// 时间复杂度 O(n)，空间复杂度 O(n)（memo + 递归栈）。
func robMemo(root *TreeNode) int {
	memo := map[*TreeNode]int{}
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if v, ok := memo[node]; ok {
			return v
		}
		take := node.Val
		if node.Left != nil {
			take += f(node.Left.Left) + f(node.Left.Right)
		}
		if node.Right != nil {
			take += f(node.Right.Left) + f(node.Right.Right)
		}
		skip := f(node.Left) + f(node.Right)
		memo[node] = max(take, skip)
		return memo[node]
	}
	return f(root)
}
