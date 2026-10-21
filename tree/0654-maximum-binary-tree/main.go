package maxbinarytree

import "leetcode-go/template"

// 654. 最大二叉树
// https://leetcode.cn/problems/maximum-binary-tree/
// 给一个不含重复元素的整数数组 nums，按规则建树：
// 根是数组最大值；最大值左边的子数组递归建左子树，右边的子数组递归建右子树。
//
// 核心思路：
//   - 直接按定义递归：每层找区间最大值，最坏（数组有序）退化为 O(n^2)。
//   - 观察到：一个数的父节点是「左边第一个比它大的数」和「右边第一个比它大的数」中较小的那个。
//     用一个单调递减栈从左到右扫一遍就能确定所有父子关系，O(n)。

type TreeNode = template.TreeNode

// constructMaximumBinaryTree 解法一：单调栈 ⭐ 最优方案
// 维护一个值单调递减的栈。处理新节点 cur 时：
//   - 把栈里比 cur 小的节点全部弹出，最后一个被弹出的是 cur 左边最近的较大值之前那段的最大值，挂为 cur.Left；
//   - 弹完后若栈非空，栈顶比 cur 大，cur 挂为栈顶的右孩子（会被后续更大的值覆盖成正确位置）。
//
// 最终栈底就是整个数组的最大值，即根。
// 时间复杂度 O(n)（每个元素进出栈一次），空间复杂度 O(n)。
func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := []*TreeNode{}
	for _, x := range nums {
		cur := &TreeNode{Val: x}
		for len(stack) > 0 && stack[len(stack)-1].Val < x {
			cur.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			stack[len(stack)-1].Right = cur
		}
		stack = append(stack, cur)
	}
	if len(stack) == 0 {
		return nil
	}
	return stack[0]
}

// constructMaximumBinaryTreeRecursive 解法二：按定义递归（对照用）
// 在 [l, r] 里线性找最大值下标 m，以它为根，递归建 [l, m-1] 和 [m+1, r]。
// 时间复杂度 平均 O(n log n)、最坏 O(n^2)（数组有序时每层只少一个元素），空间复杂度 O(n)。
func constructMaximumBinaryTreeRecursive(nums []int) *TreeNode {
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := l
		for i := l + 1; i <= r; i++ {
			if nums[i] > nums[m] {
				m = i
			}
		}
		return &TreeNode{Val: nums[m], Left: build(l, m-1), Right: build(m+1, r)}
	}
	return build(0, len(nums)-1)
}
