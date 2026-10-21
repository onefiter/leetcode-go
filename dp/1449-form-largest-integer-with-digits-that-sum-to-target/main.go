package largestnumber

import "strings"

// 1449. 数位成本和为目标值的最大数字
// https://leetcode.cn/problems/form-largest-integer-with-digits-that-sum-to-target/
// 给一个长度为 9 的数组 cost，cost[i] 表示添加数字 i+1 的成本，以及整数 target。
// 返回成本总和恰好为 target 时能得到的最大整数（字符串形式，不含 0），无解返回 "0"。
//
// 核心思路：完全背包 + 贪心还原。物品是数字 1..9（可无限次使用），重量为 cost，容量为 target，必须恰好装满。
// 数字越长越大，所以先求最多能放几位：dp[j] 表示成本恰好为 j 时的最大位数，
// dp[0] = 0，其余初始化为 -inf（不可达），转移：dp[j] = max(dp[j], dp[j-cost[i]] + 1)。
// 完全背包内层容量「正序」遍历。
// 还原时从数字 9 往 1 贪心：只要 dp[j-cost[d]] == dp[j]-1 就放 d（高位放大数字）。

// largestNumber 解法一：完全背包求最大位数 + 贪心还原 ⭐ 最优方案
// 时间复杂度 O(9 * target)，空间复杂度 O(target)（不含答案字符串）。
func largestNumber(cost []int, target int) string {
	const negInf = -1 << 30
	dp := make([]int, target+1)
	for j := 1; j <= target; j++ {
		dp[j] = negInf
	}
	for _, c := range cost {
		for j := c; j <= target; j++ { // 正序：数字可重复使用
			dp[j] = max(dp[j], dp[j-c]+1)
		}
	}
	if dp[target] < 0 {
		return "0"
	}
	var sb strings.Builder
	for d, j := 9, target; d >= 1; d-- {
		c := cost[d-1]
		for j >= c && dp[j-c] == dp[j]-1 { // 高位优先放大数字
			sb.WriteByte(byte('0' + d))
			j -= c
		}
	}
	return sb.String()
}

// largestNumberString 解法二：完全背包直接维护最大字符串（对照用）
// dp[j] 直接存成本恰好为 j 时能得到的最大数字字符串（"" 表示 j=0，nil 状态用 ok 标记不可达）。
// 比较大小：先比长度，长度相同再比字典序。数字从小到大遍历、新数字拼到最前面，
// 保证同一字符串中数字非递增（大数字在高位）。
// 时间复杂度 O(9 * target * L)，空间复杂度 O(target * L)，L 为答案长度。
func largestNumberString(cost []int, target int) string {
	better := func(a, b string) bool { // a 是否严格大于 b
		if len(a) != len(b) {
			return len(a) > len(b)
		}
		return a > b
	}
	dp := make([]string, target+1)
	ok := make([]bool, target+1)
	ok[0] = true
	for d := 1; d <= 9; d++ {
		c := cost[d-1]
		for j := c; j <= target; j++ {
			if !ok[j-c] {
				continue
			}
			cand := string(byte('0'+d)) + dp[j-c]
			if !ok[j] || better(cand, dp[j]) {
				dp[j] = cand
				ok[j] = true
			}
		}
	}
	if !ok[target] {
		return "0"
	}
	return dp[target]
}
