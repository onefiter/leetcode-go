package movezeroes

// moveZeroes 283. 移动零
// https://leetcode.cn/problems/move-zeroes/
// 把所有 0 移到末尾，同时保持非零元素的相对顺序，必须原地操作。

// moveZeroes 【解法一】快慢指针 ⭐
// slow 指向下一个应放非零元素的位置；fast 扫描，遇到非零就与 slow 交换。
// 交换保证非零元素前移且顺序不变，0 自然被挤到后面。
// 时间 O(n)、空间 O(1)。
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
