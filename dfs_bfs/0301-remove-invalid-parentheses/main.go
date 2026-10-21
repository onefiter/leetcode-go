package removeparens

// 301. 删除无效的括号
// https://leetcode.cn/problems/remove-invalid-parentheses/
// 给一个由括号和字母组成的字符串 s，删除最小数量的无效括号使字符串有效，
// 返回所有可能的结果（去重），顺序任意。
//
// 核心思路：先一次遍历算出最少要删的左括号数 lRemove 和右括号数 rRemove，
// 再 DFS 逐个尝试删除括号，删除数量用完且串合法即为答案。
// 去重剪枝：连续相同的括号只从第一个开始删，删哪一个结果都一样。

// isValid 判断字符串括号是否匹配。
func isValid(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			cnt++
		case ')':
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

// removeInvalidParentheses 解法一：计算删除数 + DFS 回溯剪枝 ⭐ 最优方案
// 时间复杂度 O(n*2^n)（实际因剪枝远小于此），空间复杂度 O(n^2)（递归栈 × 子串）。
func removeInvalidParentheses(s string) []string {
	lRemove, rRemove := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lRemove++
		} else if s[i] == ')' {
			if lRemove > 0 {
				lRemove--
			} else {
				rRemove++
			}
		}
	}
	res := []string{}
	var dfs func(str string, start, l, r int)
	dfs = func(str string, start, l, r int) {
		if l == 0 && r == 0 {
			if isValid(str) {
				res = append(res, str)
			}
			return
		}
		for i := start; i < len(str); i++ {
			if i > start && str[i] == str[i-1] {
				continue // 连续相同括号只删第一个，避免重复结果
			}
			if l+r > len(str)-i {
				return // 剩余字符不够删
			}
			next := str[:i] + str[i+1:]
			if str[i] == '(' && l > 0 {
				dfs(next, i, l-1, r)
			} else if str[i] == ')' && r > 0 {
				dfs(next, i, l, r-1)
			}
		}
	}
	dfs(s, 0, lRemove, rRemove)
	return res
}

// removeInvalidParenthesesBFS 解法二：BFS 按删除个数逐层扩展（对照用）
// 第 k 层是删掉 k 个括号得到的所有字符串（用 set 去重），
// 第一次在某层找到合法串时，该层所有合法串就是删除最少的答案。
// 时间复杂度 O(n*2^n)，空间复杂度 O(n*C(n, n/2))。
func removeInvalidParenthesesBFS(s string) []string {
	level := map[string]bool{s: true}
	for len(level) > 0 {
		res := []string{}
		for str := range level {
			if isValid(str) {
				res = append(res, str)
			}
		}
		if len(res) > 0 {
			return res
		}
		next := map[string]bool{}
		for str := range level {
			for i := 0; i < len(str); i++ {
				if i > 0 && str[i] == str[i-1] {
					continue // 删连续相同字符中的任意一个结果相同
				}
				if str[i] == '(' || str[i] == ')' {
					next[str[:i]+str[i+1:]] = true
				}
			}
		}
		level = next
	}
	return []string{""}
}
