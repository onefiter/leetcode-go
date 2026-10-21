package subarraysumk

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/
// 给定一个整数数组 nums 和整数 k，统计并返回该数组中和为 k 的「连续子数组」的个数。
// 注意数组可能含负数，所以不能用滑动窗口（和不是单调的）。
//
// 核心思路：前缀和 + 哈希表。
//   记 prefix[i] 为前 i 个元素之和，则子数组 (j, i] 的和 = prefix[i] - prefix[j]。
//   要使其等于 k，即 prefix[j] = prefix[i] - k。
//   一边遍历累加前缀和，一边用哈希表记录每个前缀和出现的次数，
//   对当前前缀和 s，答案累加 map[s-k]（之前有多少个前缀和等于 s-k）。

// subarraySum 解法一：前缀和 + 哈希表 ⭐ 最优方案
// 哈希表 key 为前缀和，value 为该前缀和出现的次数；初始放入 {0:1}
// 表示「空前缀」的和为 0，用于处理从下标 0 开始就等于 k 的子数组。
// 时间复杂度 O(n)：一次遍历。
// 空间复杂度 O(n)：哈希表最多存 n+1 个前缀和。
func subarraySum(nums []int, k int) int {
	prefixCount := map[int]int{0: 1}
	sum, count := 0, 0
	for _, x := range nums {
		sum += x
		count += prefixCount[sum-k]
		prefixCount[sum]++
	}
	return count
}

// subarraySumBrute 解法二：双重循环枚举所有子数组（对照用）
// 外层固定左端点，内层向右扩展并累加和，等于 k 就计数。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func subarraySumBrute(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
