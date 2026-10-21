package sqrtx

// 69. x 的平方根
// https://leetcode.cn/problems/sqrtx/
// 给定非负整数 x，返回 x 的算术平方根的整数部分（向下取整），不能用内置开方函数。
// 例如 x=4 -> 2；x=8 -> 2（√8≈2.828）。
//
// 一句话理解题意：找最大的整数 r，使得 r*r <= x。
//
// 核心思路：r*r 随 r 单调递增，所以「r*r <= x」在 [0, x] 上是一段前缀为真，
// 二分找这段前缀的最后一个位置即可。

// mySqrt 解法一：二分查找 ⭐ 最优方案
// 在 [0, x] 上二分：mid*mid <= x 时 mid 是候选答案，记下并往右找更大的；否则往左。
// x 最大 2^31-1，mid*mid 最大约 2^60，在 64 位 int 范围内不会溢出。
// 时间复杂度 O(log x)，空间复杂度 O(1)。
func mySqrt(x int) int {
	lo, hi := 0, x
	ans := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid <= x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

// mySqrtNewton 解法二：牛顿迭代
// 求 f(r)=r^2-x 的零点，迭代公式 r = (r + x/r) / 2。
// 从 r=x 出发，序列单调下降并收敛到 ⌊√x⌋，当 r*r <= x 时停止。
// 时间复杂度 O(log x)（实际收敛非常快），空间复杂度 O(1)。
func mySqrtNewton(x int) int {
	if x < 2 {
		return x
	}
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// mySqrtEnumeration 解法三：从小到大枚举（对照用）
// 从 0 开始逐个尝试，直到 (r+1)^2 > x。
// 时间复杂度 O(√x)，空间复杂度 O(1)。
func mySqrtEnumeration(x int) int {
	r := 0
	for (r+1)*(r+1) <= x {
		r++
	}
	return r
}
