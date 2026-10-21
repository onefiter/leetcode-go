package largestsumafterk

import (
	"container/heap"
	"slices"
)

// 1005. K 次取反后最大化的数组和
// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
// 给一个整数数组 nums 和整数 k，每次操作选一个下标 i 把 nums[i] 替换为 -nums[i]，
// 恰好执行 k 次（同一下标可重复选择），返回数组可能的最大和。
//
// 核心思路：排序后从小到大把负数翻成正数，每次翻转绝对值最大的负数收益最大；
// 负数翻完后若 k 还剩奇数次，就反复翻转当前绝对值最小的数（偶数次翻转可以抵消），
// 这样损失最小。正确性：每一步都选「对和贡献增量最大」的翻转，且偶数次翻转对同一元素无影响。

// largestSumAfterKNegations 解法一：排序 + 贪心 ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（排序栈）。
func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)
	for i := 0; i < len(nums) && k > 0 && nums[i] < 0; i++ {
		nums[i] = -nums[i] // 最小的负数（绝对值最大）优先翻正
		k--
	}
	sum, minAbs := 0, nums[0]
	for _, v := range nums {
		sum += v
		minAbs = min(minAbs, v) // 此时所有翻过的数都非负，最小值即绝对值最小
	}
	if k%2 == 1 {
		sum -= 2 * minAbs // 剩余奇数次全打在绝对值最小的数上
	}
	return sum
}

// minHeap 整数最小堆
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

// largestSumAfterKNegationsHeap 解法二：最小堆逐次模拟（对照用）
// 每次弹出当前最小值取反后放回，严格模拟 k 次操作。
// 时间复杂度 O(n + k log n)，空间复杂度 O(n)。
func largestSumAfterKNegationsHeap(nums []int, k int) int {
	h := minHeap(slices.Clone(nums))
	heap.Init(&h)
	for ; k > 0; k-- {
		h[0] = -h[0] // 直接修改堆顶后下沉
		heap.Fix(&h, 0)
	}
	sum := 0
	for _, v := range h {
		sum += v
	}
	return sum
}
