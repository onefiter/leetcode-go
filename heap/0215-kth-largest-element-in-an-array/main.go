package kthlargest

import (
	"container/heap"
	"math/rand/v2"
)

// 215. 数组中的第K个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array/
// 给整数数组 nums 和整数 k，返回数组排序后第 k 个最大的元素（不是第 k 个不同的元素），
// 要求时间复杂度 O(n)。
//
// 核心思路：第 k 大 = 升序排序后下标为 n-k 的元素。
//   - 快速选择：借用快排的 partition，每次只递归进入包含目标下标的一侧，期望 O(n)；
//     用随机 pivot + 三路划分（< = >）防止有序数组或大量重复元素退化成 O(n^2)。
//   - 小顶堆：维护大小为 k 的小顶堆，堆顶就是当前第 k 大。

// findKthLargest 解法一：随机 pivot + 三路划分快速选择 ⭐ 最优方案
// 注意：会原地修改 nums。
// 时间复杂度 期望 O(n)，空间复杂度 O(1)（迭代实现）。
func findKthLargest(nums []int, k int) int {
	target := len(nums) - k // 升序下标
	lo, hi := 0, len(nums)-1
	for lo < hi {
		pivot := nums[lo+rand.IntN(hi-lo+1)]
		// 三路划分后：[lo, lt) < pivot，[lt, gt] == pivot，(gt, hi] > pivot
		lt, i, gt := lo, lo, hi
		for i <= gt {
			switch {
			case nums[i] < pivot:
				nums[lt], nums[i] = nums[i], nums[lt]
				lt++
				i++
			case nums[i] > pivot:
				nums[i], nums[gt] = nums[gt], nums[i]
				gt--
			default:
				i++
			}
		}
		switch {
		case target < lt:
			hi = lt - 1
		case target > gt:
			lo = gt + 1
		default:
			return pivot // 目标落在等于区
		}
	}
	return nums[target]
}

// minHeap 基于 container/heap 的小顶堆。
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

// findKthLargestHeap 解法二：大小为 k 的小顶堆（对照用）
// 堆满 k 个后，新元素比堆顶大才替换堆顶，最终堆顶即第 k 大。
// 时间复杂度 O(n log k)，空间复杂度 O(k)。
func findKthLargestHeap(nums []int, k int) int {
	h := make(minHeap, 0, k)
	for _, x := range nums {
		if h.Len() < k {
			heap.Push(&h, x)
		} else if x > h[0] {
			h[0] = x
			heap.Fix(&h, 0)
		}
	}
	return h[0]
}
