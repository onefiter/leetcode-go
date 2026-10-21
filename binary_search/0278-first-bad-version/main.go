package firstbadversion

// 278. 第一个错误的版本
// https://leetcode.cn/problems/first-bad-version/
// 版本号 1..n，某个版本出错后，之后的版本全都是错的。
// 通过 isBadVersion(version) 接口找出第一个错误版本，尽量少调用接口。
//
// 一句话理解题意：isBadVersion 在 1..n 上是 false...false true...true，找第一个 true。
//
// 核心思路：单调的布尔序列，标准的「找第一个满足条件的位置」二分。

// isBadVersion 由 LeetCode 平台提供，这里用包级变量模拟，测试中赋值。
var isBadVersion func(version int) bool

// firstBadVersion 解法一：二分查找 ⭐ 最优方案
// 区间 [lo, hi] 内一定包含答案：mid 是坏版本则答案在 mid 或更左（hi = mid），
// 否则答案在右侧（lo = mid+1）。用 lo+(hi-lo)/2 防止溢出。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// firstBadVersionEnumeration 解法二：从 1 开始逐个检查（对照用）
// 返回第一个坏版本；题目保证存在坏版本，兜底返回 n。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func firstBadVersionEnumeration(n int) int {
	for v := 1; v <= n; v++ {
		if isBadVersion(v) {
			return v
		}
	}
	return n
}
