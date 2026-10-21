package swappairs

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"swapPairs":            swapPairs,
	"swapPairsRecursive":   swapPairsRecursive,
	"swapPairsEnumeration": swapPairsEnumeration,
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"示例2-空链表", []int{}, []int{}},
		{"示例3-单元素", []int{1}, []int{1}},
		{"两个元素", []int{1, 2}, []int{2, 1}},
		{"奇数个", []int{1, 2, 3}, []int{2, 1, 3}},
		{"全相同", []int{7, 7, 7, 7}, []int{7, 7, 7, 7}},
		{"六个元素", []int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.head)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.head, got, tt.want)
				}
			})
		}
	}
}

// TestSwapPairsSwapsNodes 确认交换的是节点本身而不是节点值。
func TestSwapPairsSwapsNodes(t *testing.T) {
	for name, fn := range solutions {
		head := template.BuildList([]int{1, 2})
		first, second := head, head.Next
		got := fn(head)
		if got != second || got.Next != first {
			t.Errorf("%s: 应交换节点本身", name)
		}
	}
}
