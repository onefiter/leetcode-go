package rangesumimmutable

// NumArray 303. 区域和检索 - 数组不可变
// https://leetcode.cn/problems/range-sum-query-immutable/
// 数组初始化后不变，需多次查询区间 [left, right] 的和。

// NumArray 【解法一】前缀和设计 ⭐
// 预处理前缀和 prefix，其中 prefix[i] = nums[0..i-1] 之和，prefix[0] = 0。
// 则区间和 sum(left, right) = prefix[right+1] - prefix[left]。
// 构造时间 O(n)；每次查询时间 O(1)、空间 O(n)。
type NumArray struct {
	prefix []int
}

// Constructor 预处理前缀和数组。
func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + x
	}
	return NumArray{prefix: prefix}
}

// SumRange 返回闭区间 [left, right] 的元素和。
func (na *NumArray) SumRange(left, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}
