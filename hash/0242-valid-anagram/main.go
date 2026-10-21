package validanagram

import (
	"slices"
	"sort"
)

// 242. 有效的字母异位词
// https://leetcode.cn/problems/valid-anagram/
// 给定两个字符串 s 和 t，判断 t 是否是 s 的字母异位词。
// 字母异位词：由相同字母、相同数量、不同顺序组成，例如 s="anagram", t="nagaram" -> true。
//
// 核心思路：长度不同直接 false；否则统计每个字符出现次数，两者完全一致即为异位词。

// isAnagram 【解法1】哈希计数 ⭐ 最优
// 用一个计数表：遍历 s 时 +1，遍历 t 时 -1，最后所有计数为 0 则是异位词。
// 这里用 map[rune]int 以支持 Unicode（题目进阶要求）。
// 时间 O(n)：两次遍历。
// 空间 O(k)：k 为不同字符数。
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := make(map[rune]int)
	for _, c := range s {
		cnt[c]++
	}
	for _, c := range t {
		cnt[c]--
		if cnt[c] < 0 {
			return false
		}
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}

// isAnagramArray 【解法2】定长数组计数（仅小写字母）
// 题目基础版本只含小写字母，用长度 26 的数组计数，常数更小。
// 时间 O(n)，空间 O(1)（固定 26）。
func isAnagramArray(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}

// isAnagramSort 【解法3】Enumeration 暴力：排序后逐字符比较（对照用）
// 把两个字符串的字符分别排序，相等即为异位词。
// 时间 O(n*logn)：排序开销。
// 空间 O(n)：排序用的字符切片。
func isAnagramSort(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := []rune(s)
	b := []rune(t)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return slices.Equal(a, b)
}
