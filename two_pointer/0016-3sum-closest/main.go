package threesumclosest

import "sort"

// 16. 最接近的三数之和
// https://leetcode.cn/problems/3sum-closest/
// 给定一个整数数组 nums 和目标值 target，找出三个整数，
// 使它们的和与 target 最接近，返回这三个数的和。假定每组输入恰有一个答案。
// 例如 nums=[-1,2,1,-4], target=1 -> 2（即 -1+2+1=2）。
//
// 一句话理解题意：和 15 题几乎一样，只是把「等于 0」换成「尽量贴近 target」。
//
// 核心思路：先排序，固定第一个数 nums[i]，右侧用双指针。
// 维护一个当前最接近的和 best，每次用 |sum-target| 与 |best-target| 比较更新；
// 根据 sum 与 target 的大小关系收缩左右指针（sum 偏小则 l++，偏大则 r--）。

// threeSumClosest 解法一：排序 + 双指针 ⭐ 最优方案
// best 初始化为前三个数之和。固定 i 后，l、r 向中间逼近：
//   - sum == target：已是最接近，直接返回；
//   - sum <  target：l 右移，尝试增大；
//   - sum >  target：r 左移，尝试减小。
//
// 过程中不断用更接近 target 的 sum 更新 best。
// 时间复杂度 O(n^2)：排序 O(n log n)，双指针嵌套 O(n^2)。
// 空间复杂度 O(1)。
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < abs(best-target) {
				best = sum
			}
			switch {
			case sum == target:
				return sum // 完全命中，不可能更接近
			case sum < target:
				l++
			default:
				r--
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// threeSumClosestEnumeration 解法二：三重循环穷举（对照用）
// 不依赖有序性，枚举所有三元组，记录与 target 差距最小的那组和。
// 时间复杂度 O(n^3)，空间复杂度 O(1)。
func threeSumClosestEnumeration(nums []int, target int) int {
	n := len(nums)
	best := nums[0] + nums[1] + nums[2]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < abs(best-target) {
					best = sum
				}
			}
		}
	}
	return best
}
