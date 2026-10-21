package assigncookies

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, []int) int{
	"findContentChildren":      findContentChildren,
	"findContentChildrenBrute": findContentChildrenBrute,
}

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		name string
		g    []int
		s    []int
		want int
	}{
		{"示例1", []int{1, 2, 3}, []int{1, 1}, 1},
		{"示例2", []int{1, 2}, []int{1, 2, 3}, 2},
		{"没有饼干", []int{1, 2}, []int{}, 0},
		{"饼干都太小", []int{5, 6}, []int{1, 2, 3}, 0},
		{"乱序输入", []int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
		{"大饼干不能浪费", []int{1, 3}, []int{3, 1}, 2},
		{"重复胃口", []int{2, 2, 2}, []int{2, 2}, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.g), slices.Clone(tt.s))
				if got != tt.want {
					t.Errorf("%s(%v, %v) = %d, want %d", name, tt.g, tt.s, got, tt.want)
				}
			})
		}
	}
}
