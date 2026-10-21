package dailytemperatures

// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures/
// 给定每日温度数组 temperatures，返回数组 answer：answer[i] 表示第 i 天之后
// 要再等几天才会出现更高的温度；如果之后都不会升温，则为 0。
// 例如 [73,74,75,71,69,72,76,73] -> [1,1,4,2,1,1,0,0]。
//
// 核心思路：这是「下一个更大元素」问题，只不过要的是距离而不是值。
// 用单调递减栈存「还没等到升温的日子」的下标，新的一天比栈顶更热时，
// 栈顶那天就等到了，答案 = 今天下标 - 栈顶下标。

// dailyTemperatures 解法一：单调栈 ⭐ 最优方案
// 时间复杂度 O(n)：每个下标进出栈各一次。空间复杂度 O(n)。
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0, n)
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}
		stack = append(stack, i)
	}
	return ans
}

// dailyTemperaturesJump 解法二：从右往左跳跃查找
// 从右往左计算，找 i 右侧第一个更高温度时，若 j 不够高，
// 可以直接跳到 j + ans[j]（比 j 更高的第一天），跳过中间不可能的日子；
// ans[j]==0 说明 j 右边没有更高的了，也就不可能比 i 高。
// 时间复杂度 O(n)（均摊），空间复杂度 O(1)（不计结果）。
func dailyTemperaturesJump(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for temperatures[j] <= temperatures[i] && ans[j] > 0 {
			j += ans[j]
		}
		if temperatures[j] > temperatures[i] {
			ans[i] = j - i
		}
	}
	return ans
}

// dailyTemperaturesEnumeration 解法三：暴力向右找（对照用）
// 时间复杂度 O(n^2)，空间复杂度 O(1)（不计结果）。
func dailyTemperaturesEnumeration(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				ans[i] = j - i
				break
			}
		}
	}
	return ans
}
