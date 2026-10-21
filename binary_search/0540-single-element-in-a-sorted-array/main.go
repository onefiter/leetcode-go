package singlenonduplicate

// 540. 有序数组中的单一元素
// https://leetcode.cn/problems/single-element-in-a-sorted-array/
// 有序数组中每个元素都恰好出现两次，只有一个数只出现一次，找出它。要求 O(log n)。
// 例如 [1,1,2,3,3,4,4,8,8] -> 2。
//
// 一句话理解题意：成对出现的数里混进了一个「单身」，找出它。
//
// 核心思路：单一元素左边，每一对都是 (偶数下标, 奇数下标)；
// 单一元素右边，配对整体错位一格，变成 (奇数下标, 偶数下标)。
// 所以看偶数下标 mid 是否和 mid+1 相等，就知道单一元素在左还是在右。

// singleNonDuplicate 解法一：二分查找 ⭐ 最优方案
// 只在偶数下标上二分：mid 取偶数，若 nums[mid] == nums[mid+1]，
// 说明 mid 之前都配对正常，答案在 mid+2 及之后；否则答案在 mid 或更左。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid%2 == 1 {
			mid-- // 对齐到偶数下标
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

// singleNonDuplicateEnumeration 解法二：全部异或（对照用）
// a^a = 0，a^0 = a，成对的数异或后抵消，剩下的就是单一元素。不利用有序性。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func singleNonDuplicateEnumeration(nums []int) int {
	ans := 0
	for _, x := range nums {
		ans ^= x
	}
	return ans
}
