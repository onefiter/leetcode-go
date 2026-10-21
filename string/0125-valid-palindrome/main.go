package validpalindrome

import "unicode"

// 125. 验证回文串
// https://leetcode.cn/problems/valid-palindrome/
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，
// 字符串正着读和反着读都一样，则认为它是一个回文串。
// 只考虑字母和数字字符，忽略大小写。
// 例如 "A man, a plan, a canal: Panama" -> true；"race a car" -> false。
//
// 一句话理解题意：先过滤掉非字母数字、统一大小写，再判断是否回文。
//
// 核心思路：双指针从两端向中间走，跳过所有非字母数字字符，
// 比较时统一转小写，只要有一次不等即返回 false。

// isPalindrome 【解法1】双指针原地过滤 ⭐ 最优方案
// 思路：左右指针各自跳过非字母数字字符，再比较（忽略大小写）。
// 时间复杂度 O(n)：每个字符至多被访问一次。
// 空间复杂度 O(1)：不额外开辟字符串。
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlnum(s[i]) {
			i++
		}
		for i < j && !isAlnum(s[j]) {
			j--
		}
		if i < j {
			if toLower(s[i]) != toLower(s[j]) {
				return false
			}
			i++
			j--
		}
	}
	return true
}

// isAlnum 判断 ASCII 字符是否为字母或数字。
func isAlnum(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// toLower 将 ASCII 大写字母转为小写，其余原样返回。
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// isPalindromeFilter 【解法2】先构造过滤后的切片再双指针（对照用）
// 思路：先遍历一遍，保留字母数字并统一转小写，存入新切片，
// 再对新切片做标准双指针回文判断。逻辑直白，但多用 O(n) 空间。
// 这里用 unicode 包演示通用的大小写/字母数字判断。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func isPalindromeFilter(s string) bool {
	cs := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cs = append(cs, unicode.ToLower(r))
		}
	}
	for i, j := 0, len(cs)-1; i < j; i, j = i+1, j-1 {
		if cs[i] != cs[j] {
			return false
		}
	}
	return true
}
