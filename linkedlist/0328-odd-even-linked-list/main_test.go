package oddevenlist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"oddEvenList":            oddEvenList,
	"oddEvenListEnumeration": oddEvenListEnumeration,
}

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{"示例2", []int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
		{"空链表", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"两个节点", []int{1, 2}, []int{1, 2}},
		{"三个节点", []int{1, 2, 3}, []int{1, 3, 2}},
		{"偶数长度", []int{1, 2, 3, 4, 5, 6}, []int{1, 3, 5, 2, 4, 6}},
		{"全部相同", []int{9, 9, 9, 9}, []int{9, 9, 9, 9}},
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
