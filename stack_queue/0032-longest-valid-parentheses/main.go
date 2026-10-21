package longestvalid

// 32. 最长有效括号
// https://leetcode.cn/problems/longest-valid-parentheses/
// 给定只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
// 例如 ")()())" -> 4（"()()"）。
//
// 核心思路：有效子串的特点是「左右括号数量相等，且任意前缀中左括号不少于右括号」。
// 可以用栈记录「最后一个没被匹配的位置」作为分界；也可以用 DP 以每个位置结尾来递推；
// 还可以左右各扫一遍计数，不需要额外空间。

// longestValidParentheses 解法一：栈 ⭐ 最优方案（之一）
// 栈底始终保存「最后一个没有被匹配的右括号下标」，初始放 -1 作为哨兵。
// 遇到 '(' 压入下标；遇到 ')' 先弹出栈顶：
//   - 弹完栈空，说明这个 ')' 没人匹配，把它的下标压入作为新的分界；
//   - 栈非空，当前有效长度 = i - 栈顶下标。
//
// 时间复杂度 O(n)，空间复杂度 O(n)。
func longestValidParentheses(s string) int {
	best := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			best = max(best, i-stack[len(stack)-1])
		}
	}
	return best
}

// longestValidParenthesesDP 解法二：动态规划
// dp[i] 表示以 s[i] 结尾的最长有效子串长度，只有 s[i]==')' 时可能非 0：
//   - s[i-1]=='('：形如 "...()"，dp[i] = dp[i-2] + 2；
//   - s[i-1]==')'：形如 "...))"，看 dp[i-1] 那段之前的字符 j=i-dp[i-1]-1，
//     若 s[j]=='('，则 dp[i] = dp[i-1] + 2 + dp[j-1]（把 j 之前紧挨着的有效段也接上）。
//
// 时间复杂度 O(n)，空间复杂度 O(n)。
func longestValidParenthesesDP(s string) int {
	n := len(s)
	dp := make([]int, n)
	best := 0
	for i := 1; i < n; i++ {
		if s[i] != ')' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else if j := i - dp[i-1] - 1; j >= 0 && s[j] == '(' {
			dp[i] = dp[i-1] + 2
			if j >= 1 {
				dp[i] += dp[j-1]
			}
		}
		best = max(best, dp[i])
	}
	return best
}

// longestValidParenthesesCount 解法三：左右双向计数 ⭐ 空间最优
// 从左往右扫，分别统计 left、right：相等时更新答案；right > left 时说明断开，清零。
// 但像 "(()" 这种左括号一直多的情况，从左扫永远不相等，所以再从右往左扫一遍，
// 规则对称（left > right 时清零）。两遍取最大。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func longestValidParenthesesCount(s string) int {
	best := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			best = max(best, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			best = max(best, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return best
}

// longestValidParenthesesEnumeration 解法四：枚举起点暴力（对照用）
// 固定起点 i 向右扩展，用计数 bal 维护「左括号 - 右括号」：
// bal < 0 说明从 i 开始已不可能再有效，停止；bal == 0 时 [i, j] 是有效串。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func longestValidParenthesesEnumeration(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		bal := 0
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				bal++
			} else {
				bal--
			}
			if bal < 0 {
				break
			}
			if bal == 0 {
				best = max(best, j-i+1)
			}
		}
	}
	return best
}
