package containsduplicateiii

import (
	"sort"
)

// 220. 存在重复元素 III
// https://leetcode.cn/problems/contains-duplicate-iii/
// 给定整数数组 nums 和两个整数 indexDiff(k)、valueDiff(t)，判断是否存在
// 两个「不同」的下标 i、j，满足：
//   - abs(i - j) <= indexDiff
//   - abs(nums[i] - nums[j]) <= valueDiff
//
// 核心思路（桶排序思想）：把数值按宽度 (valueDiff+1) 分桶，桶号相同的两个数差值
// 必然 <= valueDiff，直接命中；否则只需再检查相邻桶。维护一个大小为 indexDiff
// 的滑动窗口（每个桶最多存一个数），即可同时满足下标距离约束。

// containsNearbyAlmostDuplicate 【解法1】桶 + 滑动窗口 ⭐ 最优
// 桶宽取 valueDiff+1，保证同桶内任意两数差 <= valueDiff。
// 时间 O(n)：每个元素入桶、查相邻桶都是常数操作。
// 空间 O(min(n, indexDiff))：窗口内至多 indexDiff+1 个桶。
func containsNearbyAlmostDuplicate(nums []int, indexDiff, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}
	buckets := make(map[int]int)
	width := valueDiff + 1
	// bucketID 对负数做向下取整映射，保证分桶连续。
	bucketID := func(x int) int {
		if x >= 0 {
			return x / width
		}
		return (x+1)/width - 1
	}
	for i, x := range nums {
		id := bucketID(x)
		// 同桶：差值必 <= valueDiff。
		if _, ok := buckets[id]; ok {
			return true
		}
		// 相邻左桶。
		if v, ok := buckets[id-1]; ok && abs(x-v) <= valueDiff {
			return true
		}
		// 相邻右桶。
		if v, ok := buckets[id+1]; ok && abs(x-v) <= valueDiff {
			return true
		}
		buckets[id] = x
		// 维护窗口大小，移除下标超出 indexDiff 的旧元素。
		if i >= indexDiff {
			delete(buckets, bucketID(nums[i-indexDiff]))
		}
	}
	return false
}

// containsNearbyAlmostDuplicateSortedWindow 【解法2】有序结构 + 滑动窗口
// 维护一个按值有序的窗口（这里用已排序切片 + 二分），对每个新元素查找窗口中
// 第一个 >= x-valueDiff 的值，若它 <= x+valueDiff 则命中。
// 时间 O(n*logk)：每次二分查找与插入/删除。
// 空间 O(min(n, indexDiff))。
func containsNearbyAlmostDuplicateSortedWindow(nums []int, indexDiff, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}
	window := make([]int, 0, indexDiff+1)
	for i, x := range nums {
		// 查找窗口中第一个 >= x-valueDiff 的位置。
		pos := sort.SearchInts(window, x-valueDiff)
		if pos < len(window) && window[pos] <= x+valueDiff {
			return true
		}
		// 把 x 按序插入窗口。
		ins := sort.SearchInts(window, x)
		window = append(window, 0)
		copy(window[ins+1:], window[ins:])
		window[ins] = x
		// 移除最旧元素，保持窗口大小 <= indexDiff。
		if i >= indexDiff {
			old := nums[i-indexDiff]
			del := sort.SearchInts(window, old)
			window = append(window[:del], window[del+1:]...)
		}
	}
	return false
}

// containsNearbyAlmostDuplicateEnumeration 【解法3】Enumeration 暴力（对照用）
// 对每个 i，检查其后 indexDiff 个元素的值差是否 <= valueDiff。
// 时间 O(n*k)：k 为 indexDiff。
// 空间 O(1)。
func containsNearbyAlmostDuplicateEnumeration(nums []int, indexDiff, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}
	for i := range nums {
		for j := i + 1; j <= i+indexDiff && j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= valueDiff {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
