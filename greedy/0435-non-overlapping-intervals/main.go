package nonoverlapping

import "slices"

// 435. 无重叠区间
// https://leetcode.cn/problems/non-overlapping-intervals/
// 给一个区间集合 intervals，intervals[i] = [start, end]，
// 返回需要移除的最少区间数，使剩余区间互不重叠（端点相接 [1,2],[2,3] 不算重叠）。
//
// 核心思路：等价于「最多能保留多少个互不重叠的区间」（经典活动选择问题）。
// 按右端点升序排序，每次选结束最早且与已选区间不冲突的区间。
// 正确性（交换论证）：任一最优解的第一个区间都可以换成结束最早的那个，
// 不会与后续区间产生新冲突，且为后面留出的空间最多。答案 = n - 最多保留数。

// cmpInt 用比较而非减法，避免极端值相减溢出。
func cmpInt(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// eraseOverlapIntervals 解法一：按右端点排序 + 贪心选最早结束 ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（排序栈）。
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	slices.SortFunc(intervals, func(a, b []int) int { return cmpInt(a[1], b[1]) })
	keep, end := 1, intervals[0][1]
	for _, iv := range intervals[1:] {
		if iv[0] >= end { // 不重叠，保留并更新边界
			keep++
			end = iv[1]
		}
	}
	return len(intervals) - keep
}

// eraseOverlapIntervalsDP 解法二：按左端点排序 + 动态规划（对照用）
// dp[i] 表示以第 i 个区间结尾时最多能保留的区间数，
// dp[i] = max(dp[j] + 1)，其中 j < i 且 intervals[j][1] <= intervals[i][0]。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func eraseOverlapIntervalsDP(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	slices.SortFunc(intervals, func(a, b []int) int { return cmpInt(a[0], b[0]) })
	dp := make([]int, n)
	best := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		best = max(best, dp[i])
	}
	return n - best
}
