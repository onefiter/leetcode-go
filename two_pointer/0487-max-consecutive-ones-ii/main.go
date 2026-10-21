package maxonesii

// 487. 最大连续1的个数 II
// https://leetcode.cn/problems/max-consecutive-ones-ii/
// 给定二进制数组 nums，最多可以把一个 0 翻转成 1，返回最大连续 1 的个数。
// 例如 nums=[1,0,1,1,0] -> 4（翻转第一个 0 得到 [1,1,1,1,0]）。
//
// 核心思路：等价于「找最长的、至多包含一个 0 的子数组」。
// 滑动窗口记录窗口内 0 的个数，超过 1 就收缩左端；
// 也可以只记住上一个 0 的位置，遇到新 0 时左端直接跳到它后面。

// findMaxConsecutiveOnes 解法一：滑动窗口（记录上一个 0 的位置）⭐ 最优方案
// left 是窗口左端，lastZero 是窗口内唯一那个 0 的下标（-1 表示没有）。
// 遇到新的 0：窗口左端跳到 lastZero+1，再把 lastZero 更新为当前下标。
// 只需 O(1) 状态，也适用于数据流（进阶要求）。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func findMaxConsecutiveOnes(nums []int) int {
	best, left, lastZero := 0, 0, -1
	for right, x := range nums {
		if x == 0 {
			left = lastZero + 1
			lastZero = right
		}
		best = max(best, right-left+1)
	}
	return best
}

// findMaxConsecutiveOnesDP 解法二：动态规划
// dp0 = 以当前位置结尾、未翻转过的连续 1 长度；
// dp1 = 以当前位置结尾、已用掉一次翻转的连续 1 长度。
// 遇到 1：dp0++，dp1++；遇到 0：dp1 = dp0+1（翻转当前 0），dp0 = 0。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func findMaxConsecutiveOnesDP(nums []int) int {
	best, dp0, dp1 := 0, 0, 0
	for _, x := range nums {
		if x == 1 {
			dp0++
			dp1++
		} else {
			dp1 = dp0 + 1
			dp0 = 0
		}
		best = max(best, dp0, dp1)
	}
	return best
}

// findMaxConsecutiveOnesEnumeration 解法三：枚举起点暴力扩展（对照用）
// 对每个起点向右扩展，统计 0 的个数，超过 1 个就停止。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func findMaxConsecutiveOnesEnumeration(nums []int) int {
	best := 0
	for i := range nums {
		zeros := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				zeros++
			}
			if zeros > 1 {
				break
			}
			best = max(best, j-i+1)
		}
	}
	return best
}
