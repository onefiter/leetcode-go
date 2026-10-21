package repeateddna

// 187. 重复的DNA序列
// https://leetcode.cn/problems/repeated-dna-sequences/
// DNA 序列由 'A','C','G','T' 组成。给定字符串 s，返回所有在 s 中出现不止一次、
// 长度为 10 的子串（顺序任意）。
// 例如 s="AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT" -> ["AAAAACCCCC","CCCCCAAAAA"]。
//
// 核心思路：用一个长度固定为 10 的窗口从左往右滑，
// 用哈希表统计每个窗口子串出现的次数，次数恰好变成 2 时加入答案（避免重复加入）。
// 进一步：4 种字符各用 2 bit 编码，10 个字符正好 20 bit，
// 窗口右移时「左移 2 位 + 并入新字符 + 掩掉最高 2 位」即可 O(1) 得到新编码。

const seqLen = 10

// findRepeatedDnaSequences 解法一：滑动窗口 + 位运算编码 + 哈希 ⭐ 最优方案
// A=00, C=01, G=10, T=11，窗口编码 x 只保留低 20 位。
// 每右移一位：x = ((x << 2) | code(s[i])) & (1<<20 - 1)，无需切片和字符串哈希。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	if len(s) <= seqLen {
		return res
	}
	code := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	const mask = 1<<(2*seqLen) - 1
	x := 0
	for i := 0; i < seqLen-1; i++ {
		x = x<<2 | code[s[i]]
	}
	cnt := map[int]int{}
	for i := seqLen - 1; i < len(s); i++ {
		x = (x<<2 | code[s[i]]) & mask
		cnt[x]++
		if cnt[x] == 2 {
			res = append(res, s[i-seqLen+1:i+1])
		}
	}
	return res
}

// findRepeatedDnaSequencesHash 解法二：滑动窗口 + 字符串哈希表
// 直接把每个长度为 10 的子串作为 key 计数，次数为 2 时加入答案。
// 时间复杂度 O(n * L)，L=10 为切片与哈希的代价；空间复杂度 O(n * L)。
func findRepeatedDnaSequencesHash(s string) []string {
	res := []string{}
	cnt := map[string]int{}
	for i := 0; i+seqLen <= len(s); i++ {
		sub := s[i : i+seqLen]
		cnt[sub]++
		if cnt[sub] == 2 {
			res = append(res, sub)
		}
	}
	return res
}

// findRepeatedDnaSequencesEnumeration 解法三：两两比较所有窗口（对照用）
// 对每个起点 i，向后找是否存在另一个起点 j 的子串与之相同，
// 并用集合去重（同一子串只收一次）。
// 时间复杂度 O(n^2 * L)，空间复杂度 O(n * L)。
func findRepeatedDnaSequencesEnumeration(s string) []string {
	res := []string{}
	added := map[string]bool{}
	for i := 0; i+seqLen <= len(s); i++ {
		sub := s[i : i+seqLen]
		if added[sub] {
			continue
		}
		for j := i + 1; j+seqLen <= len(s); j++ {
			if s[j:j+seqLen] == sub {
				added[sub] = true
				res = append(res, sub)
				break
			}
		}
	}
	return res
}
