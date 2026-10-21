package removeduplicates

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int // 去重后保留的前缀
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		nums := slices.Clone(tt.nums)
		k := removeDuplicates(nums)
		got := slices.Clone(nums[:k])
		if !slices.Equal(got, tt.want) {
			t.Errorf("removeDuplicates(%v) = %v (k=%d), want %v", tt.nums, got, k, tt.want)
		}
	}
}
