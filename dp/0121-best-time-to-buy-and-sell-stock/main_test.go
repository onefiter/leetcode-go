package maxprofit

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"maxProfit":      maxProfit,
	"maxProfitBrute": maxProfitBrute,
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"示例1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"示例2价格单调下降", []int{7, 6, 4, 3, 1}, 0},
		{"单天无法交易", []int{5}, 0},
		{"单调上涨", []int{1, 2, 3, 4, 5}, 4},
		{"最低点在最高点之后", []int{3, 8, 1, 2}, 5},
		{"价格不变", []int{2, 2, 2}, 0},
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
