package rob

// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/
// 一排房屋，每间存有非负金额 nums[i]，相邻两间不能同时偷，求能偷到的最高金额。
//
// 核心思路：dp[i] 表示前 i 间房能偷到的最大金额。
// 对第 i 间房只有两种选择：不偷，则为 dp[i-1]；偷，则前一间不能偷，为 dp[i-2] + nums[i-1]。
// 转移方程 dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])，dp[0] = 0，dp[1] = nums[0]。
// dp[i] 只依赖前两项，可以用两个变量滚动。

// rob 解法一：滚动变量 DP ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func rob(nums []int) int {
	prev, cur := 0, 0 // prev = dp[i-2]，cur = dp[i-1]
	for _, x := range nums {
		prev, cur = cur, max(cur, prev+x)
	}
	return cur
}

// robMemo 解法二：记忆化递归（自顶向下，对照用）
// f(i) 表示从第 i 间房开始往后能偷到的最大金额，f(i) = max(f(i+1), nums[i]+f(i+2))。
// 时间复杂度 O(n)，空间复杂度 O(n)（memo + 递归栈）。
func robMemo(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var f func(i int) int
	f = func(i int) int {
		if i >= n {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = max(f(i+1), nums[i]+f(i+2))
		return memo[i]
	}
	return f(0)
}
