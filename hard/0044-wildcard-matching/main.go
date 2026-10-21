package wildcard

// 44. 通配符匹配
// https://leetcode.cn/problems/wildcard-matching/
// 给字符串 s 和模式 p，实现支持 '?' 和 '*' 的通配符匹配：
// '?' 匹配任意单个字符，'*' 匹配任意字符序列（包括空序列），要求覆盖整个 s。
//
// 核心思路：
//   - DP：dp[i][j] 表示 s 前 i 个与 p 前 j 个能否匹配；
//     p[j-1] == '*' 时 dp[i][j] = dp[i][j-1]（* 匹配空）|| dp[i-1][j]（* 再吃一个字符）；
//     否则 dp[i][j] = dp[i-1][j-1] && (p[j-1] == '?' || s[i-1] == p[j-1])。
//   - 贪心：'*' 能吸收任意串，只需记住「最近一个 *」的位置，
//     失配时回溯到它，让它多吃一个字符再继续，前面的 * 不必再考虑。

// isMatch 解法一：贪心双指针 + 回溯到最近的 * ⭐ 最优方案
// 时间复杂度 最坏 O(m*n)、平均接近 O(m+n)，空间复杂度 O(1)。
func isMatch(s string, p string) bool {
	i, j := 0, 0
	star, mark := -1, 0 // star：最近 * 在 p 中的位置；mark：该 * 开始匹配时 s 的位置
	for i < len(s) {
		switch {
		case j < len(p) && (p[j] == '?' || p[j] == s[i]):
			i++
			j++
		case j < len(p) && p[j] == '*':
			star, mark = j, i // 先让 * 匹配空串
			j++
		case star != -1:
			mark++ // 失配：让最近的 * 多吃一个字符
			i, j = mark, star+1
		default:
			return false
		}
	}
	for j < len(p) && p[j] == '*' { // s 用完，p 剩下的只能全是 *
		j++
	}
	return j == len(p)
}

// isMatchDP 解法二：二维 DP（对照用）
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func isMatchDP(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n && p[j-1] == '*'; j++ {
		dp[0][j] = true // 前缀全是 * 才能匹配空串
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
