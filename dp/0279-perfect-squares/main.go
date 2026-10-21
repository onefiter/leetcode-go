package perfectsquares

import "math"

// 279. 完全平方数
// https://leetcode.cn/problems/perfect-squares/
// 给一个正整数 n，返回和为 n 的完全平方数的最少数量（如 1、4、9、16...）。
//
// 核心思路：完全背包求最少物品数。物品是 1², 2², ..., k²（k² <= n），每个可以无限次使用，容量为 n。
// dp[j] 表示凑出 j 所需的最少平方数个数，dp[0] = 0，
// 转移：dp[j] = min(dp[j], dp[j-i*i] + 1)。
// 完全背包内层容量「正序」遍历，使同一物品在本轮内可以被重复选择。

// numSquares 解法一：完全背包 + 一维滚动数组 ⭐ 最优方案
// 时间复杂度 O(n * sqrt(n))，空间复杂度 O(n)。
func numSquares(n int) int {
	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt32
	}
	for i := 1; i*i <= n; i++ {
		sq := i * i
		for j := sq; j <= n; j++ { // 正序：物品可重复使用
			dp[j] = min(dp[j], dp[j-sq]+1)
		}
	}
	return dp[n]
}

// numSquaresMath 解法二：四平方和定理（对照用）
// 拉格朗日四平方和定理：任何正整数都能表示为至多 4 个平方数之和；
// 勒让德三平方和定理：n 需要恰好 4 个，当且仅当 n = 4^a * (8b+7)。
// 于是依次判断：是否完全平方数 -> 1；是否 4^a(8b+7) -> 4；能否拆成两个平方数 -> 2；否则 3。
// 时间复杂度 O(sqrt(n))，空间复杂度 O(1)。
func numSquaresMath(n int) int {
	isSquare := func(x int) bool {
		r := int(math.Sqrt(float64(x)))
		for r*r > x {
			r--
		}
		for (r+1)*(r+1) <= x {
			r++
		}
		return r*r == x
	}
	if isSquare(n) {
		return 1
	}
	m := n
	for m%4 == 0 {
		m /= 4
	}
	if m%8 == 7 {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}
