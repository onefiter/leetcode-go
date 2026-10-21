package mergetwolists

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode, *ListNode) *ListNode{
	"mergeTwoLists":            mergeTwoLists,
	"mergeTwoListsRecursive":   mergeTwoListsRecursive,
	"mergeTwoListsEnumeration": mergeTwoListsEnumeration,
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		l1, l2 []int
		want   []int
	}{
		{"示例1", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"示例2-都为空", []int{}, []int{}, []int{}},
		{"示例3-一个为空", []int{}, []int{0}, []int{0}},
		{"另一个为空", []int{1, 2}, []int{}, []int{1, 2}},
		{"单元素", []int{2}, []int{1}, []int{1, 2}},
		{"全部相同", []int{3, 3}, []int{3, 3, 3}, []int{3, 3, 3, 3, 3}},
		{"不交叉", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"含负数", []int{-5, 0, 7}, []int{-3, 2}, []int{-5, -3, 0, 2, 7}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法会原地拼接链表，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.l1), template.BuildList(tt.l2)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.l1, tt.l2, got, tt.want)
				}
			})
		}
	}
}
