package balancedstring

// 1221. 分割平衡字符串
// https://leetcode.cn/problems/split-a-string-in-balanced-strings/
// 平衡字符串中 'L' 和 'R' 的数量相同。给一个平衡字符串 s，把它分割成尽可能多的平衡子串，
// 返回可以得到的平衡字符串的最大数量。
//
// 核心思路：从左往右计数，遇 'R' 加一、遇 'L' 减一，计数归零时说明前缀是平衡的，立即切一刀。
// 正确性：若某个最优解在计数归零处没切，那么这段可以从归零处拆成两段（两部分都平衡），
// 数量只增不减；所以「一平衡就切」一定能得到最多的片段。

// balancedStringSplit 解法一：计数器贪心 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func balancedStringSplit(s string) int {
	res, balance := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			res++
		}
	}
	return res
}

// balancedStringSplitDP 解法二：区间划分 DP（对照用）
// dp[i] 表示前 i 个字符最多能分成多少个平衡子串（-1 表示无法划分），
// 枚举最后一段 s[j:i]，若它平衡则 dp[i] = max(dp[i], dp[j]+1)；用前缀和 O(1) 判断平衡。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func balancedStringSplitDP(s string) int {
	n := len(s)
	prefix := make([]int, n+1) // prefix[i]：前 i 个字符中 R 数 - L 数
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if s[i] == 'R' {
			prefix[i+1]++
		} else {
			prefix[i+1]--
		}
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = -1
		for j := 0; j < i; j++ {
			if dp[j] >= 0 && prefix[i] == prefix[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return max(dp[n], 0)
}
