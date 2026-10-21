package targetsum

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"findTargetSumWays":     findTargetSumWays,
	"findTargetSumWaysMemo": findTargetSumWaysMemo,
}

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"示例1", []int{1, 1, 1, 1, 1}, 3, 5},
		{"示例2", []int{1}, 1, 1},
		{"目标超出总和", []int{1, 2}, 4, 0},
		{"奇偶不匹配", []int{1, 2}, 2, 0},
		{"负数目标", []int{1, 2, 3}, -6, 1},
		{"含零翻倍", []int{0, 0, 1}, 1, 4},
		{"目标为零", []int{1, 1, 2}, 0, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums), tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}
