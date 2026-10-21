package maxonesii

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"findMaxConsecutiveOnes":            findMaxConsecutiveOnes,
	"findMaxConsecutiveOnesDP":          findMaxConsecutiveOnesDP,
	"findMaxConsecutiveOnesEnumeration": findMaxConsecutiveOnesEnumeration,
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{1, 0, 1, 1, 0}, 4},
		{"示例2", []int{1, 0, 1, 1, 0, 1}, 4},
		{"空数组", []int{}, 0},
		{"单个0", []int{0}, 1},
		{"单个1", []int{1}, 1},
		{"全1", []int{1, 1, 1}, 3},
		{"全0", []int{0, 0, 0}, 1},
		{"0在开头", []int{0, 1, 1, 1}, 4},
		{"0在末尾", []int{1, 1, 0}, 3},
		{"多段", []int{1, 1, 0, 0, 1, 1, 1, 0, 1, 1}, 6},
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
