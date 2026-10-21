package threesum

import "sort"

// 15. 三数之和
// https://leetcode.cn/problems/3sum/
// 给定一个整数数组 nums，找出所有「和为 0」且「不重复」的三元组 [a, b, c]。
// 答案中不可以包含重复的三元组（三元组本身内部顺序不限，但集合要去重）。
// 例如 nums=[-1,0,1,2,-1,-4] -> [[-1,-1,2],[-1,0,1]]。
//
// 一句话理解题意：固定一个数 a，问题就退化为在剩余元素里找「两数之和 = -a」。
//
// 核心思路：先排序，然后固定第一个数 nums[i]，
// 在 [i+1, n-1] 区间内用双指针找两数之和为 -nums[i]。
// 排序带来两大好处：
//   - 双指针能根据和的大小决定左右收缩方向；
//   - 相同元素相邻，方便跳过重复，天然完成去重。

// threeSum 解法一：排序 + 双指针 ⭐ 最优方案
// 排序后固定 i，l、r 在右侧区间向中间逼近：
//   - 三数和 == 0：记录答案，同时跳过与当前 l、r 相同的值去重；
//   - 三数和 <  0：l 右移增大；
//   - 三数和 >  0：r 左移减小。
//
// 外层跳过与上一个相同的 nums[i]，避免产生重复三元组。
// 时间复杂度 O(n^2)：排序 O(n log n)，双指针嵌套 O(n^2)。
// 空间复杂度 O(1)（不计返回值与排序栈）。
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		// 最小的数都 > 0，三数之和不可能为 0，提前结束。
		if nums[i] > 0 {
			break
		}
		// 跳过重复的第一个数。
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 跳过重复的 l 与 r。
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			case sum < 0:
				l++
			default:
				r--
			}
		}
	}
	return res
}

// threeSumEnumeration 解法二：三重循环穷举 + set 去重（对照用）
// 不依赖有序性，枚举所有下标三元组，和为 0 则把「排序后的三元组」作为
// map 的 key 去重，最后收集所有不同的三元组。
// 时间复杂度 O(n^3)，空间复杂度 O(n)（去重 set）。
func threeSumEnumeration(nums []int) [][]int {
	n := len(nums)
	seen := map[[3]int]struct{}{}
	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] != 0 {
					continue
				}
				triple := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triple[:])
				if _, ok := seen[triple]; ok {
					continue
				}
				seen[triple] = struct{}{}
				res = append(res, []int{triple[0], triple[1], triple[2]})
			}
		}
	}
	return res
}
