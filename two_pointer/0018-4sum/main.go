package foursum

import "sort"

// 18. 四数之和
// https://leetcode.cn/problems/4sum/
// 给定一个整数数组 nums 和目标值 target，找出所有「和为 target」且「不重复」
// 的四元组 [a, b, c, d]。答案中不可以包含重复的四元组。
// 例如 nums=[1,0,-1,0,-2,2], target=0 -> [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]。
//
// 一句话理解题意：在 15 题（三数之和）外面再套一层循环，固定两个数后双指针。
//
// 核心思路：先排序，两层循环固定前两个数 nums[i]、nums[j]，
// 剩下用双指针在 [j+1, n-1] 找两数之和为 target-nums[i]-nums[j]。
// 注意 target 可能为任意整数（含负数），不能像 3sum 那样靠 nums[i]>0 剪枝为 0，
// 这里用「当前可能的最小/最大和」做剪枝，并在每层跳过重复元素去重。

// fourSum 解法一：排序 + 双重固定 + 双指针 ⭐ 最优方案
// 外两层固定 i、j（都跳过重复值并做上下界剪枝），内层 l、r 双指针逼近：
//   - 四数和 == target：记录答案并跳过重复的 l、r；
//   - 四数和 <  target：l 右移；
//   - 四数和 >  target：r 左移。
//
// 时间复杂度 O(n^3)：排序 O(n log n)，三层嵌套 O(n^3)。
// 空间复杂度 O(1)（不计返回值与排序栈）。
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复的第一个数
		}
		// 当前 i 下的最小可能和都大于 target，后面更大，直接结束。
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 当前 i 配上最大的三个数都小于 target，换更大的 i。
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // 跳过重复的第二个数
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case sum == target:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				case sum < target:
					l++
				default:
					r--
				}
			}
		}
	}
	return res
}

// fourSumEnumeration 解法二：四重循环穷举 + set 去重（对照用）
// 不依赖有序性，枚举所有下标四元组，和为 target 则把「排序后的四元组」
// 作为 map 的 key 去重，最后收集所有不同的四元组。
// 时间复杂度 O(n^4)，空间复杂度 O(n)（去重 set）。
func fourSumEnumeration(nums []int, target int) [][]int {
	n := len(nums)
	seen := map[[4]int]struct{}{}
	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for m := k + 1; m < n; m++ {
					if nums[i]+nums[j]+nums[k]+nums[m] != target {
						continue
					}
					quad := [4]int{nums[i], nums[j], nums[k], nums[m]}
					sort.Ints(quad[:])
					if _, ok := seen[quad]; ok {
						continue
					}
					seen[quad] = struct{}{}
					res = append(res, []int{quad[0], quad[1], quad[2], quad[3]})
				}
			}
		}
	}
	return res
}
