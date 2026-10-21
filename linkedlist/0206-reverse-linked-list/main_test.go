package reverselist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"reverseList":            reverseList,
	"reverseListRecursive":   reverseListRecursive,
	"reverseListEnumeration": reverseListEnumeration,
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"示例2", []int{1, 2}, []int{2, 1}},
		{"示例3-空链表", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"全部相同", []int{4, 4, 4}, []int{4, 4, 4}},
		{"含负数", []int{-1, 0, 3}, []int{3, 0, -1}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法会原地修改链表，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.vals)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
