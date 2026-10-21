package removedupsii

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"removeDuplicates":            removeDuplicates,
	"removeDuplicatesEnumeration": removeDuplicatesEnumeration,
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"示例1", []int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}},
		{"示例2", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}},
		{"空数组", []int{}, []int{}},
		{"单个元素", []int{7}, []int{7}},
		{"两个相同", []int{2, 2}, []int{2, 2}},
		{"全部相同", []int{5, 5, 5, 5, 5}, []int{5, 5}},
		{"无重复", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"含负数", []int{-3, -3, -3, -1, 0, 0, 0}, []int{-3, -3, -1, 0, 0}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				nums := slices.Clone(tt.nums) // 原地修改，避免污染用例
				k := fn(nums)
				if k != len(tt.want) || !slices.Equal(nums[:k], tt.want) {
					t.Errorf("%s(%v) = %d, %v, want %d, %v", name, tt.nums, k, nums[:k], len(tt.want), tt.want)
				}
			})
		}
	}
}
