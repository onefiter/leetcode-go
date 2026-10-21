package minstack

import "testing"

// stackAPI 抽象出两种实现的公共方法，便于跑同一组操作序列。
type stackAPI interface {
	Push(int)
	Pop()
	Top() int
	GetMin() int
}

var solutions = map[string]func() stackAPI{
	"MinStack": func() stackAPI { s := Constructor(); return &s },
	"MinStackEnumeration": func() stackAPI {
		s := ConstructorEnumeration()
		return &s
	},
}

func TestMinStack(t *testing.T) {
	// op: "push" 带参数 val；"pop" 无返回；"top"/"getMin" 断言 want。
	type step struct {
		op   string
		val  int
		want int
	}
	tests := []struct {
		name  string
		steps []step
	}{
		{"示例1", []step{
			{"push", -2, 0}, {"push", 0, 0}, {"push", -3, 0},
			{"getMin", 0, -3}, {"pop", 0, 0}, {"top", 0, 0}, {"getMin", 0, -2},
		}},
		{"单元素", []step{{"push", 5, 0}, {"top", 0, 5}, {"getMin", 0, 5}}},
		{"重复最小值", []step{
			{"push", 1, 0}, {"push", 1, 0}, {"push", 2, 0},
			{"pop", 0, 0}, {"pop", 0, 0}, {"getMin", 0, 1}, {"top", 0, 1},
		}},
		{"递减入栈", []step{
			{"push", 3, 0}, {"push", 2, 0}, {"push", 1, 0}, {"getMin", 0, 1},
			{"pop", 0, 0}, {"getMin", 0, 2}, {"pop", 0, 0}, {"getMin", 0, 3},
		}},
		{"清空后重用", []step{
			{"push", 7, 0}, {"pop", 0, 0}, {"push", 9, 0}, {"getMin", 0, 9}, {"top", 0, 9},
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
						s.Pop()
					case "top":
						if got := s.Top(); got != st.want {
							t.Fatalf("step %d Top() = %d, want %d", i, got, st.want)
						}
					case "getMin":
						if got := s.GetMin(); got != st.want {
							t.Fatalf("step %d GetMin() = %d, want %d", i, got, st.want)
						}
					}
				}
			})
		}
	}
}
