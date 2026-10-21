package removeduplicateletters

// 316. 去除重复字母
// https://leetcode.cn/problems/remove-duplicate-letters/
// 给定只含小写字母的字符串 s，去除重复字母使每个字母只出现一次，
// 并保证返回结果的字典序最小（且必须是原串的子序列）。例如 "cbacdcbc" -> "acdb"。
//
// 核心思路：想让字典序小，就希望前面的字符尽量小。从左到右构造答案，
// 遇到一个更小的字符时，如果栈顶字符比它大、并且栈顶字符在后面还会再出现，
// 那就把栈顶扔掉（以后再取），这正是「单调栈 + 贪心」。

// removeDuplicateLetters 解法一：单调栈 + 贪心 ⭐ 最优方案
// last[c] 记录字符 c 最后出现的位置；inStack[c] 表示 c 是否已在答案里。
// 遍历到 s[i]：已在栈中则跳过；否则当栈顶 > s[i] 且栈顶在 i 之后还会出现时弹出，最后压入 s[i]。
// 时间复杂度 O(n)，空间复杂度 O(|Σ|)（栈最多 26 个字符）。
func removeDuplicateLetters(s string) string {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	var inStack [26]bool
	stack := make([]byte, 0, 26)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if inStack[c-'a'] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && last[stack[len(stack)-1]-'a'] > i {
			inStack[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		inStack[c-'a'] = true
	}
	return string(stack)
}

// removeDuplicateLettersEnumeration 解法二：逐位贪心选字符（对照用）
// 每一步从 'a' 到 'z' 尝试答案的下一个字符 c：取 c 在剩余串中第一次出现的位置 p，
// 若 s[p:] 仍包含所有尚未选的字符，就选 c，然后从 p+1 继续，并删去剩余串中的 c。
// 时间复杂度 O(|Σ|^2 · n)，空间复杂度 O(n)。
func removeDuplicateLettersEnumeration(s string) string {
	res := []byte{}
	for len(s) > 0 {
		need := [26]bool{}
		for i := 0; i < len(s); i++ {
			need[s[i]-'a'] = true
		}
		for c := byte('a'); c <= 'z'; c++ {
			if !need[c-'a'] {
				continue
			}
			p := indexByte(s, c)
			if containsAll(s[p:], need) {
				res = append(res, c)
				rest := make([]byte, 0, len(s)-p)
				for i := p + 1; i < len(s); i++ {
					if s[i] != c {
						rest = append(rest, s[i])
					}
				}
				s = string(rest)
				break
			}
		}
	}
	return string(res)
}

func indexByte(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func containsAll(s string, need [26]bool) bool {
	var has [26]bool
	for i := 0; i < len(s); i++ {
		has[s[i]-'a'] = true
	}
	for i := 0; i < 26; i++ {
		if need[i] && !has[i] {
			return false
		}
	}
	return true
}
