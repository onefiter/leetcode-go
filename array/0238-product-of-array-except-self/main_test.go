package productexceptself

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) []int{
	"productExceptSelf":            productExceptSelf,
	"productExceptSelfEnumeration": productExceptSelfEnumeration,
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		// 单元素
		{[]int{5}, []int{1}},
		// 两元素
		{[]int{2, 3}, []int{3, 2}},
		// 含多个 0
		{[]int{0, 0}, []int{0, 0}},
		// 全相同
		{[]int{2, 2, 2}, []int{4, 4, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			got := fn(slices.Clone(tt.nums))
			if !slices.Equal(got, tt.want) {
				t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
			}
		}
	}
}
