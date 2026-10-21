package longestcommonprefix

// 14. 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 例如 strs=["flower","flow","flight"] -> "fl"；
// strs=["dog","racecar","car"] -> ""。
//
// 一句话理解题意：找出所有字符串开头都相同的那段最长串。
//
// 核心思路：以第一个字符串为基准，逐列比较所有字符串在该位置的字符，
// 一旦某个字符串在该位置越界或字符不同，当前位置之前即为公共前缀。

// longestCommonPrefix 【解法1】纵向逐列扫描 ⭐ 最优实用方案
// 思路：固定以 strs[0] 为参照，按列（下标 i）遍历，
// 检查其余每个字符串第 i 位是否与 strs[0][i] 相同，
// 出现越界或不同就返回 strs[0][:i]。
// 时间复杂度 O(S)：S 为所有字符总数，最坏情况下每个字符都比较一次。
// 空间复杂度 O(1)。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// longestCommonPrefixH 【解法2】横向逐个归并（对照用）
// 思路：用 prefix 保存当前公共前缀，初始为 strs[0]，
// 依次与后续字符串比较并缩短 prefix，直到匹配为止。
// 时间复杂度 O(S)，空间复杂度 O(1)。
func longestCommonPrefixH(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		// 缩短 prefix 直到它是 s 的前缀。
		k := 0
		for k < len(prefix) && k < len(s) && prefix[k] == s[k] {
			k++
		}
		prefix = prefix[:k]
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
