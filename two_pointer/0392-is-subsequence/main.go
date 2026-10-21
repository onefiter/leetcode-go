package issubsequence

// 392. 判断子序列
// https://leetcode.cn/problems/is-subsequence/
// 给定字符串 s 和 t，判断 s 是否为 t 的子序列（删去 t 中部分字符、不改变相对顺序能得到 s）。
// 例如 s="abc", t="ahbgdc" -> true；s="axc", t="ahbgdc" -> false。
//
// 核心思路：贪心 + 双指针。i 指向 s 中下一个待匹配的字符，j 扫描 t，
// 匹配上就让 i 前进。尽早匹配不会让结果变差，所以贪心是正确的。

// isSubsequence 解法一：贪心双指针 ⭐ 最优方案
// 时间复杂度 O(|s|+|t|)，空间复杂度 O(1)。
func isSubsequence(s, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

// isSubsequenceDP 解法二：预处理「下一个字符位置」表（适合大量 s 查询同一个 t 的进阶场景）
// next[j][c] 表示 t 中从位置 j 开始（含 j）字符 c 第一次出现的下标，不存在为 len(t)。
// 匹配时从 pos=0 出发，每个字符直接跳到 next[pos][c]+1。
// 时间复杂度 O(26·|t| + |s|)，空间复杂度 O(26·|t|)。
func isSubsequenceDP(s, t string) bool {
	m := len(t)
	next := make([][26]int, m+1)
	for c := range 26 {
		next[m][c] = m
	}
	for j := m - 1; j >= 0; j-- {
		next[j] = next[j+1]
		next[j][t[j]-'a'] = j
	}
	pos := 0
	for i := 0; i < len(s); i++ {
		p := next[pos][s[i]-'a']
		if p == m {
			return false
		}
		pos = p + 1
	}
	return true
}

// isSubsequenceEnumeration 解法三：枚举 + 递归（对照用）
// 对 s 的第一个字符，枚举它在 t 中的每个匹配位置，递归判断剩余部分。
// 最坏时间复杂度 O(2^|t|) 级别（实际受剪枝影响），空间复杂度 O(|s|) 递归栈。
func isSubsequenceEnumeration(s, t string) bool {
	if len(s) == 0 {
		return true
	}
	for j := 0; j < len(t); j++ {
		if t[j] == s[0] && isSubsequenceEnumeration(s[1:], t[j+1:]) {
			return true
		}
	}
	return false
}
