package sumcounts

import "leetcode-go/template"

const mod = 1_000_000_007

// sumCounts 2916. 子数组不同元素数目的平方和 II
// https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/
// 枚举右端点 i，设 f(j) 为子数组 nums[j..i] 的不同元素数。
// 加入 nums[i] 后，左端点 j ∈ (last[nums[i]], i] 的 f(j) 都 +1，
// 平方和增量 = Σ((f+1)^2 - f^2) = 2*Σf + 区间长度。
// 用支持区间加、区间求和的树状数组维护 f，时间 O(n log n)。
func sumCounts(nums []int) int {
	n := len(nums)
	bit := template.NewRangeBIT(n)
	last := make(map[int]int) // 值 -> 上次出现位置（1-indexed），未出现为 0
	ans, cur := 0, 0          // cur：以 i 结尾的所有子数组的平方和
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		l := last[x] + 1
		cur = (cur + 2*bit.RangeQuery(l, i) + (i - l + 1)) % mod
		bit.RangeAdd(l, i, 1)
		ans = (ans + cur) % mod
		last[x] = i
	}
	return ans
}

// sumCountsEnumeration 解法二：穷举/朴素做法（对照用，大数据会超时）
// 枚举每个左端点 i，向右扩展右端点 j，用集合统计 [i, j] 的不同元素数并累加平方。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func sumCountsEnumeration(nums []int) int {
	ans := 0
	for i := range nums {
		set := map[int]bool{}
		for j := i; j < len(nums); j++ {
			set[nums[j]] = true
			ans = (ans + len(set)*len(set)) % mod
		}
	}
	return ans
}
