package longestconsecutive

import (
	"slices"
	"sort"
)

// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/
// 给定一个未排序的整数数组 nums，找出数字连续的最长序列（不要求在原数组中连续）的长度。
// 要求时间复杂度 O(n)。
// 例如 nums=[100,4,200,1,3,2] -> 连续序列 [1,2,3,4]，长度 4。
//
// 核心思路：先把所有数去重放进哈希集合。对每个数 x，只有当 x-1 不在集合里时，
// x 才可能是某条连续序列的「起点」，从起点往后 x+1, x+2... 计数。
// 因为每个数只会在它是起点时被向后扩展一次，整体是 O(n)。

// longestConsecutive 【解法1】哈希集合 + 只从序列起点扩展 ⭐ 最优
// 时间 O(n)：每个元素最多被访问两次（一次判断起点，一次被扩展）。
// 空间 O(n)：哈希集合。
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		set[x] = struct{}{}
	}
	best := 0
	for x := range set {
		// x 不是起点就跳过，保证每条序列只从最小值开始扩展。
		if _, ok := set[x-1]; ok {
			continue
		}
		cur := x
		length := 1
		for {
			if _, ok := set[cur+1]; !ok {
				break
			}
			cur++
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}

// longestConsecutiveSort 【解法2】排序后线性扫描（对照用）
// 排序后相邻元素要么相等（跳过）、要么差 1（延长当前序列）、要么断开（重置）。
// 时间 O(n*logn)：排序开销，不满足 O(n) 要求，作为正确性对照。
// 空间 O(n)：克隆一份避免修改入参（或视排序实现为 O(logn)）。
func longestConsecutiveSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	a := slices.Clone(nums)
	sort.Ints(a)
	best, cur := 1, 1
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			continue // 重复元素不影响连续长度
		}
		if a[i] == a[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
