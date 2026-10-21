package largestsumafterk

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"largestSumAfterKNegations":     largestSumAfterKNegations,
	"largestSumAfterKNegationsHeap": largestSumAfterKNegationsHeap,
}

func TestLargestSumAfterKNegations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{4, 2, 3}, 1, 5},
		{"示例2", []int{3, -1, 0, 2}, 3, 6},
		{"示例3", []int{2, -3, -1, 5, -4}, 2, 13},
		{"负数翻完剩奇数次", []int{-2, -3, 1}, 3, 4},
		{"负数翻完剩偶数次", []int{-2, -3, 1}, 4, 6},
		{"k 小于负数个数", []int{-8, -5, -1}, 2, 12},
		{"单元素", []int{-5}, 3, 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.nums), tt.k)
				if got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
