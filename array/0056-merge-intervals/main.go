package mergeintervals

import "slices"

// merge 56. 合并区间
// https://leetcode.cn/problems/merge-intervals/
// 把有重叠的区间合并。先按左端点排序，重叠的区间在排序后必然相邻。

// merge 【解法一】排序 + 贪心 ⭐
// 按左端点升序排序后依次扫描：
// 若当前区间左端点 <= 结果中最后一个区间的右端点，说明重叠，更新右端点取较大者；
// 否则不重叠，直接追加为新区间。
// 时间 O(n log n)（排序主导）、空间 O(n)（结果数组，忽略排序栈）。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sorted := slices.Clone(intervals)
	slices.SortFunc(sorted, func(a, b []int) int {
		return a[0] - b[0]
	})
	res := [][]int{{sorted[0][0], sorted[0][1]}}
	for _, iv := range sorted[1:] {
		last := res[len(res)-1]
		if iv[0] <= last[1] {
			if iv[1] > last[1] {
				last[1] = iv[1]
			}
		} else {
			res = append(res, []int{iv[0], iv[1]})
		}
	}
	return res
}
