package detectcycle

import (
	"leetcode-go/template"
	"testing"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"detectCycle":            detectCycle,
	"detectCycleEnumeration": detectCycleEnumeration,
}

// buildCycleList 构建链表，把尾节点接到下标 pos 处（pos=-1 表示无环），
// 返回头节点和期望的入环节点指针（无环时为 nil）。
func buildCycleList(vals []int, pos int) (head, entry *ListNode) {
	head = template.BuildList(vals)
	if pos < 0 || head == nil {
		return head, nil
	}
	var tail *ListNode
	for i, cur := 0, head; cur != nil; i, cur = i+1, cur.Next {
		if i == pos {
			entry = cur
		}
		tail = cur
	}
	tail.Next = entry
	return head, entry
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
	}{
		{"示例1", []int{3, 2, 0, -4}, 1},
		{"示例2", []int{1, 2}, 0},
		{"示例3-单节点无环", []int{1}, -1},
		{"空链表", []int{}, -1},
		{"单节点自环", []int{1}, 0},
		{"无环", []int{1, 2, 3, 4, 5}, -1},
		{"尾节点自环", []int{1, 2, 3, 4, 5}, 4},
		{"整条成环", []int{1, 2, 3, 4, 5, 6}, 0},
		{"长尾短环", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6},
		{"全部相同有环", []int{7, 7, 7, 7}, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				head, want := buildCycleList(tt.vals, tt.pos)
				// 值可能重复，必须按指针比较。
				if got := fn(head); got != want {
					t.Errorf("%s(%v, pos=%d) = %p, want %p", name, tt.vals, tt.pos, got, want)
				}
			})
		}
	}
}
