package coinchange

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"coinChange":    coinChange,
	"coinChangeBFS": coinChangeBFS,
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{"示例1", []int{1, 2, 5}, 11, 3},
		{"示例2-无解", []int{2}, 3, -1},
		{"示例3-金额为0", []int{1}, 0, 0},
		{"贪心失效", []int{1, 3, 4}, 6, 2},
		{"大面额用不上", []int{5, 10}, 3, -1},
		{"单一面额整除", []int{7}, 28, 4},
		{"较大金额", []int{186, 419, 83, 408}, 6249, 20},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.coins), tt.amount); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.coins, tt.amount, got, tt.want)
				}
			})
		}
	}
}
