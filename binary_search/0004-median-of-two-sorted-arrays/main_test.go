package mediantwosorted

import "testing"

var solutions = map[string]func([]int, []int) float64{
	"findMedianSortedArrays":            findMedianSortedArrays,
	"findMedianSortedArraysKth":         findMedianSortedArraysKth,
	"findMedianSortedArraysEnumeration": findMedianSortedArraysEnumeration,
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		want         float64
	}{
		{"示例1", []int{1, 3}, []int{2}, 2},
		{"示例2", []int{1, 2}, []int{3, 4}, 2.5},
		{"两个都空", []int{}, []int{}, 0},
		{"第一个空", []int{}, []int{1}, 1},
		{"第二个空", []int{2}, []int{}, 2},
		{"第一个空-偶数", []int{}, []int{1, 2, 3, 4}, 2.5},
		{"全部相同", []int{1, 1}, []int{1, 1}, 1},
		{"含负数", []int{-5, -3}, []int{-2, -1}, -2.5},
		{"完全分离", []int{1, 2, 3}, []int{10, 11}, 3},
		{"交错", []int{1, 3, 5, 7}, []int{2, 4, 6}, 4},
		{"长度悬殊", []int{100}, []int{1, 2, 3, 4, 5, 6}, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums1, tt.nums2); got != tt.want {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.nums1, tt.nums2, got, tt.want)
				}
			})
		}
	}
}

// TestFindMedianAgainstMerge 用归并结果核对其他解法，覆盖各种长度组合与重复值。
func TestFindMedianAgainstMerge(t *testing.T) {
	build := func(n, start, step int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = start + i*step/2 // step/2 制造重复值
		}
		return s
	}
	for m := 0; m <= 6; m++ {
		for n := 0; n <= 6; n++ {
			for start := -3; start <= 3; start++ {
				a := build(m, start, 3)
				b := build(n, 0, 2)
				want := findMedianSortedArraysEnumeration(a, b)
				for name, fn := range solutions {
					if got := fn(a, b); got != want {
						t.Errorf("%s(%v, %v) = %v, want %v", name, a, b, got, want)
					}
				}
			}
		}
	}
}
