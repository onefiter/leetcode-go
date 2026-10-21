package decodestring

import "strings"

// 394. 字符串解码
// https://leetcode.cn/problems/decode-string/
// 编码规则为 k[encoded_string]，表示方括号内的字符串重复 k 次（k 为正整数），可以嵌套。
// 例如 "3[a2[c]]" -> "accaccacc"，"2[abc]3[cd]ef" -> "abcabccdcdcdef"。
//
// 核心思路：括号嵌套 = 先处理最内层，天然用栈。
// 遇到 '[' 时把「当前已拼好的串」和「重复次数」压栈，进入新的一层从空串开始拼；
// 遇到 ']' 时把这一层的串重复 k 次，接到上一层的串后面。

// decodeString 解法一：双栈 ⭐ 最优方案
// numStack 存每层的重复次数，strStack 存进入该层前已拼好的前缀。
// 时间复杂度 O(S)，S 为解码后字符串长度；空间复杂度 O(S)。
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	var cur strings.Builder
	k := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			k = k*10 + int(c-'0')
		case c == '[':
			numStack = append(numStack, k)
			strStack = append(strStack, cur.String())
			cur.Reset()
			k = 0
		case c == ']':
			n := numStack[len(numStack)-1]
			prev := strStack[len(strStack)-1]
			numStack = numStack[:len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
			inner := strings.Repeat(cur.String(), n)
			cur.Reset()
			cur.WriteString(prev)
			cur.WriteString(inner)
		default:
			cur.WriteByte(c)
		}
	}
	return cur.String()
}

// decodeStringRecursive 解法二：递归
// dfs(i) 从位置 i 开始解码，直到遇到 ']' 或串尾，返回解码结果和停下的位置。
// 遇到数字就读出 k，跳过 '[' 后递归解码括号内部，再重复 k 次。
// 时间复杂度 O(S)，空间复杂度 O(S)（含递归栈）。
func decodeStringRecursive(s string) string {
	res, _ := dfs(s, 0)
	return res
}

func dfs(s string, i int) (string, int) {
	var sb strings.Builder
	for i < len(s) && s[i] != ']' {
		if s[i] >= '0' && s[i] <= '9' {
			k := 0
			for s[i] >= '0' && s[i] <= '9' {
				k = k*10 + int(s[i]-'0')
				i++
			}
			inner, next := dfs(s, i+1) // i 指向 '['
			sb.WriteString(strings.Repeat(inner, k))
			i = next + 1 // 跳过 ']'
		} else {
			sb.WriteByte(s[i])
			i++
		}
	}
	return sb.String(), i
}

// decodeStringEnumeration 解法三：反复展开最内层括号（对照用）
// 每轮找到第一个 ']'，它前面最近的 '[' 与之构成最内层括号，再向前读出数字 k，
// 把 "k[xxx]" 替换为 xxx 重复 k 次，直到没有括号。
// 时间复杂度 O(S · 括号数)，空间复杂度 O(S)。
func decodeStringEnumeration(s string) string {
	for {
		r := strings.IndexByte(s, ']')
		if r == -1 {
			return s
		}
		l := strings.LastIndexByte(s[:r], '[')
		d := l
		for d > 0 && s[d-1] >= '0' && s[d-1] <= '9' {
			d--
		}
		k := 0
		for i := d; i < l; i++ {
			k = k*10 + int(s[i]-'0')
		}
		s = s[:d] + strings.Repeat(s[l+1:r], k) + s[r+1:]
	}
}
