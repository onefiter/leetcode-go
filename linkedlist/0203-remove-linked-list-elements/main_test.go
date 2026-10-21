package removeelements

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode, int) *ListNode{
	"removeElements":            removeElements,
	"removeElementsRecursive":   removeElementsRecursive,
	"removeElementsEnumeration": removeElementsEnumeration,
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		val  int
		want []int
	}{
		{"示例1", []int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{"示例2-空链表", []int{}, 1, []int{}},
		{"示例3-全部相同全删", []int{7, 7, 7, 7}, 7, []int{}},
		{"单节点删除", []int{1}, 1, []int{}},
		{"单节点保留", []int{1}, 2, []int{1}},
		{"删除头节点", []int{1, 2, 3}, 1, []int{2, 3}},
		{"删除尾节点", []int{1, 2, 3}, 3, []int{1, 2}},
		{"连续多个待删", []int{1, 2, 2, 2, 3}, 2, []int{1, 3}},
		{"无需删除", []int{1, 2, 3}, 9, []int{1, 2, 3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法会原地修改链表，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.vals), tt.val))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.vals, tt.val, got, tt.want)
				}
			})
		}
	}
}
