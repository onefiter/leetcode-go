package partition

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) bool{
	"canPartition":   canPartition,
	"canPartition2D": canPartition2D,
}

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"示例1", []int{1, 5, 11, 5}, true},
		{"示例2", []int{1, 2, 3, 5}, false},
		{"单元素", []int{1}, false},
		{"两个相等", []int{7, 7}, true},
		{"和为奇数", []int{1, 2, 4}, false},
		{"和为偶数但无法分割", []int{1, 2, 5}, false},
		{"多元素可分", []int{3, 3, 3, 4, 5}, true},
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
