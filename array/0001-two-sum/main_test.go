package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		if got := twoSum(tt.nums, tt.target); !slices.Equal(got, tt.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func TestTwoSumEnumeration(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
		{[]int{1, 2, 3}, 7, nil},
		{[]int{}, 0, nil},
	}
	for _, tt := range tests {
		if got := twoSumEnumeration(tt.nums, tt.target); !slices.Equal(got, tt.want) {
			t.Errorf("twoSumEnumeration(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
