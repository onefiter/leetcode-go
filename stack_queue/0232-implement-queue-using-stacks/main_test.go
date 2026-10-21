package queueusingstacks

import "testing"

type queueAPI interface {
	Push(int)
	Pop() int
	Peek() int
	Empty() bool
}

var solutions = map[string]func() queueAPI{
	"MyQueue": func() queueAPI { q := Constructor(); return &q },
	"MyQueueEnumeration": func() queueAPI {
		q := ConstructorEnumeration()
		return &q
	},
}

func TestMyQueue(t *testing.T) {
	// op: push(val) / pop -> want / peek -> want / empty -> wantB
	type step struct {
		op    string
		val   int
		want  int
		wantB bool
	}
	tests := []struct {
		name  string
		steps []step
	}{
		{"示例1", []step{
			{op: "push", val: 1}, {op: "push", val: 2}, {op: "peek", want: 1},
			{op: "pop", want: 1}, {op: "empty", wantB: false},
		}},
		{"初始为空", []step{{op: "empty", wantB: true}}},
		{"单元素", []step{
			{op: "push", val: 3}, {op: "peek", want: 3}, {op: "pop", want: 3}, {op: "empty", wantB: true},
		}},
		{"交替进出", []step{
			{op: "push", val: 1}, {op: "push", val: 2}, {op: "pop", want: 1},
			{op: "push", val: 3}, {op: "push", val: 4}, {op: "pop", want: 2},
			{op: "peek", want: 3}, {op: "pop", want: 3}, {op: "pop", want: 4}, {op: "empty", wantB: true},
		}},
		{"清空后重用", []step{
			{op: "push", val: 8}, {op: "pop", want: 8}, {op: "push", val: 9},
			{op: "peek", want: 9}, {op: "empty", wantB: false},
		}},
	}
	for name, newFn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				q := newFn()
				for i, st := range tt.steps {
					switch st.op {
					case "push":
						q.Push(st.val)
					case "pop":
						if got := q.Pop(); got != st.want {
							t.Fatalf("step %d Pop() = %d, want %d", i, got, st.want)
						}
					case "peek":
						if got := q.Peek(); got != st.want {
							t.Fatalf("step %d Peek() = %d, want %d", i, got, st.want)
						}
					case "empty":
						if got := q.Empty(); got != st.wantB {
							t.Fatalf("step %d Empty() = %v, want %v", i, got, st.wantB)
						}
					}
				}
			})
		}
	}
}
