package strstr

// 28. 找出字符串中第一个匹配项的下标
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
// 给你两个字符串 haystack 和 needle，请你在 haystack 字符串中找出 needle 字符串
// 的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，
// 则返回 -1。约定：当 needle 为空字符串时返回 0。
// 例如 haystack="sadbutsad", needle="sad" -> 0；
// haystack="leetcode", needle="leeto" -> -1。
//
// 一句话理解题意：实现字符串查找（indexOf），返回首次出现位置。
//
// 核心思路：朴素做法是枚举每个起点逐字符比较；
// KMP 通过预处理 needle 的「最长相等前后缀」避免主串指针回退，达到线性。

// strStr 【解法1】朴素匹配 Enumeration（对照 / 直观）
// 思路：枚举 haystack 中每个可能的起点 i（范围 0..len-m），
// 从该起点开始逐字符与 needle 比较，全部匹配则返回 i。
// 时间复杂度 O((n-m+1)*m)，最坏 O(n*m)。
// 空间复杂度 O(1)。
func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	for i := 0; i+m <= n; i++ {
		j := 0
		for j < m && haystack[i+j] == needle[j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

// strStrKMP 【解法2】KMP ⭐ 最优方案
// 思路：先对 needle 计算 next 数组（next[j] 表示 needle[:j] 的最长相等前后缀长度），
// 匹配时主串指针 i 不回退，模式串指针 j 失配时跳到 next[j-1]。
// 时间复杂度 O(n+m)：构建 next O(m)，匹配 O(n)。
// 空间复杂度 O(m)：next 数组。
func strStrKMP(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	next := buildNext(needle)
	j := 0
	for i := 0; i < n; i++ {
		// 失配时回退 j 到最长相等前后缀位置。
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

// buildNext 计算 needle 的部分匹配表（最长相等真前后缀长度）。
func buildNext(needle string) []int {
	m := len(needle)
	next := make([]int, m)
	k := 0 // 当前最长相等前后缀长度
	for i := 1; i < m; i++ {
		for k > 0 && needle[i] != needle[k] {
			k = next[k-1]
		}
		if needle[i] == needle[k] {
			k++
		}
		next[i] = k
	}
	return next
}
