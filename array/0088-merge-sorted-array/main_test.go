package mergesortedarray

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int, []int, int){
	"merge":            merge,
	"mergeEnumeration": mergeEnumeration,
}

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			got := slices.Clone(tt.nums1)
			fn(got, tt.m, tt.nums2, tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("%s(%v, %d, %v, %d) = %v, want %v", name, tt.nums1, tt.m, tt.nums2, tt.n, got, tt.want)
			}
		}
	}
}
