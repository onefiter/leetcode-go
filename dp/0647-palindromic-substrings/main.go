package palindromic

// 647. 回文子串
// https://leetcode.cn/problems/palindromic-substrings/
// 给一个字符串 s，统计并返回其中回文子串的数目。
// 具有不同开始位置或结束位置的子串，即使字符相同也视为不同的子串。
//
// 核心思路（DP）：dp[i][j] 表示 s[i..j] 是否为回文串。
// 转移：s[i] == s[j] 且（j-i <= 1 或 dp[i+1][j-1]）时 dp[i][j] = true。
// 遍历顺序：dp[i][j] 依赖左下角 dp[i+1][j-1]，因此 i 倒序、j 从 i 正序。
// 每得到一个 true 计数加一。

// countSubstrings 解法一：中心扩展 ⭐ 最优方案
// 回文串关于中心对称，中心共 2n-1 个（单字符 n 个 + 两字符间隙 n-1 个），
// 从每个中心向两侧扩展，每扩展成功一次即多一个回文子串。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func countSubstrings(s string) int {
	n := len(s)
	count := 0
	for center := 0; center < 2*n-1; center++ {
		l, r := center/2, center/2+center%2
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			l--
			r++
		}
	}
	return count
}

// countSubstringsDP 解法二：区间 DP（i 倒序）
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func countSubstringsDP(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	count := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}
