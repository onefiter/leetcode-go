package findpeak

// 162. 寻找峰值
// https://leetcode.cn/problems/find-peak-element/
// 峰值元素是指严格大于左右相邻值的元素。nums[-1] = nums[n] = -∞，
// 且相邻元素不相等。返回任意一个峰值的下标，要求 O(log n)。
// 例如 [1,2,3,1] -> 2；[1,2,1,3,5,6,4] -> 1 或 5。
//
// 一句话理解题意：把数组想成山路，两端都是悬崖（-∞），找任意一个山顶。
//
// 核心思路：看 mid 和 mid+1 谁高。如果 nums[mid] < nums[mid+1]，往右是上坡，
// 右边一直走下去最终一定会遇到下坡或右端悬崖，所以右侧必有峰值；否则左侧（含 mid）必有峰值。

// findPeakElement 解法一：二分查找 ⭐ 最优方案
// 区间 [lo, hi] 内保证有峰值；比较 nums[mid] 与 nums[mid+1]，
// 上坡则 lo = mid+1，下坡则 hi = mid，直到 lo == hi。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func findPeakElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// findPeakElementEnumeration 解法二：线性扫描（对照用）
// 从左往右找第一个比下一个元素大的位置；都没有说明数组单调递增，最后一个就是峰值。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func findPeakElementEnumeration(nums []int) int {
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}
