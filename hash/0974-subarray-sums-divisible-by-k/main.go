package subarraydivk

// 974. 和可被 K 整除的子数组
// https://leetcode.cn/problems/subarray-sums-divisible-by-k/
// 给定整数数组 nums 和整数 k，返回其中和可以被 k 整除的「连续子数组」的数目。
//
// 核心思路：前缀和 + 同余 + 哈希表。
//   子数组 (j, i] 的和 = prefix[i] - prefix[j]，它能被 k 整除
//   当且仅当 prefix[i] 与 prefix[j] 对 k 「同余」（除以 k 余数相同）。
//   于是只需统计每个余数出现的次数，对每个新前缀和的余数 r，
//   把之前出现过 r 的次数累加进答案即可。
//   注意 Go 的 % 对负数会得到负余数，要用 ((s%k)+k)%k 归一化到 [0, k)。

// subarraysDivByK 解法一：前缀和 + 同余 + 哈希表 ⭐ 最优方案
// remCount[r] 记录「余数为 r 的前缀和」出现次数，初始 {0:1} 表示空前缀。
// 时间复杂度 O(n)：一次遍历。
// 空间复杂度 O(k)：哈希表最多存 k 个不同余数。
func subarraysDivByK(nums []int, k int) int {
	remCount := map[int]int{0: 1}
	sum, count := 0, 0
	for _, x := range nums {
		sum += x
		r := ((sum % k) + k) % k // 归一化到 [0, k)
		count += remCount[r]
		remCount[r]++
	}
	return count
}

// subarraysDivByKBrute 解法二：双重循环枚举所有子数组（对照用）
// 外层固定左端点，内层向右扩展累加和，和能被 k 整除就计数。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func subarraysDivByKBrute(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum%k == 0 {
				count++
			}
		}
	}
	return count
}
