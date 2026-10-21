package searchrotatedii

import (
	"testing"
)

var solutions = map[string]func([]int, int) bool{
	"search":            search,
	"searchEnumeration": searchEnumeration,
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"示例1-存在", []int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{"示例2-不存在", []int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{"空数组", []int{}, 0, false},
		{"单元素-命中", []int{1}, 1, true},
		{"单元素-不命中", []int{1}, 0, false},
		{"全相同-命中", []int{1, 1, 1, 1, 1}, 1, true},
		{"全相同-不命中", []int{1, 1, 1, 1, 1}, 2, false},
		{"重复干扰-命中", []int{1, 0, 1, 1, 1}, 0, true},
		{"重复干扰-命中2", []int{1, 1, 1, 0, 1}, 0, true},
		{"未旋转-命中", []int{1, 2, 3, 4, 5}, 4, true},
		{"旋转点在头部", []int{5, 1, 2, 3, 4}, 5, true},
		{"旋转点在尾部", []int{2, 3, 4, 5, 1}, 1, true},
		{"比全部都小", []int{4, 5, 6, 0, 1, 2}, -1, false},
		{"比全部都大", []int{4, 5, 6, 0, 1, 2}, 100, false},
		{"两元素重复-命中", []int{1, 1}, 1, true},
		{"两元素重复-不命中", []int{1, 1}, 2, false},
		{"含负数-命中", []int{0, 1, 2, -3, -2, -1}, -3, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}

// TestSearchAgainstEnumeration 用顺序遍历的结果核对二分解法，
// 使用含大量重复元素的基数组，枚举所有旋转位置与目标值。
func TestSearchAgainstEnumeration(t *testing.T) {
	bases := [][]int{
		{1, 1, 1, 2, 2, 3, 1, 1},
		{0, 0, 1, 1, 2, 2},
		{2, 2, 2, 2, 2},
		{-1, -1, 0, 0, 0, 1, 2, 2},
	}
	for _, base := range bases {
		n := len(base)
		for k := 0; k < n; k++ {
			rotated := append(append([]int{}, base[k:]...), base[:k]...)
			for target := -2; target <= 4; target++ {
				want := searchEnumeration(rotated, target)
				if got := search(rotated, target); got != want {
					t.Errorf("search(%v, %d) = %v, want %v", rotated, target, got, want)
				}
			}
		}
	}
}
