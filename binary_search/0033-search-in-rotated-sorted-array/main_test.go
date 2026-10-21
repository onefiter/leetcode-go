package searchrotated

import (
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"search":            search,
	"searchEnumeration": searchEnumeration,
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"示例1", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"示例2-不存在", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"示例3-空数组", []int{}, 0, -1},
		{"单元素-命中", []int{1}, 1, 0},
		{"单元素-不命中", []int{1}, 0, -1},
		{"未旋转-命中头", []int{1, 2, 3, 4, 5}, 1, 0},
		{"未旋转-命中尾", []int{1, 2, 3, 4, 5}, 5, 4},
		{"旋转点在头部", []int{5, 1, 2, 3, 4}, 1, 1},
		{"旋转点在尾部", []int{2, 3, 4, 5, 1}, 1, 4},
		{"命中旋转点", []int{6, 7, 0, 1, 2}, 0, 2},
		{"命中最大值", []int{4, 5, 6, 7, 0, 1, 2}, 7, 3},
		{"比全部都小", []int{4, 5, 6, 7, 0, 1, 2}, -1, -1},
		{"比全部都大", []int{4, 5, 6, 7, 0, 1, 2}, 100, -1},
		{"两元素-命中前", []int{3, 1}, 3, 0},
		{"两元素-命中后", []int{3, 1}, 1, 1},
		{"含负数", []int{3, 4, -2, -1, 0, 1, 2}, -2, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}

// TestSearchAgainstEnumeration 用顺序遍历的结果核对二分解法，
// 枚举所有旋转位置 + 所有可能目标值（含不存在的值）。
func TestSearchAgainstEnumeration(t *testing.T) {
	base := []int{-5, -3, -1, 0, 2, 4, 6, 8, 10}
	n := len(base)
	for k := 0; k < n; k++ {
		// 旋转 k 位：rotated = base[k:] + base[:k]
		rotated := append(append([]int{}, base[k:]...), base[:k]...)
		for target := -6; target <= 11; target++ {
			want := searchEnumeration(rotated, target)
			if got := search(rotated, target); got != want {
				t.Errorf("search(%v, %d) = %d, want %d", rotated, target, got, want)
			}
		}
	}
}
