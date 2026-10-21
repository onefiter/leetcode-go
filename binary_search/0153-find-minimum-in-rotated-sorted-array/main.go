package findminrotated

// 153. 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
// 一个「升序且元素互不相同」的数组，在某个未知下标处被旋转，
// 例如 [0,1,2,4,5,6,7] 旋转后可能变成 [4,5,6,7,0,1,2]。
// 求旋转后数组中的最小元素。要求时间复杂度 O(log n)。
//
// 核心思路：最小值是数组中唯一的「拐点」（它比左边的元素小）。
// 用 nums[mid] 与 nums[right] 比较判断拐点在哪半：
//   - nums[mid] > nums[right]：mid 落在前半段（较大的那段），最小值在 mid 右侧；
//   - nums[mid] < nums[right]：mid 落在后半段，最小值在 mid 或其左侧。
// 因为元素互不相同，不存在相等的歧义，始终 O(log n)。

// findMin 解法一：二分比较 nums[mid] 与 nums[right] ⭐ 最优方案
// 左闭右闭 [left, right]，循环条件 left < right：
//   - 若 nums[mid] > nums[right]，最小值必在 (mid, right]，令 left = mid+1；
//   - 否则最小值在 [left, mid]，令 right = mid。
//
// 循环结束时 left == right，即最小值下标。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1 // mid 在前段，最小值在右侧
		} else {
			right = mid // mid 在后段，最小值在 mid 或左侧
		}
	}
	return nums[left]
}

// findMinEnumeration 解法二：顺序遍历取最小值（对照用，不满足 O(log n) 要求）
// 直接扫描所有元素取最小值。仅用于对照验证二分解法的正确性。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func findMinEnumeration(nums []int) int {
	m := nums[0]
	for _, x := range nums {
		if x < m {
			m = x
		}
	}
	return m
}
