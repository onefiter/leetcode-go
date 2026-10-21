package burstballoons

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"maxCoins":     maxCoins,
	"maxCoinsMemo": maxCoinsMemo,
}

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{3, 1, 5, 8}, 167},
		{"示例2", []int{1, 5}, 10},
		{"单个气球", []int{7}, 7},
		{"含零", []int{0, 2}, 2},
		{"全为一", []int{1, 1, 1}, 3},
		{"三个气球", []int{2, 3, 4}, 36},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums)); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
