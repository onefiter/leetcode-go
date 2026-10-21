package sumsubarraymins

// 907. 子数组的最小值之和
// https://leetcode.cn/problems/sum-of-subarray-minimums/
// 给定整数数组 arr，求所有连续子数组的最小值之和，结果对 1e9+7 取模。
// 例如 [3,1,2,4] -> 17。
//
// 核心思路（贡献法）：与其枚举每个子数组再找最小值，不如反过来问：
// arr[i] 是多少个子数组的最小值？设 left[i] 为左侧第一个「< arr[i]」的下标，
// right[i] 为右侧第一个「<= arr[i]」的下标，则以 arr[i] 为最小值的子数组个数为
// (i - left[i]) * (right[i] - i)，贡献 = arr[i] * 个数。
// 一侧严格、一侧非严格，是为了在有相同值时每个子数组只被统计一次。

const mod = 1_000_000_007

// sumSubarrayMins 解法一：单调栈 + 贡献法（一次遍历）⭐ 最优方案
// 维护单调递增栈，当 arr[i] <= 栈顶值时弹出栈顶 top：
// 此时 top 的右边界就是 i，左边界是弹出后的新栈顶（栈空则为 -1），直接结算 top 的贡献。
// 末尾用 i == n（视为负无穷）作为哨兵，结算剩余元素。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func sumSubarrayMins(arr []int) int {
	n := len(arr)
	res := 0
	stack := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		for len(stack) > 0 && (i == n || arr[i] <= arr[stack[len(stack)-1]]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			res = (res + arr[top]*(top-left)%mod*(i-top)) % mod
		}
		stack = append(stack, i)
	}
	return res
}

// sumSubarrayMinsTwoPass 解法二：单调栈分别求左右边界
// 先从左到右求 left[i]（左侧第一个 < arr[i]），再从右到左求 right[i]（右侧第一个 <= arr[i]），
// 最后累加 arr[i] * (i-left[i]) * (right[i]-i)。和解法一等价，边界含义更直观。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func sumSubarrayMinsTwoPass(arr []int) int {
	n := len(arr)
	left, right := make([]int, n), make([]int, n)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		left[i] = -1
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		right[i] = n
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + arr[i]*(i-left[i])%mod*(right[i]-i)) % mod
	}
	return res
}

// sumSubarrayMinsEnumeration 解法三：枚举左端点暴力（对照用）
// 固定左端点 i，向右扩展时维护当前最小值并累加。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func sumSubarrayMinsEnumeration(arr []int) int {
	res := 0
	for i := range arr {
		m := arr[i]
		for j := i; j < len(arr); j++ {
			m = min(m, arr[j])
			res = (res + m) % mod
		}
	}
	return res
}
