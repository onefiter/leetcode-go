package twosum

// twoSum 1. 两数之和
// https://leetcode.cn/problems/two-sum/
// 哈希表记录已遍历的值 -> 下标，时间 O(n)，空间 O(n)。
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}

func twoSumEnumeration(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-x {
				return []int{i, j}
			}
		}
	}
	return nil
}
