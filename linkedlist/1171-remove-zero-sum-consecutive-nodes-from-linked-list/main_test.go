package removezerosum

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"removeZeroSumSublists":            removeZeroSumSublists,
	"removeZeroSumSublistsEnumeration": removeZeroSumSublistsEnumeration,
}

// hasZeroSumSegment 判断切片中是否还存在和为 0 的连续段。
func hasZeroSumSegment(a []int) bool {
	for i := range a {
		s := 0
		for j := i; j < len(a); j++ {
			s += a[j]
			if s == 0 {
				return true
			}
		}
	}
	return false
}

func TestRemoveZeroSumSublists(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		wants [][]int // 题目允许多个正确答案
	}{
		{"示例1", []int{1, 2, -3, 3, 1}, [][]int{{3, 1}, {1, 2, 1}}},
		{"示例2", []int{1, 2, 3, -3, 4}, [][]int{{1, 2, 4}}},
		{"示例3", []int{1, 2, 3, -3, -2}, [][]int{{1}}},
		{"单个0", []int{0}, [][]int{{}}},
		{"单个非0", []int{1}, [][]int{{1}}},
		{"全为0", []int{0, 0, 0}, [][]int{{}}},
		{"整体和为0", []int{1, -1}, [][]int{{}}},
		{"尾部是0", []int{2, 0}, [][]int{{2}}},
		{"多段", []int{1, 3, 2, -3, -2, 5, 5, -5, 1}, [][]int{{1, 5, 1}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 原地修改，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.vals)))
				if hasZeroSumSegment(got) {
					t.Fatalf("%s(%v) = %v 仍含零和段", name, tt.vals, got)
				}
				ok := false
				for _, w := range tt.wants {
					if slices.Equal(got, w) {
						ok = true
						break
					}
				}
				if !ok {
					t.Errorf("%s(%v) = %v, want one of %v", name, tt.vals, got, tt.wants)
				}
			})
		}
	}
}
