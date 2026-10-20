package removeelement

// removeElement 27. 移除元素
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
