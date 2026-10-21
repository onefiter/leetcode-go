package containerwater

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/
// 给定 n 个非负整数 height，每个数代表坐标上一条垂直线的高度。
// 找出两条线，使它们与 x 轴围成的容器能容纳最多的水，返回最大容量。
// 容量 = 两线间距（下标差） * 较短的那条线的高度。
// 例如 height=[1,8,6,2,5,4,8,3,7] -> 49。
//
// 一句话理解题意：在所有「左右两条线」的组合里，找面积最大的一组。
//
// 核心思路：左右双指针分别指向最左、最右两条线。
// 面积受限于较短的那条线，移动较高的线只会让宽度变小、面积不增；
// 因此每次移动「较短」的那条线，才有可能遇到更高的线、得到更大面积。

// maxArea 解法一：左右双指针 ⭐ 最优方案
// l、r 从两端开始，每轮用 min(height[l],height[r]) * (r-l) 更新最大面积，
// 然后移动较短的一侧（相等时移动任意一侧均可，这里移动左侧）。
// 时间复杂度 O(n)：双指针一次遍历。
// 空间复杂度 O(1)。
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	best := 0
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		if area := h * (r - l); area > best {
			best = area
		}
		if height[l] < height[r] {
			l++ // 左侧较矮，移动左侧寻找更高线
		} else {
			r-- // 右侧较矮或相等，移动右侧
		}
	}
	return best
}

// maxAreaEnumeration 解法二：双重循环穷举（对照用）
// 枚举所有左右线的组合 (i, j)，逐一计算面积取最大值。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func maxAreaEnumeration(height []int) int {
	n := len(height)
	best := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			if area := h * (j - i); area > best {
				best = area
			}
		}
	}
	return best
}
