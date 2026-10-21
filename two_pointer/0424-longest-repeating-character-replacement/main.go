package charreplace

// 424. 替换后的最长重复字符
// https://leetcode.cn/problems/longest-repeating-character-replacement/
// 给定只含大写英文字母的字符串 s 和整数 k，最多可把 k 个字符替换成任意大写字母，
// 返回替换后「全由同一字母组成」的最长子串长度。
// 例如 s="ABAB", k=2 -> 4；s="AABABBA", k=1 -> 4。
//
// 核心思路：一个窗口能变成全相同字母，当且仅当
// 「窗口长度 - 窗口内出现最多的字母个数 <= k」（其余字母全部替换掉）。
// 用滑动窗口维护这个条件，不满足时左指针右移一位。

// characterReplacement 解法一：滑动窗口 + 历史最大频次 ⭐ 最优方案
// maxCnt 记录窗口曾经达到过的最大单字母频次，左移时不回退它：
// 答案只会在 maxCnt 变大时变大，所以用「过时的」maxCnt 不会得到错误的更大答案，
// 窗口长度只增不减或平移，最终 len(s)-left 就是答案。
// 时间复杂度 O(n)，空间复杂度 O(1)（26 个计数）。
func characterReplacement(s string, k int) int {
	var cnt [26]int
	maxCnt, left := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right] - 'A'
		cnt[c]++
		maxCnt = max(maxCnt, cnt[c])
		// 需要替换的字符数超过 k，窗口整体右移一位（长度不变）。
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

// characterReplacementEnumeration 解法二：枚举起点暴力扩展（对照用）
// 对每个起点 i 向右扩展，实时维护窗口内最大频次，满足条件就更新答案。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func characterReplacementEnumeration(s string, k int) int {
	best := 0
	for i := 0; i < len(s); i++ {
		var cnt [26]int
		maxCnt := 0
		for j := i; j < len(s); j++ {
			c := s[j] - 'A'
			cnt[c]++
			maxCnt = max(maxCnt, cnt[c])
			if j-i+1-maxCnt <= k {
				best = max(best, j-i+1)
			}
		}
	}
	return best
}
