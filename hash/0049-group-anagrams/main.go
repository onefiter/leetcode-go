package groupanagrams

import (
	"slices"
	"sort"
	"strings"
)

// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/
// 给定一个字符串数组，把「字母异位词」分到同一组。
// 字母异位词：由相同字母、相同数量、不同顺序组成的字符串，例如 "eat" 和 "tea"。
//
// 核心思路：互为异位词的字符串，经过某种「规范化」后会得到相同的 key。
// 只要用这个 key 做哈希分组即可。常见的规范化方式有两种：
//   - 把字符排序后的字符串当 key（通用，任意字符集都可用）；
//   - 统计 26 个小写字母出现次数拼成的计数串当 key（只适用小写字母，但更快）。

// groupAnagrams 【解法1】排序作为 key 的哈希分组 ⭐ 最优（通用）
// 对每个字符串把字符排序，得到的新串作为同组标识塞进 map。
// 时间 O(n*k*logk)，n 为字符串个数、k 为字符串平均长度（排序开销）。
// 空间 O(n*k)。
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))
	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		key := string(b)
		groups[key] = append(groups[key], s)
	}
	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}

// groupAnagramsCount 【解法2】计数作为 key 的哈希分组
// 题目保证只含小写字母，统计 26 个字母出现次数拼成的串作为 key，
// 省去排序，分组更快，但只适用于固定小字符集。
// 时间 O(n*k)，空间 O(n*k)。
func groupAnagramsCount(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))
	for _, s := range strs {
		var cnt [26]int
		for i := 0; i < len(s); i++ {
			cnt[s[i]-'a']++
		}
		// 把计数数组转成分隔的字符串，避免 "1,11" 与 "11,1" 冲突。
		var sb strings.Builder
		for _, c := range cnt {
			sb.WriteByte('#')
			sb.WriteString(itoa(c))
		}
		key := sb.String()
		groups[key] = append(groups[key], s)
	}
	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}

// itoa 把非负整数转成十进制字符串，内部小工具，避免引入额外依赖语义。
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}

// 让 sort 包保持被使用（解法注释中提及排序思路，这里给出等价的稳定实现入口）。
// groupAnagramsSortStable 与解法1等价，仅用 sort.Slice 展示另一种排序写法。
func groupAnagramsSortStable(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], s)
	}
	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}
