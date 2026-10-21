package searchrotatedii

// 81. 搜索旋转排序数组 II
// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/
// 与 33 题类似，但数组「可能包含重复元素」，在某个未知下标处被旋转。
// 给定旋转后的数组 nums 和目标值 target，判断 target 是否存在，返回 bool。
//
// 核心思路：仍是「判断哪半有序再收缩」的二分，
// 但重复元素会破坏判断：当 nums[left] == nums[mid] == nums[right] 时，
// 无法确定哪半有序（例如 [1,0,1,1,1]）。此时只能把左右端各收缩一步
// （left++, right--）排除干扰，这一步使最坏复杂度退化为 O(n)。

// search 解法一：二分 + 去重端点 ⭐ 最优方案
// 命中直接返回 true。否则：
//   - 若 nums[left]、nums[mid]、nums[right] 三者相等，无法判断，left++、right--；
//   - 若 nums[left] <= nums[mid] 左半有序，target 在左半区间则收缩左半，否则右半；
//   - 否则右半有序，target 在右半区间则收缩右半，否则左半。
//
// 时间复杂度 O(log n) 平均；最坏 O(n)（大量重复，如全相同元素）。
// 空间复杂度 O(1)。
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			// 三端相等，无法判断有序半区，收缩两端。
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// 左半 [left, mid] 有序。
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半 [mid, right] 有序。
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// searchEnumeration 解法二：顺序遍历（对照用，不满足 O(log n) 要求）
// 直接扫描数组看是否存在 target。仅用于对照验证二分解法的正确性。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func searchEnumeration(nums []int, target int) bool {
	for _, x := range nums {
		if x == target {
			return true
		}
	}
	return false
}
