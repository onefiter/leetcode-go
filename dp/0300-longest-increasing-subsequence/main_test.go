package lis

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"lengthOfLIS":   lengthOfLIS,
	"lengthOfLISDP": lengthOfLISDP,
}

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"示例2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"示例3-全相同", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"单元素", []int{5}, 1},
		{"严格递减", []int{5, 4, 3, 2, 1}, 1},
		{"严格递增", []int{1, 2, 3, 4, 5}, 5},
		{"含负数", []int{-2, -1, -3, 0, -5, 1}, 4},
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
