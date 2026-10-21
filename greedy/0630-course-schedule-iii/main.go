package scheduleiii

import (
	"container/heap"
	"slices"
)

// 630. 课程表 III
// https://leetcode.cn/problems/course-schedule-iii/
// 有 n 门课，courses[i] = [duration, lastDay] 表示这门课需要连续上 duration 天，且必须在第 lastDay 天或之前结束。
// 从第 1 天开始，同一时间只能上一门课，求最多能修多少门课。
//
// 核心思路：按截止时间 lastDay 升序处理（截止早的先上，排在前面不会更差），
// 用最大堆维护已选课程的时长；每来一门课先选上，若总时长超过它的截止时间，
// 就踢掉堆中时长最长的那门。正确性：课程数只能 +1 或不变，在课程数相同的前提下
// 踢掉最长的课让总耗时最小，给后面的课留出最多的时间，所以每一步都保持「数量最多且耗时最少」。

// maxHeap 时长最大堆
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// scheduleCourse 解法一：按截止时间排序 + 最大堆反悔贪心 ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func scheduleCourse(courses [][]int) int {
	slices.SortFunc(courses, func(a, b []int) int { return a[1] - b[1] })
	h := &maxHeap{}
	total := 0
	for _, c := range courses {
		heap.Push(h, c[0])
		total += c[0]
		if total > c[1] {
			total -= heap.Pop(h).(int) // 反悔：踢掉最耗时的课
		}
	}
	return h.Len()
}

// scheduleCourseDP 解法二：按截止时间排序 + 0/1 背包 DP（对照用）
// dp[t] 表示总耗时恰好为 t 时最多能修的课程数，按截止时间顺序逐门考虑「选 / 不选」，
// 选第 i 门课要求 t <= lastDay。
// 时间复杂度 O(n * D)，空间复杂度 O(D)，D 为最大截止时间。
func scheduleCourseDP(courses [][]int) int {
	slices.SortFunc(courses, func(a, b []int) int { return a[1] - b[1] })
	maxDay := 0
	for _, c := range courses {
		maxDay = max(maxDay, c[1])
	}
	dp := make([]int, maxDay+1)
	for i := 1; i <= maxDay; i++ {
		dp[i] = -1 // -1 表示该耗时不可达
	}
	for _, c := range courses {
		d, last := c[0], c[1]
		for t := last; t >= d; t-- { // 倒序保证每门课只用一次
			if dp[t-d] >= 0 {
				dp[t] = max(dp[t], dp[t-d]+1)
			}
		}
	}
	return slices.Max(dp)
}
