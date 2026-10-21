package shortestsubarray

// 862. 和至少为 K 的最短子数组
// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
// 给定整数数组 nums（可能含负数）和整数 k，找出和 >= k 的最短非空子数组长度，不存在返回 -1。
// 例如 nums=[2,-1,2], k=3 -> 3。
//
// 核心思路：前缀和 pre[i] = nums[0..i-1] 之和，子数组 [i, j) 的和 = pre[j]-pre[i]。
// 对每个右端点 j，想找最靠右的 i 使 pre[i] <= pre[j]-k。
// 因为有负数，普通滑动窗口不适用，改用「前缀和单调递增的双端队列」（存下标）：
//   - 队头：若 pre[j]-pre[队头] >= k，队头找到了它最短的答案，记录后弹出；
//   - 队尾：若 pre[队尾] >= pre[j]，队尾永远不如 j（j 更靠右且前缀和更小），弹出。

// shortestSubarray 解法一：前缀和 + 单调双端队列 ⭐ 最优方案
// 每个下标最多入队、出队各一次。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}
	ans := n + 1
	q := make([]int, 0, n+1) // 存下标，对应 pre 值单调递增
	for j := 0; j <= n; j++ {
		// 队头满足条件：记录长度后弹出（之后的 j 只会更长，不需要它了）。
		for len(q) > 0 && pre[j]-pre[q[0]] >= k {
			ans = min(ans, j-q[0])
			q = q[1:]
		}
		// 维护单调性：队尾前缀和 >= 当前前缀和的都没用了。
		for len(q) > 0 && pre[q[len(q)-1]] >= pre[j] {
			q = q[:len(q)-1]
		}
		q = append(q, j)
	}
	if ans == n+1 {
		return -1
	}
	return ans
}

// shortestSubarrayEnumeration 解法二：前缀和 + 枚举所有子数组（对照用）
// 枚举左端点 i，从短到长枚举右端点 j，用前缀和 O(1) 求区间和，第一个满足的就是该左端点的最短答案。
// 时间复杂度 O(n²)，空间复杂度 O(n)。
func shortestSubarrayEnumeration(nums []int, k int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}
	ans := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if pre[j]-pre[i] >= k {
				if ans == -1 || j-i < ans {
					ans = j - i
				}
				break
			}
		}
	}
	return ans
}
