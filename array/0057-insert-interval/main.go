package insertinterval

// insert 57. 插入区间
// https://leetcode.cn/problems/insert-interval/
// 原区间列表已按左端点升序且互不重叠，插入新区间后仍保持有序不重叠。

// insert 【解法一】一次线性扫描 + 贪心合并 ⭐
// 分三段处理：
//  1. 所有在 newInterval 左侧且不重叠的区间（右端点 < newInterval 左端点）直接加入；
//  2. 所有与 newInterval 重叠的区间（左端点 <= newInterval 右端点）不断合并，
//     取左端点最小、右端点最大，形成一个新的合并区间；
//  3. 剩余在右侧不重叠的区间直接加入。
//
// 由于原数组有序，只需一次遍历。时间 O(n)、空间 O(n)（结果数组）。
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i, n := 0, len(intervals)
	start, end := newInterval[0], newInterval[1]

	// 1. 左侧不重叠
	for i < n && intervals[i][1] < start {
		res = append(res, intervals[i])
		i++
	}
	// 2. 重叠部分合并
	for i < n && intervals[i][0] <= end {
		if intervals[i][0] < start {
			start = intervals[i][0]
		}
		if intervals[i][1] > end {
			end = intervals[i][1]
		}
		i++
	}
	res = append(res, []int{start, end})
	// 3. 右侧不重叠
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}
