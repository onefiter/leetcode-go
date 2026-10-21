package nextgreater2

// 503. 下一个更大元素 II
// https://leetcode.cn/problems/next-greater-element-ii/
// 给定一个循环数组 nums（最后一个元素的下一个是第一个元素），
// 返回每个元素的下一个更大元素；不存在则为 -1。例如 [1,2,1] -> [2,-1,2]。
//
// 核心思路：循环数组的常用技巧是「把数组拼接一遍」，即下标遍历 0..2n-1 并对 n 取模，
// 这样每个元素都能看到它右边绕一圈的所有元素。再用单调递减栈求「下一个更大」即可。

// nextGreaterElements 解法一：单调栈 + 遍历两遍 ⭐ 最优方案
// 栈中存下标（对应值单调递减）。遍历 i in [0, 2n)，x = nums[i%n]，
// 当 x 大于栈顶对应的值时，栈顶的答案就是 x。只在第一遍压栈，避免重复。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	stack := make([]int, 0, n)
	for i := 0; i < 2*n; i++ {
		x := nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < x {
			res[stack[len(stack)-1]] = x
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return res
}

// nextGreaterElementsEnumeration 解法二：暴力绕一圈（对照用）
// 对每个 i，依次检查 i+1, i+2, ..., i+n-1（对 n 取模），遇到第一个更大的即为答案。
// 时间复杂度 O(n^2)，空间复杂度 O(1)（不计结果）。
func nextGreaterElementsEnumeration(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range nums {
		res[i] = -1
		for k := 1; k < n; k++ {
			if v := nums[(i+k)%n]; v > nums[i] {
				res[i] = v
				break
			}
		}
	}
	return res
}
