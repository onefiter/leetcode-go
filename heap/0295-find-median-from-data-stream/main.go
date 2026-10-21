package median

import (
	"container/heap"
	"slices"
)

// 295. 数据流的中位数
// https://leetcode.cn/problems/find-median-from-data-stream/
// 设计一个数据结构，支持：
//   - AddNum(num)：从数据流中添加一个整数；
//   - FindMedian()：返回目前所有元素的中位数（偶数个时取中间两数的平均值）。
//
// 核心思路：对顶堆。用大顶堆 lo 存较小的一半，小顶堆 hi 存较大的一半，
// 并保持 len(lo) == len(hi) 或 len(lo) == len(hi)+1。
// 这样中位数要么是 lo 堆顶，要么是两个堆顶的平均值。
// 添加时先放进 lo，再把 lo 的最大值挪到 hi（保证 lo 全部 <= hi），
// 若 hi 比 lo 多了，再把 hi 的最小值挪回 lo。

// intHeap 基于 container/heap 的小顶堆；存相反数即可当大顶堆用。
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// MedianFinder 解法一：大顶堆 + 小顶堆（对顶堆）⭐ 最优方案
// 时间复杂度 AddNum O(log n)、FindMedian O(1)，空间复杂度 O(n)。
type MedianFinder struct {
	lo *intHeap // 大顶堆（存相反数）：较小的一半
	hi *intHeap // 小顶堆：较大的一半
}

func Constructor() MedianFinder {
	return MedianFinder{lo: &intHeap{}, hi: &intHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lo, -num)
	heap.Push(this.hi, -heap.Pop(this.lo).(int)) // lo 的最大值移到 hi
	if this.hi.Len() > this.lo.Len() {
		heap.Push(this.lo, -heap.Pop(this.hi).(int)) // 保持 lo 不少于 hi
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(-(*this.lo)[0])
	}
	return float64(-(*this.lo)[0]+(*this.hi)[0]) / 2
}

// MedianFinderSorted 解法二：有序切片 + 二分插入（对照用）
// AddNum 二分找到插入位置后插入（需要搬移元素），FindMedian 直接取中间。
// 时间复杂度 AddNum O(n)（二分 O(log n) + 搬移 O(n)）、FindMedian O(1)，空间复杂度 O(n)。
type MedianFinderSorted struct {
	nums []int
}

func ConstructorSorted() MedianFinderSorted {
	return MedianFinderSorted{}
}

func (this *MedianFinderSorted) AddNum(num int) {
	i, _ := slices.BinarySearch(this.nums, num)
	this.nums = slices.Insert(this.nums, i, num)
}

func (this *MedianFinderSorted) FindMedian() float64 {
	n := len(this.nums)
	if n%2 == 1 {
		return float64(this.nums[n/2])
	}
	return float64(this.nums[n/2-1]+this.nums[n/2]) / 2
}
