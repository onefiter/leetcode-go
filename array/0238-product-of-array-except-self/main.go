package productexceptself

// productExceptSelf 238. 除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/
// answer[i] 等于除 nums[i] 外所有元素的乘积，要求不使用除法且 O(n) 时间。

// productExceptSelf 【解法一】前缀积 + 后缀积 ⭐
// 先用一次正向遍历把 answer[i] 填成 i 左侧所有元素之积（前缀积）；
// 再用一个变量 right 从右向左累乘后缀积，乘到 answer[i] 上。
// 这样 answer[i] = 左积 * 右积，即除自身外的乘积。
// 时间 O(n)、空间 O(1)（输出数组不计）。
func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	// 前缀积：answer[i] = nums[0..i-1] 的乘积
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	// 后缀积：right 累乘 nums[i+1..n-1]
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}

// productExceptSelfEnumeration 【解法二】后缀暴力枚举（对照用）
// 对每个下标 i，内层从 i 的后一个位置起向后再乘一遍，并配合外层累积的前缀，
// 直接双重循环求出除自身外的乘积。逻辑最直白，但效率低。
// 时间 O(n^2)、空间 O(1)（输出数组不计）。
func productExceptSelfEnumeration(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		prod := 1
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			prod *= nums[j]
		}
		answer[i] = prod
	}
	return answer
}
