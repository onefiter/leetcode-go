package searchrotated

// 33. 搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/
// 一个「升序且元素互不相同」的数组，在某个未知下标处被旋转，
// 例如 [0,1,2,4,5,6,7] 旋转后变成 [4,5,6,7,0,1,2]。
// 给定旋转后的数组 nums 和目标值 target，返回其下标，不存在返回 -1。
// 要求时间复杂度 O(log n)。
//
// 核心思路：旋转后的数组从中点 mid 一分为二，
// 左半 [left, mid] 和右半 [mid, right] 至少有一半是「完全有序」的。
// 每一步先判断哪一半有序，再看 target 是否落在那半有序区间内：
//   - 在有序半区间内 -> 往那半收缩；
//   - 不在 -> 往另一半收缩。
// 这样仍是标准二分，复杂度 O(log n)。

// search 解法一：一次二分判断有序半区 ⭐ 最优方案
// 每次用 nums[mid] 命中则直接返回；否则判断 [left, mid] 是否有序：
//   - 若 nums[left] <= nums[mid] 说明左半有序，
//     target 落在 [nums[left], nums[mid]) 时收缩到左半，否则收缩到右半；
//   - 否则右半 [mid, right] 有序，
//     target 落在 (nums[mid], nums[right]] 时收缩到右半，否则收缩到左半。
//
// 时间复杂度 O(log n)：标准二分。
// 空间复杂度 O(1)。
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			// 左半 [left, mid] 有序。
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // target 在有序左半，收缩到左半
			} else {
				left = mid + 1 // 否则到右半
			}
		} else {
			// 右半 [mid, right] 有序。
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // target 在有序右半，收缩到右半
			} else {
				right = mid - 1 // 否则到左半
			}
		}
	}
	return -1
}

// searchEnumeration 解法二：顺序遍历（对照用，不满足 O(log n) 要求）
// 直接扫描每个元素，命中即返回下标。仅用于对照验证二分解法的正确性。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func searchEnumeration(nums []int, target int) int {
	for i, x := range nums {
		if x == target {
			return i
		}
	}
	return -1
}
