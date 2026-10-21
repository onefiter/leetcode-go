package hascycle

import (
	"leetcode-go/template"
	"testing"
)

var solutions = map[string]func(*ListNode) bool{
	"hasCycle":            hasCycle,
	"hasCycleEnumeration": hasCycleEnumeration,
}

// buildCycleList 构建链表，并把尾节点接到下标 pos 处（pos=-1 表示无环）。
func buildCycleList(vals []int, pos int) *ListNode {
	head := template.BuildList(vals)
	if pos < 0 || head == nil {
		return head
	}
	var entry, tail *ListNode
	for i, cur := 0, head; cur != nil; i, cur = i+1, cur.Next {
		if i == pos {
			entry = cur
		}
		tail = cur
	}
	tail.Next = entry
	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{"示例1", []int{3, 2, 0, -4}, 1, true},
		{"示例2", []int{1, 2}, 0, true},
		{"示例3-单节点无环", []int{1}, -1, false},
		{"空链表", []int{}, -1, false},
		{"单节点自环", []int{1}, 0, true},
		{"无环", []int{1, 2, 3, 4, 5}, -1, false},
		{"尾节点自环", []int{1, 2, 3, 4, 5}, 4, true},
		{"全部相同无环", []int{7, 7, 7}, -1, false},
		{"全部相同有环", []int{7, 7, 7}, 0, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildCycleList(tt.vals, tt.pos)); got != tt.want {
					t.Errorf("%s(%v, pos=%d) = %v, want %v", name, tt.vals, tt.pos, got, tt.want)
				}
			})
		}
	}
}
