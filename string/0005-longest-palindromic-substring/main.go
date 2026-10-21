package longestpalindrome

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/
// 给定一个字符串 s，找到 s 中最长的回文子串。
// 例如 s="babad" -> "bab"（"aba" 也符合）；s="cbbd" -> "bb"。
//
// 一句话理解题意：在 s 的所有子串中，找出「正着读和反着读一样」且最长的那一段。
//
// 核心思路：回文串关于中心对称。枚举每一个可能的中心，向两侧扩展，
// 记录能扩展出的最长回文。中心既可能是单个字符（奇数长度回文），
// 也可能是两个相邻字符之间的缝隙（偶数长度回文）。

// longestPalindrome 【解法1】中心扩展 ⭐ 最优实用方案
// 思路：遍历每个下标 i，分别以 (i,i) 为中心（奇数长度）、以 (i,i+1) 为中心
// （偶数长度）向两侧扩展，取所有扩展结果中最长的子串。
// 时间复杂度 O(n^2)：n 个中心，每个中心最多扩展 O(n)。
// 空间复杂度 O(1)：只用常数额外变量（返回切片除外）。
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		// 退出循环时 (l, r) 已越过边界，真正的回文是 [l+1, r-1]。
		if cur := r - l - 1; cur > maxLen {
			maxLen = cur
			start = l + 1
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // 奇数长度中心
		expand(i, i+1) // 偶数长度中心
	}
	return s[start : start+maxLen]
}

// longestPalindromeBrute 【解法2】暴力枚举 Enumeration（对照用）
// 思路：枚举所有子串 s[i:j+1]，逐个判断是否回文，记录最长的。
// 时间复杂度 O(n^3)：O(n^2) 个子串，每次判断回文 O(n)。
// 空间复杂度 O(1)。
func longestPalindromeBrute(s string) string {
	best := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if j-i+1 > len(best) && isPalin(s, i, j) {
				best = s[i : j+1]
			}
		}
	}
	return best
}

// isPalin 判断 s[l:r+1] 是否为回文。
func isPalin(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// longestPalindromeDP 【解法3】动态规划
// 思路：dp[i][j] 表示 s[i:j+1] 是否回文。
// 转移：dp[i][j] = (s[i]==s[j]) && (j-i<2 || dp[i+1][j-1])。
// 按子串长度从小到大递推。
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func longestPalindromeDP(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	start, maxLen := 0, 1
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] != s[j] {
				continue
			}
			if length == 2 || dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLen {
					maxLen = length
					start = i
				}
			}
		}
	}
	return s[start : start+maxLen]
}
