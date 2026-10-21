package guessnumber

// 374. 猜数字大小
// https://leetcode.cn/problems/guess-number-higher-or-lower/
// 系统从 1..n 中选了一个数 pick，调用 guess(num) 得到提示：
// num > pick 返回 -1，num < pick 返回 1，相等返回 0。找出 pick。
//
// 一句话理解题意：经典「我心里想一个数，你来猜，我告诉你大了还是小了」。
//
// 核心思路：每次猜区间中点，根据提示砍掉一半区间。

// guess 由 LeetCode 平台提供，这里用包级变量模拟，测试中赋值。
var guess func(num int) int

// guessNumber 解法一：二分查找 ⭐ 最优方案
// 在 [lo, hi] 上猜 mid：返回 -1 说明猜大了（hi = mid-1），返回 1 说明猜小了（lo = mid+1）。
// 用 lo+(hi-lo)/2 防止溢出。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func guessNumber(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			hi = mid - 1
		default:
			lo = mid + 1
		}
	}
	return -1 // 题目保证有解，不会走到这里
}

// guessNumberEnumeration 解法二：从 1 开始逐个猜（对照用）
// 时间复杂度 O(n)，空间复杂度 O(1)。
func guessNumberEnumeration(n int) int {
	for v := 1; v <= n; v++ {
		if guess(v) == 0 {
			return v
		}
	}
	return -1
}
