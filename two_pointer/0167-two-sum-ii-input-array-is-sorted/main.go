package twosumii

// 167. 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 给定一个「非递减」排列的整数数组 numbers 和目标值 target，
// 找出两个数使它们相加之和等于 target，返回这两个数的下标（下标从 1 开始）。
// 题目保证有且仅有一组解，同一个元素不能使用两次。
// 例如 numbers=[2,7,11,15], target=9 -> [1,2]。
//
// 一句话理解题意：数组有序，利用「两头往中间逼近」的双指针即可线性求解。
//
// 核心思路：左指针 l 指向最小端、右指针 r 指向最大端。
//   - 两数之和 == target：命中，返回下标（+1 转成 1 起始）；
//   - 两数之和 <  target：说明和太小，左指针右移增大和；
//   - 两数之和 >  target：说明和太大，右指针左移减小和。

// twoSum 解法一：排序数组上的双指针 ⭐ 最优方案
// 利用数组已经有序的前提，l、r 从两端向中间收缩，
// 每一步都能排除一个不可能的元素，整体一次遍历即可。
// 时间复杂度 O(n)：每个元素最多被访问一次。
// 空间复杂度 O(1)。
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		switch {
		case sum == target:
			return []int{l + 1, r + 1}
		case sum < target:
			l++ // 和太小，增大左值
		default:
			r-- // 和太大，减小右值
		}
	}
	return nil // 题目保证有解，正常不会到这里
}

// twoSumEnumeration 解法二：双重循环穷举（对照用）
// 不依赖有序性，枚举所有下标对 (i, j)，找到和为 target 的一对。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func twoSumEnumeration(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}
