package removeelement

// removeElement 27. 移除元素
// https://leetcode.cn/problems/remove-element/
// 快慢指针：slow 指向下一个保留位置，时间 O(n)，空间 O(1)。
func removeElement(nums []int, val int) int {
	slow := 0
	for _, x := range nums {
		if x != val {
			nums[slow] = x
			slow++
		}
	}
	return slow
}

// removeElementEnumeration 解法二：穷举/朴素做法（对照用）
// 把不等于 val 的元素依次收集到新切片，再拷回原数组。
// 逻辑最直白，但多用了一个辅助数组。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func removeElementEnumeration(nums []int, val int) int {
	kept := make([]int, 0, len(nums))
	for _, x := range nums {
		if x != val {
			kept = append(kept, x)
		}
	}
	copy(nums, kept)
	return len(kept)
}
