package mylinkedlist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

// op 一次操作；get 时 want 为期望返回值。
type op struct {
	name       string
	index, val int
	want       int
}

func TestMyLinkedList(t *testing.T) {
	tests := []struct {
		name  string
		ops   []op
		final []int // 全部操作结束后链表内容
	}{
		{"示例1", []op{
			{"addAtHead", 0, 1, 0}, {"addAtTail", 0, 3, 0}, {"addAtIndex", 1, 2, 0},
			{"get", 1, 0, 2}, {"deleteAtIndex", 1, 0, 0}, {"get", 1, 0, 3},
		}, []int{1, 3}},
		{"空链表get与delete", []op{
			{"get", 0, 0, -1}, {"deleteAtIndex", 0, 0, 0}, {"get", -1, 0, -1},
		}, []int{}},
		{"index等于长度追加到尾部", []op{
			{"addAtIndex", 0, 10, 0}, {"addAtIndex", 1, 20, 0}, {"get", 1, 0, 20},
		}, []int{10, 20}},
		{"index大于长度不插入", []op{
			{"addAtHead", 0, 1, 0}, {"addAtIndex", 5, 9, 0}, {"get", 1, 0, -1},
		}, []int{1}},
		{"删除头节点和尾节点", []op{
			{"addAtTail", 0, 1, 0}, {"addAtTail", 0, 2, 0}, {"addAtTail", 0, 3, 0},
			{"deleteAtIndex", 0, 0, 0}, {"get", 0, 0, 2},
			{"deleteAtIndex", 1, 0, 0}, {"get", 1, 0, -1}, {"get", 0, 0, 2},
		}, []int{2}},
		{"删光后再添加", []op{
			{"addAtHead", 0, 1, 0}, {"deleteAtIndex", 0, 0, 0}, {"get", 0, 0, -1},
			{"addAtTail", 0, 5, 0}, {"get", 0, 0, 5},
		}, []int{5}},
		{"多次头插", []op{
			{"addAtHead", 0, 3, 0}, {"addAtHead", 0, 2, 0}, {"addAtHead", 0, 1, 0},
			{"get", 0, 0, 1}, {"get", 2, 0, 3}, {"get", 3, 0, -1},
		}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Constructor()
			for i, o := range tt.ops {
				switch o.name {
				case "addAtHead":
					l.AddAtHead(o.val)
				case "addAtTail":
					l.AddAtTail(o.val)
				case "addAtIndex":
					l.AddAtIndex(o.index, o.val)
				case "deleteAtIndex":
					l.DeleteAtIndex(o.index)
				case "get":
					if got := l.Get(o.index); got != o.want {
						t.Errorf("op#%d Get(%d) = %d, want %d", i, o.index, got, o.want)
					}
				}
			}
			if got := template.ListToSlice(l.dummy.Next); !slices.Equal(got, tt.final) {
				t.Errorf("final list = %v, want %v", got, tt.final)
			}
			if l.size != len(tt.final) {
				t.Errorf("size = %d, want %d", l.size, len(tt.final))
			}
		})
	}
}
