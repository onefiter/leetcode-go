package firstuniqchar

// 387. 字符串中的第一个唯一字符
// https://leetcode.cn/problems/first-unique-character-in-a-string/
// 给定一个字符串 s，找到它的第一个「不重复」的字符，并返回它的索引；
// 如果不存在，则返回 -1。s 只包含小写字母。
//
// 核心思路：哈希表（或定长数组）统计频次。
//   先遍历一遍统计每个字符出现的次数，再按原顺序遍历一遍，
//   返回第一个出现次数为 1 的字符下标。

// firstUniqChar 解法一：两次遍历 + 频次数组 ⭐ 最优方案
// s 只含小写字母，用长度 26 的数组计数比 map 更快、常数更小。
// 时间复杂度 O(n)：两次线性遍历。
// 空间复杂度 O(1)：固定 26 个字母的计数数组。
func firstUniqChar(s string) int {
	var count [26]int
	for _, c := range s {
		count[c-'a']++
	}
	for i, c := range s {
		if count[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

// firstUniqCharBrute 解法二：对每个字符向后扫描判断是否重复（对照用）
// 对位置 i 的字符，扫描整个串看是否只出现一次。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func firstUniqCharBrute(s string) int {
	for i := 0; i < len(s); i++ {
		unique := true
		for j := 0; j < len(s); j++ {
			if j != i && s[j] == s[i] {
				unique = false
				break
			}
		}
		if unique {
			return i
		}
	}
	return -1
}
