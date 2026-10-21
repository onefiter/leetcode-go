package kdistinct

// 340. 至多包含 K 个不同字符的最长子串
// https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/
// 给定字符串 s 和整数 k，返回至多包含 k 个不同字符的最长子串的长度。
// 例如 s="eceba", k=2 -> 3（"ece"）；s="aa", k=1 -> 2。
//
// 核心思路：159 题的推广。滑动窗口 + 计数表，
// 右指针扩张时更新字符种类数，种类超过 k 就收缩左指针直到合法，
// 每个 right 对应的合法窗口取最大长度。

// lengthOfLongestSubstringKDistinct 解法一：滑动窗口 + 计数表 ⭐ 最优方案
// cnt 记录窗口内每个字节的个数，kinds 为窗口内不同字符数。
// k=0 时窗口永远无法容纳字符，答案为 0。
// 时间复杂度 O(n)：左右指针各最多移动 n 次。
// 空间复杂度 O(1)：计数表固定 256 大小。
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var cnt [256]int
	kinds, best, left := 0, 0, 0
	for right := 0; right < len(s); right++ {
		if cnt[s[right]] == 0 {
			kinds++
		}
		cnt[s[right]]++
		for kinds > k {
			cnt[s[left]]--
			if cnt[s[left]] == 0 {
				kinds--
			}
			left++
		}
		best = max(best, right-left+1)
	}
	return best
}

// lengthOfLongestSubstringKDistinctEnumeration 解法二：枚举起点暴力扩展（对照用）
// 对每个起点向右扩展并统计种类数，超过 k 即停止。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func lengthOfLongestSubstringKDistinctEnumeration(s string, k int) int {
	best := 0
	for i := 0; i < len(s); i++ {
		var seen [256]bool
		kinds := 0
		for j := i; j < len(s); j++ {
			if !seen[s[j]] {
				seen[s[j]] = true
				kinds++
			}
			if kinds > k {
				break
			}
			best = max(best, j-i+1)
		}
	}
	return best
}
