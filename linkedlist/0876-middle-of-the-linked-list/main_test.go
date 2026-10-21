package middlenode

import (
	"leetcode-go/template"
	"testing"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"middleNode":            middleNode,
	"middleNodeEnumeration": middleNodeEnumeration,
}

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		idx  int // 期望中间节点的下标，-1 表示 nil
	}{
		{"示例1-奇数长度", []int{1, 2, 3, 4, 5}, 2},
		{"示例2-偶数长度取第二个", []int{1, 2, 3, 4, 5, 6}, 3},
		{"空链表", []int{}, -1},
		{"单节点", []int{1}, 0},
		{"两个节点", []int{1, 2}, 1},
		{"三个节点", []int{1, 2, 3}, 1},
		{"全部相同", []int{8, 8, 8, 8}, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				head := template.BuildList(tt.vals)
				var want *ListNode
				if tt.idx >= 0 {
					want = head
					for i := 0; i < tt.idx; i++ {
						want = want.Next
					}
				}
				// 值可能重复，按指针比较。
				if got := fn(head); got != want {
					t.Errorf("%s(%v) = %v, want node at index %d", name, tt.vals, template.ListToSlice(got), tt.idx)
				}
			})
		}
	}
}
