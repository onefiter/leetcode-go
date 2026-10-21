package findanagrams

// 438. 找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
// 给定两个字符串 s 和 p，找到 s 中所有 p 的「异位词」的子串，返回这些子串的
// 起始索引。异位词指字母相同、顺序可以不同的字符串。s、p 只含小写字母。
//
// 核心思路：固定长度滑动窗口 + 计数比较。
//   窗口长度恒等于 len(p)。维护窗口内各字符计数，与 p 的计数比较，
//   相等则当前窗口是一个异位词，记录左端点。窗口右移时右端点字符入窗、
//   左端点字符出窗，计数增量更新，无需每次重建。

// findAnagrams 解法一：固定窗口滑动 + 计数数组 ⭐ 最优方案
// 用两个长度 26 的数组分别记录 p 和窗口的字符频次，逐步滑动并比较。
// 返回的下标天然按升序排列（从左到右遍历）。
// 时间复杂度 O(n)：每个字符入窗出窗各一次，比较数组是 O(26)=O(1)。
// 空间复杂度 O(1)：两个固定 26 大小的计数数组（不计返回值）。
func findAnagrams(s, p string) []int {
	res := []int{}
	if len(s) < len(p) || len(p) == 0 {
		return res
	}
	var need, window [26]int
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++
		if i >= len(p) {
			window[s[i-len(p)]-'a']-- // 窗口左端移出
		}
		if i >= len(p)-1 && window == need {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

// findAnagramsBrute 解法二：枚举每个起点并逐个比较计数（对照用）
// 对每个可能的起点截取长度 len(p) 的子串，统计其字符频次与 p 比较。
// 时间复杂度 O(n * (len(p)+26))，空间复杂度 O(1)。
func findAnagramsBrute(s, p string) []int {
	res := []int{}
	if len(s) < len(p) || len(p) == 0 {
		return res
	}
	var need [26]int
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
	}
	for i := 0; i+len(p) <= len(s); i++ {
		var window [26]int
		for j := i; j < i+len(p); j++ {
			window[s[j]-'a']++
		}
		if window == need {
			res = append(res, i)
		}
	}
	return res
}
