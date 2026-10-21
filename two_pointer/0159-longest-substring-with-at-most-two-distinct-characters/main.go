package twodistinct

// 159. 至多包含两个不同字符的最长子串
// https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/
// 给定字符串 s，找出至多包含两个不同字符的最长子串 t，返回其长度。
// 例如 s="eceba" -> 3（"ece"）；s="ccaabbb" -> 5（"aabbb"）。
//
// 核心思路：滑动窗口。右指针不断扩张，用计数表记录窗口内每个字符的个数；
// 一旦窗口内不同字符超过 2 个，就移动左指针直到恢复合法，
// 每一步窗口都是「以 right 结尾的最长合法子串」，取最大长度即可。

// lengthOfLongestSubstringTwoDistinct 解法一：滑动窗口 + 计数表 ⭐ 最优方案
// cnt 记录窗口内每个字节出现次数，kinds 记录不同字符数。
// kinds > 2 时收缩左端，直到某个字符计数归零、kinds 回到 2。
// 时间复杂度 O(n)：左右指针各最多走 n 步。
// 空间复杂度 O(1)：计数表固定 256 大小。
func lengthOfLongestSubstringTwoDistinct(s string) int {
	var cnt [256]int
	kinds, best, left := 0, 0, 0
	for right := 0; right < len(s); right++ {
		if cnt[s[right]] == 0 {
			kinds++
		}
		cnt[s[right]]++
		for kinds > 2 {
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

// lengthOfLongestSubstringTwoDistinctEnumeration 解法二：枚举起点暴力扩展（对照用）
// 对每个起点 i 向右扩展，记录出现过的字符种类，种类超过 2 就停止。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func lengthOfLongestSubstringTwoDistinctEnumeration(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		var seen [256]bool
		kinds := 0
		for j := i; j < len(s); j++ {
			if !seen[s[j]] {
				seen[s[j]] = true
				kinds++
			}
			if kinds > 2 {
				break
			}
			best = max(best, j-i+1)
		}
	}
	return best
}
