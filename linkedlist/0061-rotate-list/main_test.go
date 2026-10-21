package rotatelist

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode, int) *ListNode{
	"rotateRight":            rotateRight,
	"rotateRightEnumeration": rotateRightEnumeration,
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"示例2-k大于长度", []int{0, 1, 2}, 4, []int{2, 0, 1}},
		{"空链表", []int{}, 3, []int{}},
		{"单元素", []int{1}, 99, []int{1}},
		{"k=0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"k等于长度", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k为长度倍数", []int{1, 2}, 4, []int{1, 2}},
		{"k=长度-1", []int{1, 2, 3, 4}, 3, []int{2, 3, 4, 1}},
		{"全相同", []int{7, 7, 7}, 1, []int{7, 7, 7}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.head), tt.k))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.head, tt.k, got, tt.want)
				}
			})
		}
	}
}
