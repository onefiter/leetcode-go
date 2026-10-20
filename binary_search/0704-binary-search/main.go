package binarysearch

import "slices"

// 704. 二分查找
// https://leetcode.cn/problems/binary-search/
// 给定升序且元素不重复的整数数组 nums 和目标值 target，存在则返回下标，否则返回 -1。
//
// 解题思路：
// 数组有序，每次取中点 mid 与 target 比较，就能排除一半的区间：
//   - nums[mid] == target：找到答案；
//   - nums[mid] <  target：答案只可能在 mid 右侧；
//   - nums[mid] >  target：答案只可能在 mid 左侧。
// 关键是保持“区间定义”（循环不变量）一致：选闭区间 [left, right] 还是左闭右开 [left, right)，
// 会决定循环条件写 < 还是 <=，以及收缩边界时写 mid 还是 mid±1。

// search 解法一：二分查找（闭区间 [left, right]）
// 区间内有元素的条件是 left <= right；排除 mid 后边界取 mid±1。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// searchHalfOpen 解法二：二分查找（左闭右开 [left, right)）
// right 本身不在区间内，因此循环条件是 left < right，右边界收缩时取 right = mid。
// mid 写成 left + (right-left)/2，避免 left+right 在其他语言中溢出。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func searchHalfOpen(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

// searchRecursive 解法三：递归二分（闭区间）
// 思路同解法一，用递归代替循环。
// 时间复杂度 O(log n)，空间复杂度 O(log n)（递归栈深度）。
func searchRecursive(nums []int, target int) int {
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return dfs(mid+1, right)
		}
		return dfs(left, mid-1)
	}
	return dfs(0, len(nums)-1)
}

// searchStdlib 解法四：标准库 slices.BinarySearch
// 返回 target 应插入的位置以及是否找到。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func searchStdlib(nums []int, target int) int {
	if i, found := slices.BinarySearch(nums, target); found {
		return i
	}
	return -1
}

// searchLinear 解法五：暴力遍历（对照用，没有利用有序性）
// 时间复杂度 O(n)，空间复杂度 O(1)。
func searchLinear(nums []int, target int) int {
	for i, x := range nums {
		if x == target {
			return i
		}
	}
	return -1
}
