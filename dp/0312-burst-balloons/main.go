package burstballoons

// 312. 戳气球
// https://leetcode.cn/problems/burst-balloons/
// 有 n 个气球，编号 0 到 n-1，每个气球上标有数字 nums[i]。戳破第 i 个气球可获得
// nums[i-1] * nums[i] * nums[i+1] 枚硬币（越界视为 1），之后左右相邻。求能获得的最大硬币数。
//
// 核心思路：区间 DP，逆向思考「最后戳哪个」。在两端补 1 得到 val[0..n+1]，
// dp[i][j] 表示戳破开区间 (i, j) 内所有气球能得到的最大硬币数。
// 枚举 (i, j) 内最后一个被戳的气球 k，此时它左右邻居恰为 i 和 j：
// dp[i][j] = max(dp[i][k] + val[i]*val[k]*val[j] + dp[k][j])，i < k < j。
// 遍历顺序：dp[i][j] 依赖 dp[i][k]（同行、更靠左）和 dp[k][j]（更靠下的行），
// 因此 i 倒序、j 正序（等价于按区间长度从小到大）。答案为 dp[0][n+1]。

// maxCoins 解法一：区间 DP（i 倒序） ⭐ 最优方案
// 时间复杂度 O(n^3)，空间复杂度 O(n^2)。
func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	copy(val[1:], nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+val[i]*val[k]*val[j]+dp[k][j])
			}
		}
	}
	return dp[0][n+1]
}

// maxCoinsMemo 解法二：记忆化递归（对照用）
// solve(i, j) 含义同 dp[i][j]，自顶向下枚举最后戳破的气球 k。
// 时间复杂度 O(n^3)，空间复杂度 O(n^2)。
func maxCoinsMemo(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	copy(val[1:], nums)
	memo := make([][]int, n+2)
	for i := range memo {
		memo[i] = make([]int, n+2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var solve func(i, j int) int
	solve = func(i, j int) int {
		if i+1 >= j {
			return 0 // 开区间内没有气球
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		best := 0
		for k := i + 1; k < j; k++ {
			best = max(best, solve(i, k)+val[i]*val[k]*val[j]+solve(k, j))
		}
		memo[i][j] = best
		return best
	}
	return solve(0, n+1)
}
