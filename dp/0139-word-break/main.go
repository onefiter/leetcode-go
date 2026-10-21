package wordbreak

// 139. 单词拆分
// https://leetcode.cn/problems/word-break/
// 给一个字符串 s 和字符串列表 wordDict 作为字典，判断能否用字典中的单词拼接出 s。
// 字典中的单词可以重复使用。
//
// 核心思路：完全背包（排列型）可行性问题。dp[i] 表示 s 前 i 个字符能否被拆分，dp[0] = true。
// 转移：dp[i] = 存在 j < i 使得 dp[j] && s[j:i] 在字典中。
// 遍历顺序：i 正序（外层背包容量），内层枚举切分点 j；
// 可用字典中最长单词长度剪枝，只枚举 i-maxLen <= j < i。

// wordBreak 解法一：一维 DP + 哈希集合 ⭐ 最优方案
// 时间复杂度 O(n * L * L)（L 为最长单词长度，含子串哈希），空间复杂度 O(n + 字典总长)。
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(0, i-maxLen); j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// wordBreakBFS 解法二：BFS（对照用）
// 把下标看作图节点，若 s[i:j] 在字典中则 i -> j 有边，问 0 能否到达 n。
// visited 保证每个起点只扩展一次。
// 时间复杂度 O(n^2 * L)，空间复杂度 O(n + 字典总长)。
func wordBreakBFS(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
	}
	n := len(s)
	visited := make([]bool, n+1)
	queue := []int{0}
	visited[0] = true
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if start == n {
			return true
		}
		for end := start + 1; end <= n; end++ {
			if !visited[end] && words[s[start:end]] {
				visited[end] = true
				queue = append(queue, end)
			}
		}
	}
	return false
}
