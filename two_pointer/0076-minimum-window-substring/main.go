package minwindow

// 76. 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/
// 给定字符串 s 和 t，返回 s 中涵盖 t 所有字符（含重复）的最小子串。
// 若 s 中不存在这样的子串，返回空字符串 ""。LeetCode 保证答案唯一。
//
// 核心思路：滑动窗口 + 计数。
//   用 need 记录 t 中每个字符还「需要」的数量，用 missing 记录窗口内
//   尚未满足的字符总数。右指针不断扩张窗口并补齐需求，当 missing==0
//   表示窗口已覆盖 t，此时尝试收缩左指针以求更小窗口，并更新答案。

// minWindow 解法一：滑动窗口 ⭐ 最优方案
// need[c] > 0 表示还缺 c；<=0 表示窗口内该字符已够（甚至多余）。
// missing 为窗口还缺的字符总数，missing==0 即完成覆盖。
// 时间复杂度 O(|s| + |t|)：左右指针各最多走一遍 s。
// 空间复杂度 O(1)：计数表大小固定（ASCII 128）。
func minWindow(s, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t)
	start, end := 0, 0 // 记录最优窗口 [start, end)
	left := 0
	for right := 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			missing-- // 该字符是 t 需要的，补上一个
		}
		need[s[right]]--
		for missing == 0 { // 窗口已覆盖 t，尝试收缩左边界
			if end == 0 || right-left+1 < end-start {
				start, end = left, right+1
			}
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++ // 弹出的字符重新变成缺失
			}
			left++
		}
	}
	return s[start:end]
}

// minWindowBrute 解法二：枚举所有子串并检查是否覆盖 t（对照用）
// 外层枚举左端点，内层扩张右端点并维护计数，一旦覆盖就记录并换下一个左端点。
// 时间复杂度 O(|s|^2)，空间复杂度 O(1)（计数表固定）。
func minWindowBrute(s, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	best := ""
	for i := 0; i < len(s); i++ {
		var window [128]int
		missing := len(t)
		for j := i; j < len(s); j++ {
			if window[s[j]] < need[s[j]] {
				missing--
			}
			window[s[j]]++
			if missing == 0 {
				if best == "" || j-i+1 < len(best) {
					best = s[i : j+1]
				}
				break
			}
		}
	}
	return best
}
