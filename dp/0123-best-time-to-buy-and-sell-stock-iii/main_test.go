package maxprofit3

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"maxProfit":      maxProfit,
	"maxProfitSplit": maxProfitSplit,
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"示例1", []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{"示例2一笔即最优", []int{1, 2, 3, 4, 5}, 4},
		{"示例3单调下降", []int{7, 6, 4, 3, 1}, 0},
		{"示例4单天", []int{1}, 0},
		{"三段上涨只取两段", []int{1, 5, 2, 8, 3, 10}, 14},
		{"短数组", []int{2, 4, 1}, 2},
		{"两段完整波动", []int{1, 4, 2, 7}, 8},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.prices)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.prices, got, tt.want)
				}
			})
		}
	}
}
