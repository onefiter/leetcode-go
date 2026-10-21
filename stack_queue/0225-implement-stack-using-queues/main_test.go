package stackusingqueues

import "testing"

type stackAPI interface {
	Push(int)
	Pop() int
	Top() int
	Empty() bool
}

var solutions = map[string]func() stackAPI{
	"MyStack": func() stackAPI { s := Constructor(); return &s },
	"MyStackTwoQueues": func() stackAPI {
		s := ConstructorTwoQueues()
		return &s
	},
}

func TestMyStack(t *testing.T) {
	// op: push(val) / pop -> want / top -> want / empty -> wantB
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
			{op: "push", val: 1}, {op: "push", val: 2}, {op: "top", want: 2},
			{op: "pop", want: 2}, {op: "empty", wantB: false},
		}},
		{"初始为空", []step{{op: "empty", wantB: true}}},
		{"单元素", []step{
			{op: "push", val: 9}, {op: "top", want: 9}, {op: "pop", want: 9}, {op: "empty", wantB: true},
		}},
		{"后进先出", []step{
			{op: "push", val: 1}, {op: "push", val: 2}, {op: "push", val: 3},
			{op: "pop", want: 3}, {op: "pop", want: 2}, {op: "push", val: 4},
			{op: "top", want: 4}, {op: "pop", want: 4}, {op: "pop", want: 1}, {op: "empty", wantB: true},
		}},
		{"清空后重用", []step{
			{op: "push", val: 5}, {op: "pop", want: 5}, {op: "push", val: 6}, {op: "push", val: 7},
			{op: "top", want: 7}, {op: "pop", want: 7}, {op: "top", want: 6},
		}},
	}
	for name, newFn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				s := newFn()
				for i, st := range tt.steps {
					switch st.op {
					case "push":
						s.Push(st.val)
					case "pop":
						if got := s.Pop(); got != st.want {
							t.Fatalf("step %d Pop() = %d, want %d", i, got, st.want)
						}
					case "top":
						if got := s.Top(); got != st.want {
							t.Fatalf("step %d Top() = %d, want %d", i, got, st.want)
						}
					case "empty":
						if got := s.Empty(); got != st.wantB {
							t.Fatalf("step %d Empty() = %v, want %v", i, got, st.wantB)
						}
					}
				}
			})
		}
	}
}
