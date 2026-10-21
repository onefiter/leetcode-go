package maxproduct

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"maxProduct":      maxProduct,
	"maxProductBrute": maxProductBrute,
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{2, 3, -2, 4}, 6},
		{"示例2零隔断", []int{-2, 0, -1}, 0},
		{"两个负数相乘", []int{-2, 3, -4}, 24},
		{"单个负数", []int{-2}, -2},
		{"零后正数", []int{0, 2}, 2},
		{"三个负数取后两个", []int{-1, -2, -3}, 6},
		{"负数翻转最小值", []int{2, -5, -2, -4, 3}, 24},
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
