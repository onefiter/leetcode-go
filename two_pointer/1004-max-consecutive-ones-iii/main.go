package longestones

// 1004. 最大连续1的个数 III
// https://leetcode.cn/problems/max-consecutive-ones-iii/
// 给定一个二进制数组 nums 和一个整数 k，最多可以把 k 个 0 翻转成 1，
// 返回翻转后数组中连续 1 的最大个数。
// 例如 nums=[1,1,1,0,0,0,1,1,1,1,0], k=2 -> 6。
//
// 一句话理解题意：找一段最长的连续子数组，里面 0 的个数不超过 k。
//
// 核心思路：可变长滑动窗口。右端点扩张时统计窗口内 0 的个数，
// 超过 k 就收缩左端点，直到 0 的个数 <= k，每次更新最大窗口长度。

// longestOnes 解法一：滑动窗口 ⭐ 最优方案
// zeros 记录窗口内 0 的个数；zeros > k 时左端点右移，移出的是 0 就 zeros--。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func longestOnes(nums []int, k int) int {
	ans, left, zeros := 0, 0, 0
	for right, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// longestOnesEnumeration 解法二：枚举起点向右扩展（对照用）
// 固定起点 i，向右走并统计 0 的个数，超过 k 时停下，记录长度。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func longestOnesEnumeration(nums []int, k int) int {
	ans := 0
	for i := range nums {
		zeros := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				zeros++
			}
			if zeros > k {
				break
			}
			ans = max(ans, j-i+1)
		}
	}
	return ans
}
