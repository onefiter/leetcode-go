package substrconcat

// 30. 串联所有单词的子串
// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
// 给定字符串 s 和一组「长度相同」的单词 words，找出 s 中所有恰好由 words 中
// 全部单词（任意顺序、每个用一次）串联而成的子串的起始下标。
// 例如 s="barfoothefoobarman", words=["foo","bar"] -> [0,9]。
//
// 核心思路：所有单词等长 wl，所以 s 可以按 wl 「切块」。
// 起点 i 按 i%wl 分成 wl 组，每组内部按单词为单位滑动窗口：
// 窗口右端每次吞进一个单词，某单词超量就从左边吐单词，
// 窗口恰好装满 len(words) 个单词时记录左端点。

// findSubstring 解法一：按单词长度分组的滑动窗口 ⭐ 最优方案
// 对每个偏移 off ∈ [0, wl)，left/right 都以 wl 为步长前进：
//   - 取出 right 处的单词 w；不在 words 里则窗口清空，left 跳到 right+wl；
//   - 否则计数 +1，若 w 超出需要的次数，就不断把 left 处的单词移出；
//   - 窗口单词数等于 len(words) 时，left 是一个答案。
//
// 时间复杂度 O(wl * n)：每组每个单词最多进出窗口一次，每次切片/哈希 O(wl)。
// 空间复杂度 O(m * wl)：m 为单词个数，存放计数表。
func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(words) == 0 || len(words[0]) == 0 {
		return res
	}
	wl, m, n := len(words[0]), len(words), len(s)
	if wl*m > n {
		return res
	}
	need := make(map[string]int, m)
	for _, w := range words {
		need[w]++
	}
	for off := 0; off < wl; off++ {
		window := map[string]int{}
		count := 0 // 窗口内的单词个数
		left := off
		for right := off; right+wl <= n; right += wl {
			w := s[right : right+wl]
			if _, ok := need[w]; !ok {
				// 遇到不在词表里的单词，窗口整体作废。
				clear(window)
				count = 0
				left = right + wl
				continue
			}
			window[w]++
			count++
			// w 超量：从左侧吐出单词，直到 w 的数量合法。
			for window[w] > need[w] {
				window[s[left:left+wl]]--
				count--
				left += wl
			}
			if count == m {
				res = append(res, left)
			}
		}
	}
	return res
}

// findSubstringEnumeration 解法二：逐起点暴力检查（对照用）
// 枚举每个起点 i，从 i 开始切出 m 个长度为 wl 的单词逐个计数，
// 一旦出现超量（含不在词表里，need 为 0）的单词就提前失败。
// 时间复杂度 O(n * m * wl)，空间复杂度 O(m * wl)。
func findSubstringEnumeration(s string, words []string) []int {
	res := []int{}
	if len(words) == 0 || len(words[0]) == 0 {
		return res
	}
	wl, m := len(words[0]), len(words)
	need := map[string]int{}
	for _, w := range words {
		need[w]++
	}
	for i := 0; i+wl*m <= len(s); i++ {
		seen := map[string]int{}
		ok := true
		for j := 0; j < m; j++ {
			w := s[i+j*wl : i+(j+1)*wl]
			seen[w]++
			if seen[w] > need[w] {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, i)
		}
	}
	return res
}
