package jumpgame

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) bool{
	"canJump":   canJump,
	"canJumpDP": canJumpDP,
}

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"示例1", []int{2, 3, 1, 1, 4}, true},
		{"示例2", []int{3, 2, 1, 0, 4}, false},
		{"单元素", []int{0}, true},
		{"开头为零", []int{0, 1}, false},
		{"越过零点", []int{2, 0, 0}, true},
		{"末尾为零不影响", []int{1, 1, 0}, true},
		{"中间断点", []int{1, 0, 1, 0}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.nums))
				if got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
