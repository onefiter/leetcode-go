package intersection

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, []int) []int{
	"intersection":            intersection,
	"intersectionSort":        intersectionSort,
	"intersectionEnumeration": intersectionEnumeration,
}

// normalize 排序结果，便于顺序无关比较。
func normalize(a []int) []int {
	cp := slices.Clone(a)
	slices.Sort(cp)
	return cp
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"示例1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{"示例2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{"无交集", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{"空数组1", []int{}, []int{1, 2}, []int{}},
		{"空数组2", []int{1, 2}, []int{}, []int{}},
		{"都为空", []int{}, []int{}, []int{}},
		{"全相同", []int{1, 1, 1}, []int{1, 1}, []int{1}},
		{"含负数", []int{-1, -2, 3}, []int{-2, 3, 3}, []int{-2, 3}},
		{"完全相同集合", []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2, 3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(tt.nums1, tt.nums2))
				want := normalize(tt.want)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%v, %v) = %v, want (顺序无关) %v",
						name, tt.nums1, tt.nums2, got, want)
				}
			})
		}
	}
}
