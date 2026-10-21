package deletemiddle

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"deleteMiddle":            deleteMiddle,
	"deleteMiddleEnumeration": deleteMiddleEnumeration,
}

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{"示例2", []int{1, 2, 3, 4}, []int{1, 2, 4}},
		{"示例3-两个节点", []int{2, 1}, []int{2}},
		{"单个节点", []int{1}, []int{}},
		{"三个节点", []int{1, 2, 3}, []int{1, 3}},
		{"全相同", []int{7, 7, 7, 7, 7}, []int{7, 7, 7, 7}},
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
