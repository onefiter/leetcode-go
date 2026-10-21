package lettercombinations

// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合（顺序任意）。
// 数字到字母的映射与电话按键相同，digits 为空时返回空列表。
//
// 核心思路：每一位数字对应一组候选字母，相当于一棵「层数 = len(digits)」的决策树，
// 回溯时每层从当前数字的字母里选一个，走到叶子就得到一个组合。

var phone = [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// letterCombinations 解法一：回溯 DFS ⭐ 最优方案
// 时间复杂度 O(4^n * n)，空间复杂度 O(n)（不计结果，递归深度 + 路径缓冲区）。
func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	path := make([]byte, len(digits))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			res = append(res, string(path))
			return
		}
		for _, c := range []byte(phone[digits[i]-'0']) {
			path[i] = c // 直接覆盖第 i 位，无需显式撤销
			dfs(i + 1)
		}
	}
	dfs(0)
	return res
}

// letterCombinationsIterative 解法二：BFS 逐层扩展（对照用）
// 维护当前所有前缀，每读入一位数字，就把每个前缀分别拼上该数字的每个字母。
// 时间复杂度 O(4^n * n)，空间复杂度 O(4^n * n)（需要保存整层前缀）。
func letterCombinationsIterative(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	queue := []string{""}
	for i := 0; i < len(digits); i++ {
		letters := phone[digits[i]-'0']
		next := make([]string, 0, len(queue)*len(letters))
		for _, prefix := range queue {
			for j := 0; j < len(letters); j++ {
				next = append(next, prefix+string(letters[j]))
			}
		}
		queue = next
	}
	return queue
}
