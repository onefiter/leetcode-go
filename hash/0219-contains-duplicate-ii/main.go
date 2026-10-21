package containsduplicateii

// 219. 存在重复元素 II
// https://leetcode.cn/problems/contains-duplicate-ii/
// 给定整数数组 nums 和整数 k，判断是否存在两个「不同」的下标 i、j，
// 满足 nums[i] == nums[j] 且 abs(i-j) <= k。
//
// 核心思路：用哈希表记录每个值「最近一次出现」的下标。遍历到 nums[i] 时，
// 如果该值之前出现过且下标差 <= k，即找到答案；否则更新为当前下标。
// 只保留最近下标是因为：对同一个值，越靠近 i 的历史下标越可能满足距离约束。

// containsNearbyDuplicate 【解法1】哈希表记录最近下标 ⭐ 最优
// 时间 O(n)：一次遍历。
// 空间 O(min(n, k))：等价地也可用「滑动窗口集合」把空间限制在 k。
func containsNearbyDuplicate(nums []int, k int) bool {
	last := make(map[int]int, len(nums))
	for i, x := range nums {
		if j, ok := last[x]; ok && i-j <= k {
			return true
		}
		last[x] = i
	}
	return false
}

// containsNearbyDuplicateWindow 【解法2】滑动窗口哈希集合
// 维护一个大小不超过 k 的窗口集合：加入 nums[i] 前，若集合里已存在该值则命中；
// 当窗口超过 k 时，移除最左端元素，保证集合里只有下标在 [i-k, i-1] 的值。
// 时间 O(n)，空间 O(min(n, k))。
func containsNearbyDuplicateWindow(nums []int, k int) bool {
	window := make(map[int]struct{})
	for i, x := range nums {
		if _, ok := window[x]; ok {
			return true
		}
		window[x] = struct{}{}
		if len(window) > k {
			delete(window, nums[i-k])
		}
	}
	return false
}

// containsNearbyDuplicateEnumeration 【解法3】Enumeration 暴力（对照用）
// 对每个 i，只检查 [i+1, i+k] 范围内是否出现相同值。
// 时间 O(n*k)：每个位置最多看 k 个后继。
// 空间 O(1)。
func containsNearbyDuplicateEnumeration(nums []int, k int) bool {
	for i := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
