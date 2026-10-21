package movezeroes

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{4, 0, 5, 0, 0, 6}, []int{4, 5, 6, 0, 0, 0}},
	}
	for _, tt := range tests {
		nums := slices.Clone(tt.nums)
		moveZeroes(nums)
		if !slices.Equal(nums, tt.want) {
			t.Errorf("moveZeroes(%v) = %v, want %v", tt.nums, nums, tt.want)
		}
	}
}
