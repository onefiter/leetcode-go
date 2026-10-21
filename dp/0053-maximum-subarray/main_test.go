package maxsubarray

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"maxSubArray":       maxSubArray,
	"maxSubArrayDivide": maxSubArrayDivide,
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"单元素", []int{1}, 1},
		{"示例3", []int{5, 4, -1, 7, 8}, 23},
		{"全负数取最大单个", []int{-3, -1, -2}, -1},
		{"全正数取整体", []int{1, 2, 3}, 6},
		{"负数被跨过", []int{3, -1, 3}, 5},
		{"负数切断", []int{2, -5, 3}, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
