package findclosest

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int, int) []int{
	"findClosestElements":            findClosestElements,
	"findClosestElementsTwoPointers": findClosestElementsTwoPointers,
	"findClosestElementsEnumeration": findClosestElementsEnumeration,
}

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k, x int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{"示例2-x在左侧外", []int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{"x在右侧外", []int{1, 2, 3, 4, 5}, 2, 10, []int{4, 5}},
		{"单元素", []int{7}, 1, 3, []int{7}},
		{"k等于n", []int{1, 3, 5}, 3, 4, []int{1, 3, 5}},
		{"全相同", []int{2, 2, 2, 2}, 2, 2, []int{2, 2}},
		{"距离相等取小", []int{1, 3}, 1, 2, []int{1}},
		{"有重复", []int{1, 1, 1, 10, 10, 10}, 1, 9, []int{10}},
		{"有重复-平局", []int{1, 1, 2, 2, 2, 2, 2, 3, 3}, 3, 3, []int{2, 3, 3}},
		{"含负数", []int{-5, -3, 0, 2, 8}, 2, -1, []int{-3, 0}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				arr := slices.Clone(tt.arr)
				if got := fn(arr, tt.k, tt.x); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, k=%d, x=%d) = %v, want %v", name, tt.arr, tt.k, tt.x, got, tt.want)
				}
				if !slices.Equal(arr, tt.arr) {
					t.Errorf("%s modified input: %v", name, arr)
				}
			})
		}
	}
}

// TestFindClosestAgainstEnumeration 遍历所有 k、x 与排序对照解比对。
func TestFindClosestAgainstEnumeration(t *testing.T) {
	arr := []int{-4, -1, -1, 0, 2, 2, 2, 5, 7, 7, 11}
	for k := 1; k <= len(arr); k++ {
		for x := -8; x <= 15; x++ {
			want := findClosestElementsEnumeration(arr, k, x)
			for name, fn := range solutions {
				if got := fn(arr, k, x); !slices.Equal(got, want) {
					t.Errorf("%s(k=%d, x=%d) = %v, want %v", name, k, x, got, want)
				}
			}
		}
	}
}
