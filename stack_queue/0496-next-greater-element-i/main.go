package nextgreater

// 496. 下一个更大元素 I
// https://leetcode.cn/problems/next-greater-element-i/
// nums1 是 nums2 的子集，且元素各不相同。对 nums1 中每个 x，
// 找出它在 nums2 中的位置右侧第一个比 x 大的元素，不存在则为 -1。
//
// 核心思路：先对 nums2 求出「每个元素的下一个更大元素」，存进哈希表，
// 再对 nums1 逐个查表。求「下一个更大」用单调递减栈：
// 新元素比栈顶大时，栈顶的答案就是它。

// nextGreaterElement 解法一：单调栈 + 哈希表 ⭐ 最优方案
// 时间复杂度 O(m + n)，空间复杂度 O(n)。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := make(map[int]int, len(nums2))
	stack := []int{}
	for _, x := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < x {
			next[stack[len(stack)-1]] = x
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, x)
	}
	res := make([]int, len(nums1))
	for i, x := range nums1 {
		if v, ok := next[x]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}

// nextGreaterElementEnumeration 解法二：暴力查找（对照用）
// 对 nums1 的每个 x，先在 nums2 中找到它的位置，再往右找第一个更大的数。
// 时间复杂度 O(m · n)，空间复杂度 O(1)（不计结果）。
func nextGreaterElementEnumeration(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i, x := range nums1 {
		res[i] = -1
		j := 0
		for nums2[j] != x {
			j++
		}
		for j++; j < len(nums2); j++ {
			if nums2[j] > x {
				res[i] = nums2[j]
				break
			}
		}
	}
	return res
}
