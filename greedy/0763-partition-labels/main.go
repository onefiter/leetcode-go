package partitionlabels

// 763. 划分字母区间
// https://leetcode.cn/problems/partition-labels/
// 给一个只含小写字母的字符串 s，把它划分为尽可能多的片段，同一字母最多出现在一个片段中，
// 返回每个片段长度组成的列表。
//
// 核心思路：先记录每个字母最后一次出现的位置 last[c]。从左往右扫描，维护当前片段必须延伸到的最远位置 end，
// 当 i == end 时说明片段内所有字母都不会在后面出现，立刻切一刀。
// 正确性：片段至少要覆盖到其中所有字母的最后位置，end 正是满足约束的最小右端点，
// 在最早的合法位置切分，剩余部分能切出的片段数只会更多。

// partitionLabels 解法一：记录最后出现位置 + 一次扫描贪心 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(|Σ|)=O(26)。
func partitionLabels(s string) []int {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	res := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

// partitionLabelsMerge 解法二：区间合并（对照用）
// 每个字母对应区间 [first, last]，按 first 顺序合并重叠区间，合并后每个区间就是一个片段。
// 时间复杂度 O(n)（字母按首次出现顺序天然有序，最多 26 个区间），空间复杂度 O(26)。
func partitionLabelsMerge(s string) []int {
	var first, last [26]int
	for i := range first {
		first[i] = -1
	}
	order := []int{} // 按首次出现顺序记录字母，区间起点天然有序
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if first[c] == -1 {
			first[c] = i
			order = append(order, int(c))
		}
		last[c] = i
	}
	res := []int{}
	curL, curR := -1, -1
	for _, c := range order {
		l, r := first[c], last[c]
		if curR == -1 {
			curL, curR = l, r
		} else if l <= curR {
			curR = max(curR, r) // 有重叠，合并
		} else {
			res = append(res, curR-curL+1)
			curL, curR = l, r
		}
	}
	if curR != -1 {
		res = append(res, curR-curL+1)
	}
	return res
}
