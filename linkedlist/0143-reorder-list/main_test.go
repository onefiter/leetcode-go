package reorderlist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode){
	"reorderList":            reorderList,
	"reorderListEnumeration": reorderListEnumeration,
}

func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1-偶数长度", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{"示例2-奇数长度", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{"空链表", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"两个节点", []int{1, 2}, []int{1, 2}},
		{"三个节点", []int{1, 2, 3}, []int{1, 3, 2}},
		{"全部相同", []int{6, 6, 6, 6, 6}, []int{6, 6, 6, 6, 6}},
		{"六个节点", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法原地修改链表，每次重新构造输入。
				head := template.BuildList(tt.vals)
				fn(head)
				got := template.ListToSlice(head)
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
