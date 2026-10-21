package removeouterparens

import "strings"

// 1021. 删除最外层的括号
// https://leetcode.cn/problems/remove-outermost-parentheses/
// 有效括号字符串 s 可以唯一分解为若干「原语」（不能再拆的非空有效括号串）的拼接。
// 删除每个原语最外层的一对括号后返回结果。例如 "(()())(())" -> "()()()"。
//
// 核心思路：用计数器 depth 表示当前嵌套深度。
// 遇到 '(' 时，若进入前 depth>0 说明不是最外层，保留；遇到 ')' 时，若退出后 depth>0 也保留。

// removeOuterParentheses 解法一：深度计数 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(n)（结果字符串）。
func removeOuterParentheses(s string) string {
	var sb strings.Builder
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if depth > 0 {
				sb.WriteByte('(')
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				sb.WriteByte(')')
			}
		}
	}
	return sb.String()
}

// removeOuterParenthesesEnumeration 解法二：双指针切分原语（对照用）
// start 记录当前原语起点，扫描到平衡点（计数归零）时得到原语 s[start:i+1]，
// 取其内部 s[start+1:i] 拼接到结果，再把 start 移到 i+1。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func removeOuterParenthesesEnumeration(s string) string {
	res := ""
	start, bal := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			bal++
		} else {
			bal--
		}
		if bal == 0 {
			res += s[start+1 : i]
			start = i + 1
		}
	}
	return res
}
