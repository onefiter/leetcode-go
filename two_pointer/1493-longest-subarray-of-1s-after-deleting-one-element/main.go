package deleteone

// 1493. 删掉一个元素以后全为 1 的最长子数组
// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/
// 给你一个二进制数组 nums，你必须删掉一个元素，返回删掉后最长的且只包含 1 的非空子数组的长度。
// 如果不存在这样的子数组，返回 0。
// 例如 nums=[1,1,0,1] -> 3；nums=[1,1,1] -> 2（必须删一个，全 1 时答案是 n-1）。
//
// 一句话理解题意：找一段最多含 1 个 0 的最长连续子数组，删掉其中一个元素（有 0 删 0，没 0 删一个 1），
// 所以答案 = 窗口长度 - 1。
//
// 核心思路：可变长滑动窗口，窗口内 0 的个数不超过 1。

// longestSubarray 解法一：滑动窗口（窗口内至多一个 0）⭐ 最优方案
// zeros > 1 时收缩左端点；每个合法窗口删掉一个元素后长度为 right-left，取最大值。
// 全 1 时窗口为整个数组，答案自然是 n-1。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func longestSubarray(nums []int) int {
	ans, left, zeros := 0, 0, 0
	for right, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		ans = max(ans, right-left) // 窗口长度 right-left+1，再删掉一个
	}
	return ans
}

// longestSubarrayDP 解法二：递推（记录前一段和当前段连续 1 的长度）
// cur 为以当前位置结尾的连续 1 个数，prev 为上一个 0 之前紧邻的连续 1 个数。
// 遇到 0 时 prev = cur, cur = 0；删掉这个 0 能把 prev 与之后的 cur 拼起来，答案取 prev+cur 最大值。
// 若从未出现 0，必须删一个 1，答案为 n-1。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func longestSubarrayDP(nums []int) int {
	ans, prev, cur, hasZero := 0, 0, 0, false
	for _, x := range nums {
		if x == 1 {
			cur++
		} else {
			hasZero = true
			prev, cur = cur, 0
		}
		ans = max(ans, prev+cur)
	}
	if !hasZero && len(nums) > 0 {
		return len(nums) - 1
	}
	return ans
}

// longestSubarrayEnumeration 解法三：枚举删除的位置（对照用）
// 对每个被删除的下标 d，扫描剩余元素求最长连续 1 的长度，取最大值。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func longestSubarrayEnumeration(nums []int) int {
	ans := 0
	for d := range nums {
		run := 0
		for i, x := range nums {
			if i == d {
				continue // 被删除的元素直接跳过，两侧自然拼接
			}
			if x == 1 {
				run++
				ans = max(ans, run)
			} else {
				run = 0
			}
		}
	}
	return ans
}
