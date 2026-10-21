package trapwater

// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度，计算下雨之后能接多少雨水。
// 例如 height=[0,1,0,2,1,0,1,3,2,1,2,1] -> 6。
//
// 一句话理解题意：每个位置能存的水 = min(左边最高, 右边最高) - 自己的高度。
// 所有解法都是在想办法高效地求出「左边最高」和「右边最高」。

// trap 解法一：双指针 ⭐ 最优方案
// 左右指针向中间走，同时维护 leftMax、rightMax。
// 如果 leftMax < rightMax，说明左指针处的水位由 leftMax 决定（右边一定有更高的挡板），
// 直接结算左指针并右移；反之结算右指针并左移。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	res := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

// trapMonotonicStack 解法二：单调栈
// 栈中存下标，对应高度单调递减。遇到比栈顶高的柱子时，栈顶就是「坑底」，
// 新栈顶是左墙，当前柱子是右墙，按「横向一层」结算：宽 × (min(左墙,右墙) - 坑底)。
// 时间复杂度 O(n)：每个下标最多进出栈一次；空间复杂度 O(n)。
func trapMonotonicStack(height []int) int {
	res := 0
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break // 没有左墙，存不住水
			}
			l := stack[len(stack)-1]
			w := i - l - 1
			res += w * (min(height[l], h) - height[bottom])
		}
		stack = append(stack, i)
	}
	return res
}

// trapPrefixMax 解法三：前后缀最大值数组
// 先预处理 leftMax[i]（含 i 的左侧最高）和 rightMax[i]（含 i 的右侧最高），
// 再逐个位置累加 min(leftMax[i], rightMax[i]) - height[i]。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func trapPrefixMax(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	res := 0
	for i := range n {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

// trapEnumeration 解法四：暴力枚举（对照用）
// 对每个位置分别向左、向右扫描找最高柱子，再按公式结算。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func trapEnumeration(height []int) int {
	res := 0
	for i := range height {
		lm, rm := 0, 0
		for j := 0; j <= i; j++ {
			lm = max(lm, height[j])
		}
		for j := i; j < len(height); j++ {
			rm = max(rm, height[j])
		}
		res += min(lm, rm) - height[i]
	}
	return res
}
