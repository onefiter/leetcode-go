package productlessk

// 713. 乘积小于 K 的子数组
// https://leetcode.cn/problems/subarray-product-less-than-k/
// 给你一个正整数数组 nums 和整数 k，返回子数组内所有元素乘积严格小于 k 的连续子数组的数目。
// 例如 nums=[10,5,2,6], k=100 -> 8。
//
// 一句话理解题意：数一数有多少段连续子数组，乘起来 < k。
//
// 核心思路：元素都是正数，窗口越长乘积越大，满足单调性，可以用滑动窗口。
// 固定右端点 right，找到最小的 left 使 [left, right] 乘积 < k，
// 那么以 right 结尾的合法子数组有 right-left+1 个（左端点取 left..right 都行）。

// numSubarrayProductLessThanK 解法一：滑动窗口 ⭐ 最优方案
// 右端点每进来一个数就乘进 prod，若 prod >= k 就不断右移 left 并除掉 nums[left]。
// k <= 1 时任何正整数乘积都 >= 1 >= k，直接返回 0（也避免 left 越过 right）。
// 时间复杂度 O(n)：left、right 各最多移动 n 次。空间复杂度 O(1)。
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans, prod, left := 0, 1, 0
	for right, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

// numSubarrayProductLessThanKEnumeration 解法二：枚举左端点向右累乘（对照用）
// 固定左端点 i，向右逐个乘，一旦乘积 >= k 就停止（正数乘积只会变大）。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func numSubarrayProductLessThanKEnumeration(nums []int, k int) int {
	ans := 0
	for i := range nums {
		prod := 1
		for j := i; j < len(nums); j++ {
			prod *= nums[j]
			if prod >= k {
				break
			}
			ans++
		}
	}
	return ans
}
