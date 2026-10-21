package twosumii

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) []int{
	"twoSum":            twoSum,
	"twoSumEnumeration": twoSumEnumeration,
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"示例1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"示例2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"示例3-负数", []int{-1, 0}, -1, []int{1, 2}},
		{"两端命中", []int{1, 2, 3, 4, 4, 9, 56, 90}, 91, []int{1, 8}},
		{"含重复元素", []int{3, 3}, 6, []int{1, 2}},
		{"含负数与正数", []int{-10, -3, 0, 5, 9, 12}, 6, []int{2, 5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.numbers, tt.target); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.numbers, tt.target, got, tt.want)
				}
			})
		}
	}
}
