package removenthfromend

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode, int) *ListNode{
	"removeNthFromEnd":            removeNthFromEnd,
	"removeNthFromEndEnumeration": removeNthFromEndEnumeration,
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		n    int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"示例2-单节点", []int{1}, 1, []int{}},
		{"示例3", []int{1, 2}, 1, []int{1}},
		{"删除头节点", []int{1, 2}, 2, []int{2}},
		{"删除尾节点", []int{1, 2, 3}, 1, []int{1, 2}},
		{"全部相同", []int{7, 7, 7, 7}, 3, []int{7, 7, 7}},
		{"删除中间", []int{10, 20, 30, 40, 50}, 3, []int{10, 20, 40, 50}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法会原地修改链表，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.vals), tt.n))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.vals, tt.n, got, tt.want)
				}
			})
		}
	}
}
