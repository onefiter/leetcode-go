package removenodes

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"removeNodes":            removeNodes,
	"removeNodesReverse":     removeNodesReverse,
	"removeNodesRecursive":   removeNodesRecursive,
	"removeNodesEnumeration": removeNodesEnumeration,
}

func TestRemoveNodes(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{5, 2, 13, 3, 8}, []int{13, 8}},
		{"示例2-全相同", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"单个节点", []int{7}, []int{7}},
		{"严格递增-只剩尾", []int{1, 2, 3, 4}, []int{4}},
		{"严格递减-全保留", []int{4, 3, 2, 1}, []int{4, 3, 2, 1}},
		{"最大在头", []int{9, 1, 5, 2, 3}, []int{9, 5, 3}},
		{"含相等", []int{3, 1, 3, 2, 2}, []int{3, 3, 2, 2}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 原地修改，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.vals)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
