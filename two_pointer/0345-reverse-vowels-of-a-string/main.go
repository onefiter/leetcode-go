package reversevowels

import "strings"

// 345. 反转字符串中的元音字母
// https://leetcode.cn/problems/reverse-vowels-of-a-string/
// 给定字符串 s，仅反转其中所有的元音字母（a e i o u，大小写都算），其余字符位置不变。
// 例如 "IceCreAm" -> "AceCreIm"，"leetcode" -> "leotcede"。
//
// 核心思路：左右双指针分别从两端找元音，找到一对就交换，然后继续向中间收缩。

func isVowel(c byte) bool {
	return strings.IndexByte("aeiouAEIOU", c) >= 0
}

// reverseVowels 解法一：对撞双指针 ⭐ 最优方案
// 左指针跳过非元音，右指针跳过非元音，两者都停在元音上时交换。
// 时间复杂度 O(n)，空间复杂度 O(n)（Go 字符串不可变，需要一个 []byte 副本）。
func reverseVowels(s string) string {
	b := []byte(s)
	l, r := 0, len(b)-1
	for l < r {
		for l < r && !isVowel(b[l]) {
			l++
		}
		for l < r && !isVowel(b[r]) {
			r--
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

// reverseVowelsEnumeration 解法二：收集后回填（对照用）
// 先按顺序把所有元音取出来，再扫描一遍，遇到元音位置就从收集结果的末尾倒序回填。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func reverseVowelsEnumeration(s string) string {
	b := []byte(s)
	vowels := []byte{}
	for _, c := range b {
		if isVowel(c) {
			vowels = append(vowels, c)
		}
	}
	k := len(vowels) - 1
	for i, c := range b {
		if isVowel(c) {
			b[i] = vowels[k]
			k--
		}
	}
	return string(b)
}
