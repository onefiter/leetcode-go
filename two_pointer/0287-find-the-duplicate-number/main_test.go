package findduplicate

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"findDuplicate":             findDuplicate,
	"findDuplicateBinarySearch": findDuplicateBinarySearch,
	"findDuplicateEnumeration":  findDuplicateEnumeration,
}

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{1, 3, 4, 2, 2}, 2},
		{"示例2", []int{3, 1, 3, 4, 2}, 3},
		{"示例3-全相同", []int{3, 3, 3, 3, 3}, 3},
		{"最小规模", []int{1, 1}, 1},
		{"重复数为1", []int{1, 1, 2}, 1},
		{"重复数为n", []int{4, 1, 2, 3, 4}, 4},
		{"重复多次", []int{2, 2, 2, 2, 1}, 2},
		{"重复在开头", []int{2, 2, 1, 3}, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				nums := slices.Clone(tt.nums)
				if got := fn(nums); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
				// 题目要求不能修改数组。
				if !slices.Equal(nums, tt.nums) {
					t.Errorf("%s 修改了输入数组: %v -> %v", name, tt.nums, nums)
				}
			})
		}
	}
}
