package coinchange2

import (
	"slices"
	"testing"
)

var solutions = map[string]func(int, []int) int{
	"change":   change,
	"change2D": change2D,
}

func TestChange(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{"示例1", 5, []int{1, 2, 5}, 4},
		{"示例2", 3, []int{2}, 0},
		{"示例3", 10, []int{10}, 1},
		{"金额为零", 0, []int{7}, 1},
		{"组合不计顺序", 4, []int{1, 2}, 3},
		{"多种面额", 12, []int{2, 3, 5}, 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.amount, slices.Clone(tt.coins)); got != tt.want {
					t.Errorf("%s(%d, %v) = %d, want %d", name, tt.amount, tt.coins, got, tt.want)
				}
			})
		}
	}
}
