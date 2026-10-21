package partition

// 416. 分割等和子集
// https://leetcode.cn/problems/partition-equal-subset-sum/
// 给一个只包含正整数的非空数组 nums，判断能否将其分割成两个子集，使两个子集的元素和相等。
//
// 核心思路：转化为 01 背包可行性问题。总和 sum 为奇数直接 false；
// 否则问能否从 nums 中选出若干个数恰好凑成 target = sum/2。
// 物品是每个数字（只能选一次），容量是 target。
// dp[j] 表示能否恰好凑出和 j，dp[0] = true，转移：dp[j] = dp[j] || dp[j-num]。
// 01 背包内层容量必须「倒序」遍历，保证 dp[j-num] 还是上一轮（未用当前物品）的值。

// canPartition 解法一：01 背包 + 一维滚动数组 ⭐ 最优方案
// 时间复杂度 O(n * sum)，空间复杂度 O(sum)。
func canPartition(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- { // 倒序：每个数只用一次
			dp[j] = dp[j] || dp[j-num]
		}
		if dp[target] {
			return true // 提前结束
		}
	}
	return dp[target]
}

// canPartition2D 解法二：01 背包 + 二维 dp 表（对照用）
// dp[i][j] 表示前 i 个数能否恰好凑出 j：
// dp[i][j] = dp[i-1][j]（不选）|| dp[i-1][j-nums[i-1]]（选，需 j >= nums[i-1]）。
// 时间复杂度 O(n * sum)，空间复杂度 O(n * sum)。
func canPartition2D(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	n := len(nums)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= num {
				dp[i][j] = dp[i][j] || dp[i-1][j-num]
			}
		}
	}
	return dp[n][target]
}
