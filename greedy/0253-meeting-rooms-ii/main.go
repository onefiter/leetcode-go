package meetingrooms2

import (
	"container/heap"
	"slices"
)

// 253. 会议室 II
// https://leetcode.cn/problems/meeting-rooms-ii/
// 给一个会议时间区间数组 intervals，intervals[i] = [start, end)，
// 返回容纳所有会议所需的最少会议室数量。
//
// 核心思路：所需会议室数 = 任一时刻同时进行的会议数的最大值。
// 把所有开始时间、结束时间分别排序，用双指针按时间顺序扫描：遇到开始就占用一间，
// 遇到结束就释放一间（结束时间 <= 开始时间时先释放，因为区间左闭右开可以复用）。
// 扫描过程中占用数的峰值就是答案——峰值时刻确实需要这么多间，而贪心复用空闲房间保证不会多用。

// minMeetingRooms 解法一：开始/结束时间分别排序 + 双指针扫描 ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, iv := range intervals {
		starts[i], ends[i] = iv[0], iv[1]
	}
	slices.Sort(starts)
	slices.Sort(ends)
	rooms, best, j := 0, 0, 0
	for _, s := range starts {
		for j < n && ends[j] <= s { // 先释放已结束的会议
			rooms--
			j++
		}
		rooms++
		best = max(best, rooms)
	}
	return best
}

// minMeetingRoomsHeap 解法二：按开始时间排序 + 最小堆维护各房间结束时间（对照用）
// 堆顶是最早空出来的房间，若它在当前会议开始前结束就复用（弹出），否则新开一间；
// 最终堆的大小就是会议室数量。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func minMeetingRoomsHeap(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})
	h := &minHeap{}
	for _, iv := range intervals {
		if h.Len() > 0 && (*h)[0] <= iv[0] {
			heap.Pop(h) // 复用最早空出的房间
		}
		heap.Push(h, iv[1])
	}
	return h.Len()
}

// minHeap 存放每间会议室当前会议的结束时间。
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
