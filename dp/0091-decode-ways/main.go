package decodeways

// 91. 解码方法
// https://leetcode.cn/problems/decode-ways/
// 字母 A-Z 映射为 "1"-"26"。给一个只含数字的非空字符串 s，返回解码方法的总数。
// 含前导零的分组（如 "06"）不合法。
//
// 核心思路：dp[i] 表示 s 前 i 个字符的解码方法数，dp[0] = 1（空串一种）。
// 转移：若 s[i-1] != '0'，可单独解码，dp[i] += dp[i-1]；
// 若 i >= 2 且 s[i-2:i] 在 "10"~"26" 之间，可两位一起解码，dp[i] += dp[i-2]。
// 遍历顺序：i 正序，只依赖前两项，可压缩为两个变量。

// numDecodings 解法一：滚动变量 DP ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func numDecodings(s string) int {
	prev2, prev1 := 0, 1 // dp[i-2], dp[i-1]，初始 dp[-1] 视为 0，dp[0] = 1
	for i := 1; i <= len(s); i++ {
		cur := 0
		if s[i-1] != '0' {
			cur += prev1
		}
		if i >= 2 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			cur += prev2
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}

// numDecodingsMemo 解法二：记忆化递归（对照用）
// dfs(i) 表示 s[i:] 的解码方法数：取一位（非 '0'）或取两位（10~26）。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func numDecodingsMemo(s string) int {
	n := len(s)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 1
		}
		if s[i] == '0' {
			return 0 // 以 0 开头无法解码
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := dfs(i + 1)
		if i+1 < n && (s[i]-'0')*10+(s[i+1]-'0') <= 26 {
			res += dfs(i + 2)
		}
		memo[i] = res
		return res
	}
	return dfs(0)
}
