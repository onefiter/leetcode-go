package checkinclusion

// 567. 字符串的排列
// https://leetcode.cn/problems/permutation-in-string/
// 给你两个字符串 s1 和 s2，判断 s2 是否包含 s1 的某个排列（即 s1 的某个排列是 s2 的子串）。
// 例如 s1="ab", s2="eidbaooo" -> true（"ba" 是 "ab" 的排列）。
//
// 一句话理解题意：s2 里有没有一段长度为 len(s1) 的子串，它和 s1 的字母种类及个数完全一样。
//
// 核心思路：窗口长度固定为 len(s1)，在 s2 上滑动，维护窗口内每个字母的计数，
// 与 s1 的计数比较即可。用 diff 记录「计数不相等的字母个数」，可以 O(1) 判断是否匹配。

// checkInclusion 解法一：定长滑动窗口 + 差异计数 ⭐ 最优方案
// cnt[c] = 窗口中 c 的个数 - s1 中 c 的个数；diff = cnt 不为 0 的字母个数。
// 窗口右移时，进来一个字符 cnt++，出去一个字符 cnt--，只需更新这两个字母对 diff 的影响。
// diff == 0 说明窗口恰好是 s1 的一个排列。
// 时间复杂度 O(n+m)，空间复杂度 O(1)（字母表大小 26）。
func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt [26]int
	for i := 0; i < n; i++ {
		cnt[s1[i]-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		in, out := s2[i]-'a', s2[i-n]-'a'
		if in == out {
			continue // 进出同一个字母，计数不变
		}
		// 字符 in 进入窗口
		if cnt[in] == 0 {
			diff++
		}
		cnt[in]++
		if cnt[in] == 0 {
			diff--
		}
		// 字符 out 离开窗口
		if cnt[out] == 0 {
			diff++
		}
		cnt[out]--
		if cnt[out] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

// checkInclusionArrayCompare 解法二：定长滑动窗口 + 整体比较计数数组
// 维护两个 [26]int 计数数组，每滑动一步直接比较两个数组是否相等（Go 数组可直接 ==）。
// 时间复杂度 O(26·m)，空间复杂度 O(1)。
func checkInclusionArrayCompare(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var need, win [26]int
	for i := 0; i < n; i++ {
		need[s1[i]-'a']++
		win[s2[i]-'a']++
	}
	if need == win {
		return true
	}
	for i := n; i < m; i++ {
		win[s2[i]-'a']++
		win[s2[i-n]-'a']--
		if need == win {
			return true
		}
	}
	return false
}

// checkInclusionEnumeration 解法三：枚举每个起点重新计数（对照用）
// 对 s2 每个长度为 n 的子串都重新统计字母个数，再与 s1 的计数比较。
// 时间复杂度 O(n·m)，空间复杂度 O(1)。
func checkInclusionEnumeration(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	var need [26]int
	for i := 0; i < n; i++ {
		need[s1[i]-'a']++
	}
	for start := 0; start+n <= m; start++ {
		var win [26]int
		for j := start; j < start+n; j++ {
			win[s2[j]-'a']++
		}
		if win == need {
			return true
		}
	}
	return false
}
