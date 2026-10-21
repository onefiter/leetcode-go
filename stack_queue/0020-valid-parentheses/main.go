package validparentheses

import "strings"

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/
// 给定只包含 '(' ')' '{' '}' '[' ']' 的字符串 s，判断括号是否有效：
// 左括号必须用相同类型的右括号闭合，且闭合顺序正确。
//
// 核心思路：后出现的左括号要先被闭合，这正是「后进先出」——用栈。
// 遇到左括号入栈；遇到右括号就看栈顶是不是与之配对的左括号，是则弹出，否则无效。
// 最后栈为空才说明全部配对成功。

// isValid 解法一：栈 ⭐ 最优方案
// 用 map 记录「右括号 -> 对应左括号」，遇到右括号时与栈顶比较。
// 小优化：长度为奇数一定无效。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if open, ok := pairs[c]; ok {
			// 右括号：栈顶必须是对应的左括号。
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

// isValidEnumeration 解法二：反复删除相邻配对（对照用）
// 有效括号串里一定存在相邻的 "()"、"[]" 或 "{}"，不断把它们删掉，
// 最后能删成空串就有效；某一轮删不动了还不为空就无效。
// 时间复杂度 O(n^2)：最多 n/2 轮，每轮 O(n)。空间复杂度 O(n)。
func isValidEnumeration(s string) bool {
	for {
		next := strings.ReplaceAll(s, "()", "")
		next = strings.ReplaceAll(next, "[]", "")
		next = strings.ReplaceAll(next, "{}", "")
		if next == s {
			return s == ""
		}
		s = next
	}
}
