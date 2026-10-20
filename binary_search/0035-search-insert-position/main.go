package searchinsertposition

import "slices"

// 35. 搜索插入位置
// 给定一个「升序排列」且「元素互不相同」的整数数组 nums 和一个目标值 target，
// 如果 target 在数组里，返回它的下标；如果不在，返回它按顺序插入后应该在的下标。
// 例如 nums=[1,3,5,6], target=5 -> 2；target=2 -> 1；target=7 -> 4。
//
// 一句话理解题意：返回「第一个 大于等于 target 的元素的下标」。
// 如果数组里没有大于等于 target 的元素（target 比所有元素都大），就返回 len(nums)。
//
// 为什么插入位置等于这个下标？因为数组是升序的，把 target 放在第一个 >= target
// 的元素的位置，target 左边的元素都比它小、右边的元素都 >= 它，顺序依然成立。

// searchInsert 解法一：顺序遍历（最直观，适合理解题意）
// 从前往后找第一个 nums[i] >= target 的位置：
//   - 一旦遇到 nums[i] >= target，i 就是答案（要么相等，要么是插入点）；
//   - 如果整个数组都比 target 小，循环结束后返回 len(nums)。
// 时间复杂度 O(n)：最坏情况要看完整个数组。
// 空间复杂度 O(1)：只用了几个变量。
func searchInsert(nums []int, target int) int {
	for i, x := range nums {
		if x >= target {
			return i
		}
	}
	return len(nums)
}

// searchInsertBinary 解法二：二分查找（闭区间 [left, right]）⭐ 推荐
// 数组有序，就可以像查字典一样每次排除一半，不用一个个看。
// 关键结论：二分结束后，left 正好落在「第一个 >= target」的位置，直接返回 left。
// 为什么是 left？循环结束条件是 left > right，此时：
//   - left 左边的元素都 < target（因为每次 nums[mid] < target 才让 left 右移）；
//   - right 右边的元素都 >= target。
// 所以 left 就是第一个 >= target 的下标，即插入位置。
// 时间复杂度 O(log n)：每轮区间缩小一半。
// 空间复杂度 O(1)。
func searchInsertBinary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// searchInsertHalfOpen 解法三：二分查找（左闭右开 [left, right)）
// 和上一版同一个思路，只是 right 表示「不包含」的边界，所以：
//   - 循环条件是 left < right（当 left == right 时区间为空）；
//   - 收缩右边界时直接 right = mid（因为 mid 本来就不在下一轮区间里）。
// 区间为空时 left 就是插入位置，无需再讨论 target 是否命中。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func searchInsertHalfOpen(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// searchInsertStdlib 解法四：标准库 slices.BinarySearch
// Go 标准库已经实现了二分查找：返回 (插入位置, 是否找到)。
// 插入位置正是题目要的结果，所以无论找没找到，直接返回它。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func searchInsertStdlib(nums []int, target int) int {
	i, _ := slices.BinarySearch(nums, target)
	return i
}
