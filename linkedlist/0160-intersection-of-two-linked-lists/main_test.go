package intersection

import (
	"leetcode-go/template"
	"testing"
)

var solutions = map[string]func(*ListNode, *ListNode) *ListNode{
	"getIntersectionNode":            getIntersectionNode,
	"getIntersectionNodeHash":        getIntersectionNodeHash,
	"getIntersectionNodeEnumeration": getIntersectionNodeEnumeration,
}

// buildIntersect 构造 A = prefixA + common，B = prefixB + common，
// 两条链表共享同一段 common 尾部节点；返回 headA、headB 和交点（common 为空时为 nil）。
func buildIntersect(prefixA, prefixB, common []int) (*ListNode, *ListNode, *ListNode) {
	shared := template.BuildList(common)
	attach := func(prefix []int) *ListNode {
		head := template.BuildList(prefix)
		if head == nil {
			return shared
		}
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = shared
		return head
	}
	return attach(prefixA), attach(prefixB), shared
}

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name                     string
		prefixA, prefixB, common []int
	}{
		{"示例1", []int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5}},
		{"示例2", []int{1, 9, 1}, []int{3}, []int{2, 4}},
		{"示例3-不相交", []int{2, 6, 4}, []int{1, 5}, nil},
		{"两条都为空", nil, nil, nil},
		{"A为空", nil, []int{1, 2}, nil},
		{"同一条链表-交点为头", nil, nil, []int{1, 2, 3}},
		{"A整条是B的尾部", nil, []int{7, 8}, []int{1, 2}},
		{"交点是最后一个节点", []int{1, 2, 3}, []int{4}, []int{9}},
		{"值相同但不相交", []int{1, 2, 3}, []int{1, 2, 3}, nil},
		{"单节点相交", nil, nil, []int{1}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				headA, headB, want := buildIntersect(tt.prefixA, tt.prefixB, tt.common)
				if got := fn(headA, headB); got != want {
					t.Errorf("%s() = %p, want %p", name, got, want)
				}
				// 解法不应修改链表结构。
				if got := len(template.ListToSlice(headA)); got != len(tt.prefixA)+len(tt.common) {
					t.Errorf("%s 修改了链表 A，长度变为 %d", name, got)
				}
			})
		}
	}
}
