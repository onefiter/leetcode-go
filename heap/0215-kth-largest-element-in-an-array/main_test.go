package kthlargest

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"findKthLargest":     findKthLargest,
	"findKthLargestHeap": findKthLargestHeap,
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"示例2-有重复", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"单元素", []int{1}, 1, 1},
		{"全部相同", []int{7, 7, 7, 7}, 3, 7},
		{"取最小值", []int{5, -1, 3, 9}, 4, -1},
		{"含负数取最大", []int{-5, -2, -9, -1}, 1, -1},
		{"已升序", []int{1, 2, 3, 4, 5, 6, 7, 8}, 3, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums), tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
