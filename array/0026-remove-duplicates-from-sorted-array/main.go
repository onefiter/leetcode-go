package removeduplicates

// removeDuplicates 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// 数组已升序排序，相同元素必相邻，用快慢指针原地去重。

// removeDuplicates 【解法一】快慢指针 ⭐
// slow 指向已去重部分的末尾，fast 向后扫描；
// 当 nums[fast] 与 nums[slow] 不同，说明遇到新元素，写入 slow+1。
// 时间 O(n)、空间 O(1)。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
