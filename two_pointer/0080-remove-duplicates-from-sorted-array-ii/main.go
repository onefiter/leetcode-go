package removedupsii

// 80. 删除有序数组中的重复项 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
// 给定有序数组 nums，原地删除重复元素，使每个元素最多出现两次，返回新长度 k，
// 前 k 个元素即为结果。例如 [1,1,1,2,2,3] -> 5，前 5 个为 [1,1,2,2,3]。
//
// 核心思路：快慢指针。slow 是下一个要写入的位置，fast 扫描每个元素。
// 因为数组有序，只要 nums[fast] != nums[slow-2]，说明这个值在结果里还不足两个，可以写入。

// removeDuplicates 解法一：快慢指针 ⭐ 最优方案
// 前两个元素直接保留；之后比较 nums[fast] 和已写入部分的倒数第二个 nums[slow-2]。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// removeDuplicatesEnumeration 解法二：计数 + 额外数组（对照用）
// 记录当前值连续出现的次数，次数 <= 2 时收集到新切片，最后拷回 nums 前缀。
// 思路直白但需要额外空间。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func removeDuplicatesEnumeration(nums []int) int {
	res := make([]int, 0, len(nums))
	cnt := 0
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			res = append(res, x)
		}
	}
	copy(nums, res)
	return len(res)
}
