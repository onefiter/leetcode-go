package longestsubstring

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 给定一个字符串 s，找出其中不含有重复字符的「最长子串」的长度。
//
// 核心思路：滑动窗口 + 记录字符最近出现位置。
//   右指针不断扩张窗口，用哈希表记录每个字符最近一次出现的下标。
//   当右指针字符在窗口内已出现过时，把左指针跳到「该字符上次位置 + 1」，
//   始终保持窗口内无重复字符，过程中更新最大窗口长度。

// lengthOfLongestSubstring 解法一：滑动窗口 ⭐ 最优方案
// lastSeen[c] 记录字符 c 最近一次出现的下标。遇到窗口内的重复字符时，
// 左边界直接跳过去，避免逐格收缩。
// 时间复杂度 O(n)：右指针遍历一遍，左指针单调不回退。
// 空间复杂度 O(1)：字符集固定（最多 128 个 ASCII 字符）。
func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	best, left := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if idx, ok := lastSeen[c]; ok && idx >= left {
			left = idx + 1 // 跳过上一次出现位置
		}
		lastSeen[c] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// lengthOfLongestSubstringBrute 解法二：枚举所有子串并用集合判重（对照用）
// 外层固定起点，内层扩张并用集合检查是否出现重复，重复即停止扩张。
// 时间复杂度 O(n^2)，空间复杂度 O(min(n, 字符集大小))。
func lengthOfLongestSubstringBrute(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			if j-i+1 > best {
				best = j - i + 1
			}
		}
	}
	return best
}
