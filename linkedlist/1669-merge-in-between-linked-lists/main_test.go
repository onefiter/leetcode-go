package mergeinbetween

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode, int, int, *ListNode) *ListNode{
	"mergeInBetween":            mergeInBetween,
	"mergeInBetweenEnumeration": mergeInBetweenEnumeration,
}

func TestMergeInBetween(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		a, b  int
		list2 []int
		want  []int
	}{
		{"示例1", []int{10, 1, 13, 6, 9, 5}, 3, 4, []int{1000000, 1000001, 1000002}, []int{10, 1, 13, 1000000, 1000001, 1000002, 5}},
		{"示例2", []int{0, 1, 2, 3, 4, 5, 6}, 2, 5, []int{1000000, 1000001, 1000002, 1000003, 1000004}, []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}},
		{"最短list1-删一个", []int{1, 2, 3}, 1, 1, []int{9}, []int{1, 9, 3}},
		{"删除除头尾外全部", []int{1, 2, 3, 4, 5}, 1, 3, []int{7, 8}, []int{1, 7, 8, 5}},
		{"list2单节点", []int{1, 2, 3, 4}, 2, 2, []int{0}, []int{1, 2, 0, 4}},
		{"全相同", []int{5, 5, 5, 5}, 1, 2, []int{5}, []int{5, 5, 5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 原地修改，每次重新构造输入。
				l1, l2 := template.BuildList(tt.list1), template.BuildList(tt.list2)
				got := template.ListToSlice(fn(l1, tt.a, tt.b, l2))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d, %d, %v) = %v, want %v", name, tt.list1, tt.a, tt.b, tt.list2, got, tt.want)
				}
			})
		}
	}
}
