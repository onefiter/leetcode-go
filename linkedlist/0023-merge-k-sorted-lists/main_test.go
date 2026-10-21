package mergeksortedlists

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func([]*ListNode) *ListNode{
	"mergeKLists":            mergeKLists,
	"mergeKListsHeap":        mergeKListsHeap,
	"mergeKListsEnumeration": mergeKListsEnumeration,
}

// buildLists 每次调用重新构造输入（解法会复用并修改原节点）。
func buildLists(vals [][]int) []*ListNode {
	lists := make([]*ListNode, len(vals))
	for i, v := range vals {
		lists[i] = template.BuildList(v)
	}
	return lists
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"示例1", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"示例2-没有链表", [][]int{}, []int{}},
		{"示例3-一条空链表", [][]int{{}}, []int{}},
		{"多条空链表", [][]int{{}, {}, {}}, []int{}},
		{"单条链表", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"单元素链表们", [][]int{{3}, {1}, {2}}, []int{1, 2, 3}},
		{"全相同", [][]int{{2, 2}, {2}, {2, 2, 2}}, []int{2, 2, 2, 2, 2, 2}},
		{"含空与负数", [][]int{{}, {-2, 0, 5}, {}, {-3, 1}}, []int{-3, -2, 0, 1, 5}},
		{"区间不重叠", [][]int{{7, 8}, {1, 2}, {4, 5}}, []int{1, 2, 4, 5, 7, 8}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(buildLists(tt.lists)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.lists, got, tt.want)
				}
			})
		}
	}
}
