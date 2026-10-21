package largestrectangle

// 84. 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/
// 给定 n 个非负整数表示柱状图中各柱子的高度，柱子宽度均为 1，
// 求能勾勒出来的矩形的最大面积。例如 [2,1,5,6,2,3] -> 10。
//
// 核心思路：以每根柱子 i 的高度作为矩形高度，矩形能向左右延伸到
// 「第一个比它矮的柱子」为止。所以只要求出每根柱子左右两侧第一个更矮的位置，
// 面积就是 heights[i] * (right - left - 1)。「找两侧第一个更小」正是单调栈的拿手好戏。

// largestRectangleArea 解法一：单调栈（一次遍历）⭐ 最优方案
// 栈内保存下标，对应高度单调递增。当前高度比栈顶矮时，栈顶柱子的右边界就是 i，
// 左边界就是它在栈中的下一个元素（若栈空则为 -1），于是可以结算栈顶的面积。
// 在末尾补一个高度 0 的哨兵，保证所有柱子最终都被结算。
// 时间复杂度 O(n)：每个下标进出栈各一次。空间复杂度 O(n)。
func largestRectangleArea(heights []int) int {
	n := len(heights)
	best := 0
	stack := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		h := 0 // i == n 时作为哨兵
		if i < n {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			best = max(best, heights[top]*(i-left-1))
		}
		stack = append(stack, i)
	}
	return best
}

// largestRectangleAreaEnumeration 解法二：枚举左边界暴力（对照用）
// 固定左边界 i，向右扩展右边界 j，同时维护区间最小高度 minH，
// 面积 = minH * (j - i + 1)。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func largestRectangleAreaEnumeration(heights []int) int {
	best := 0
	for i := range heights {
		minH := heights[i]
		for j := i; j < len(heights); j++ {
			minH = min(minH, heights[j])
			best = max(best, minH*(j-i+1))
		}
	}
	return best
}
