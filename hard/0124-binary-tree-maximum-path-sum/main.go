package maxpathsum

import (
	"math"

	"leetcode-go/template"
)

// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 路径是树中一条节点序列，相邻节点有边相连，每个节点至多出现一次，不一定经过根。
// 路径和是路径上节点值之和，返回任意非空路径的最大路径和（节点值可为负）。
//
// 核心思路：后序遍历，对每个节点算「向下单边最大贡献」gain(node)：
//   gain(node) = node.Val + max(0, gain(left), gain(right))（负贡献直接舍弃）。
// 以 node 为「拐点」的路径和 = node.Val + max(0, gain(left)) + max(0, gain(right))，
// 用它更新全局答案；但返回给父节点的只能是单边，否则路径会分叉。

type TreeNode = template.TreeNode

// maxPathSum 解法一：递归后序 DFS + 全局最大值 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var gain func(node *TreeNode) int
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(gain(node.Left), 0)
		right := max(gain(node.Right), 0)
		ans = max(ans, node.Val+left+right) // 以 node 为拐点
		return node.Val + max(left, right)  // 只能向上交一条边
	}
	gain(root)
	return ans
}

// maxPathSumIterative 解法二：迭代后序遍历 + 哈希表存贡献（对照用）
// 先用栈得到「根-右-左」序列，倒过来就是后序，保证处理节点时孩子已算好。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func maxPathSumIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var order []*TreeNode
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, node)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	gain := map[*TreeNode]int{} // nil 键不存在时取到 0，正好是空子树的贡献
	ans := math.MinInt
	for i := len(order) - 1; i >= 0; i-- {
		node := order[i]
		left := max(gain[node.Left], 0)
		right := max(gain[node.Right], 0)
		ans = max(ans, node.Val+left+right)
		gain[node] = node.Val + max(left, right)
	}
	return ans
}
