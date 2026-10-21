package slidingmax

import "container/heap"

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/
// 给定整数数组 nums 和窗口大小 k，窗口从最左侧滑到最右侧，每次右移一位，
// 返回每个窗口内的最大值。
// 例如 nums=[1,3,-1,-3,5,3,6,7], k=3 -> [3,3,5,5,6,7]。
//
// 核心思路：维护一个「下标递增、值递减」的双端队列。
// 新元素进来时，把队尾所有比它小的元素踢掉（它们不可能再成为最大值）；
// 队头下标滑出窗口就弹出；此时队头就是当前窗口最大值。

// maxSlidingWindow 解法一：单调双端队列 ⭐ 最优方案
// q 存下标，对应的值从队头到队尾严格递减。
// 时间复杂度 O(n)：每个下标最多入队、出队各一次。
// 空间复杂度 O(k)：队列最多保存 k 个下标。
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	res := make([]int, 0, len(nums)-k+1)
	q := []int{}
	for i, x := range nums {
		// 队尾比 x 小（或相等）的元素永远不会再成为最大值。
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		// 队头已滑出窗口 [i-k+1, i]。
		if q[0] <= i-k {
			q = q[1:]
		}
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// maxHeap 按值降序的下标堆，用于解法二。
type maxHeap struct {
	idx  []int
	nums []int
}

func (h maxHeap) Len() int           { return len(h.idx) }
func (h maxHeap) Less(i, j int) bool { return h.nums[h.idx[i]] > h.nums[h.idx[j]] }
func (h maxHeap) Swap(i, j int)      { h.idx[i], h.idx[j] = h.idx[j], h.idx[i] }
func (h *maxHeap) Push(x any)        { h.idx = append(h.idx, x.(int)) }
func (h *maxHeap) Pop() any {
	v := h.idx[len(h.idx)-1]
	h.idx = h.idx[:len(h.idx)-1]
	return v
}

// maxSlidingWindowHeap 解法二：最大堆 + 延迟删除
// 堆里存下标，按值排序。每次取堆顶前，先把已滑出窗口的堆顶弹掉（延迟删除）。
// 时间复杂度 O(n log n)：最坏情况下堆里有 n 个元素。
// 空间复杂度 O(n)。
func maxSlidingWindowHeap(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	res := make([]int, 0, len(nums)-k+1)
	h := &maxHeap{nums: nums}
	for i := range nums {
		heap.Push(h, i)
		if i >= k-1 {
			for h.idx[0] <= i-k {
				heap.Pop(h)
			}
			res = append(res, nums[h.idx[0]])
		}
	}
	return res
}

// maxSlidingWindowEnumeration 解法三：逐窗口暴力求最大值（对照用）
// 时间复杂度 O(n * k)，空间复杂度 O(1)（不计答案）。
func maxSlidingWindowEnumeration(nums []int, k int) []int {
	res := []int{}
	if k <= 0 {
		return res
	}
	for i := 0; i+k <= len(nums); i++ {
		m := nums[i]
		for j := i + 1; j < i+k; j++ {
			m = max(m, nums[j])
		}
		res = append(res, m)
	}
	return res
}
