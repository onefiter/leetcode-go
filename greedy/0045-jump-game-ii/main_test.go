package jumpgame2

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"jump":   jump,
	"jumpDP": jumpDP,
}

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{2, 3, 1, 1, 4}, 2},
		{"示例2", []int{2, 3, 0, 1, 4}, 2},
		{"单元素无需跳", []int{0}, 0},
		{"一步到位", []int{5, 1, 1, 1}, 1},
		{"每次只能跳一格", []int{1, 1, 1, 1}, 3},
		{"选落点要看覆盖最远", []int{2, 1, 3, 1, 1, 1}, 2},
		{"两个元素", []int{1, 2}, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.nums))
				if got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
