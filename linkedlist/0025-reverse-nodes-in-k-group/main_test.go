package reversekgroup

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode, int) *ListNode{
	"reverseKGroup":            reverseKGroup,
	"reverseKGroupRecursive":   reverseKGroupRecursive,
	"reverseKGroupEnumeration": reverseKGroupEnumeration,
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"示例2", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{"空链表", []int{}, 2, []int{}},
		{"单元素k=1", []int{1}, 1, []int{1}},
		{"k=1不变", []int{1, 2, 3}, 1, []int{1, 2, 3}},
		{"k等于长度", []int{1, 2, 3, 4}, 4, []int{4, 3, 2, 1}},
		{"k大于长度", []int{1, 2}, 3, []int{1, 2}},
		{"恰好整组", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},
		{"全相同", []int{5, 5, 5}, 2, []int{5, 5, 5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.head), tt.k))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.head, tt.k, got, tt.want)
				}
			})
		}
	}
}
