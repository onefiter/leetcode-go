package deleteduplicatesii

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"deleteDuplicates":            deleteDuplicates,
	"deleteDuplicatesRecursive":   deleteDuplicatesRecursive,
	"deleteDuplicatesEnumeration": deleteDuplicatesEnumeration,
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{"示例2-表头重复", []int{1, 1, 1, 2, 3}, []int{2, 3}},
		{"空链表", []int{}, []int{}},
		{"单元素", []int{1}, []int{1}},
		{"全相同", []int{2, 2, 2}, []int{}},
		{"无重复", []int{1, 2, 3}, []int{1, 2, 3}},
		{"表尾重复", []int{1, 2, 3, 3}, []int{1, 2}},
		{"全部成对重复", []int{1, 1, 2, 2}, []int{}},
		{"含负数", []int{-3, -1, -1, 0}, []int{-3, 0}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.head)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.head, got, tt.want)
				}
			})
		}
	}
}
