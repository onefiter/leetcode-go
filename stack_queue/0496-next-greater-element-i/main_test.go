package nextgreater

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, []int) []int{
	"nextGreaterElement":            nextGreaterElement,
	"nextGreaterElementEnumeration": nextGreaterElementEnumeration,
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"示例1", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"示例2", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{"nums1 为空", []int{}, []int{1, 2}, []int{}},
		{"单元素", []int{5}, []int{5}, []int{-1}},
		{"严格递减", []int{3, 2, 1}, []int{3, 2, 1}, []int{-1, -1, -1}},
		{"严格递增", []int{1, 2, 3}, []int{1, 2, 3}, []int{2, 3, -1}},
		{"末尾元素", []int{7}, []int{1, 9, 7}, []int{-1}},
		{"跨越多个", []int{2}, []int{2, 1, 0, 5}, []int{5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums1, tt.nums2); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.nums1, tt.nums2, got, tt.want)
				}
			})
		}
	}
}
