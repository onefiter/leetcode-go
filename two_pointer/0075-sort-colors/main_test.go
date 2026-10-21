package sortcolors

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int){
	"sortColors":            sortColors,
	"sortColorsEnumeration": sortColorsEnumeration,
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"示例1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"示例2", []int{2, 0, 1}, []int{0, 1, 2}},
		{"空数组", []int{}, []int{}},
		{"单个元素", []int{1}, []int{1}},
		{"全是0", []int{0, 0, 0}, []int{0, 0, 0}},
		{"全是2", []int{2, 2, 2}, []int{2, 2, 2}},
		{"已排序", []int{0, 1, 2}, []int{0, 1, 2}},
		{"逆序", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"只有0和2", []int{2, 0, 2, 0}, []int{0, 0, 2, 2}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				nums := slices.Clone(tt.nums) // 原地修改，避免污染用例
				fn(nums)
				if !slices.Equal(nums, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, nums, tt.want)
				}
			})
		}
	}
}
