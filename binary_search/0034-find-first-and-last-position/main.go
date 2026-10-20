package findfirstlast

import "sort"

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 给定一个「非递减」排列的整数数组 nums 和目标值 target，
// 找出 target 的开始位置和结束位置。如果不存在，返回 [-1, -1]。
// 要求时间复杂度 O(log n)。
// 例如 nums=[5,7,7,8,8,10], target=8 -> [3,4]；target=6 -> [-1,-1]。
//
// 一句话理解题意：target 在数组里是「连续的一段」（因为有序），
// 我们要找这一段的左端点和右端点下标。
//
// 核心思路：拆成两个「找边界」的二分：
//   - 左边界 = 第一个 >= target 的下标（如果它正好等于 target，就是开始位置）；
//   - 右边界 = 第一个 >  target 的下标再减 1（即最后一个 == target 的位置）。
// 这两个都能用「找第一个满足条件的下标」这一个通用二分解决。

// searchRange 解法一：两次二分找左右边界 ⭐ 最优方案
// lowerBound(target)   返回第一个 >= target 的下标；
// lowerBound(target+1) 返回第一个 >= target+1 的下标，也就是第一个 > target 的下标，
// 它减 1 就是最后一个 == target 的位置。
// 找到左边界后，先判断它是否越界、是否真的等于 target，不等说明 target 不存在。
// 时间复杂度 O(log n)：两次二分。
// 空间复杂度 O(1)。
func searchRange(nums []int, target int) []int {
	left := lowerBound(nums, target)
	// target 不存在：left 越界，或 left 处的值不是 target。
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	// target 存在，右边界 = 第一个 > target 的下标 - 1。
	right := lowerBound(nums, target+1) - 1
	return []int{left, right}
}

// lowerBound 返回 nums 中第一个 >= target 的下标；都比 target 小则返回 len(nums)。
// 左闭右开区间 [left, right) 写法：循环结束时 left 即答案。
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // mid 太小，答案在右侧
		} else {
			right = mid // mid 满足 >=，答案在 mid 或更左
		}
	}
	return left
}

// searchRangeTwoBinary 解法二：分别写「找左边界」和「找右边界」两个二分
// 和解法一等价，只是把右边界的二分单独写出来，逻辑更直白，便于理解。
// findLeft：找第一个 == target 的位置；findRight：找最后一个 == target 的位置。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func searchRangeTwoBinary(nums []int, target int) []int {
	l := findLeft(nums, target)
	if l == -1 {
		return []int{-1, -1}
	}
	return []int{l, findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			// 命中 target 时继续往左收缩，寻找更靠左的位置。
			if nums[mid] == target {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			// 命中 target 时继续往右收缩，寻找更靠右的位置。
			if nums[mid] == target {
				res = mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

// searchRangeStdlib 解法三：标准库 sort.SearchInts
// sort.SearchInts 返回第一个 >= target 的下标，等价于 lowerBound。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func searchRangeStdlib(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}

// searchRangeLinear 解法四：顺序遍历（对照用，不满足 O(log n) 要求）
// 利用数组非递减的性质：一旦遇到比 target 大的元素，后面不可能再有 target，
// 直接提前返回，减少不必要的遍历（最坏仍是 O(n)，如全部等于 target）。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func searchRangeLinear(nums []int, target int) []int {
	l, r := -1, -1
	for i, x := range nums {
		if x > target {
			return []int{l, r}
		}
		if x == target {
			if l == -1 {
				l = i
			}
			r = i
		}
	}
	return []int{l, r}
}
