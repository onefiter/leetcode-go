package circularqueue

import "testing"

// op 描述一次操作及期望返回值（bool 操作用 1/0 表示 true/false）。
type op struct {
	name string
	val  int
	want int
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func run(q *MyCircularQueue, o op) int {
	switch o.name {
	case "EnQueue":
		return b2i(q.EnQueue(o.val))
	case "DeQueue":
		return b2i(q.DeQueue())
	case "Front":
		return q.Front()
	case "Rear":
		return q.Rear()
	case "IsEmpty":
		return b2i(q.IsEmpty())
	case "IsFull":
		return b2i(q.IsFull())
	}
	panic("unknown op " + o.name)
}

func TestMyCircularQueue(t *testing.T) {
	tests := []struct {
		name string
		k    int
		ops  []op
	}{
		{"示例1", 3, []op{
			{"EnQueue", 1, 1}, {"EnQueue", 2, 1}, {"EnQueue", 3, 1}, {"EnQueue", 4, 0},
			{"Rear", 0, 3}, {"IsFull", 0, 1}, {"DeQueue", 0, 1}, {"EnQueue", 4, 1}, {"Rear", 0, 4},
		}},
		{"空队列", 2, []op{
			{"IsEmpty", 0, 1}, {"Front", 0, -1}, {"Rear", 0, -1}, {"DeQueue", 0, 0}, {"IsFull", 0, 0},
		}},
		{"容量1", 1, []op{
			{"EnQueue", 7, 1}, {"IsFull", 0, 1}, {"Front", 0, 7}, {"Rear", 0, 7}, {"EnQueue", 8, 0},
			{"DeQueue", 0, 1}, {"IsEmpty", 0, 1}, {"EnQueue", 9, 1}, {"Front", 0, 9},
		}},
		{"多次绕环", 3, []op{
			{"EnQueue", 1, 1}, {"EnQueue", 2, 1}, {"DeQueue", 0, 1}, {"EnQueue", 3, 1},
			{"EnQueue", 4, 1}, {"Front", 0, 2}, {"Rear", 0, 4}, {"DeQueue", 0, 1}, {"DeQueue", 0, 1},
			{"EnQueue", 5, 1}, {"EnQueue", 6, 1}, {"Front", 0, 4}, {"Rear", 0, 6}, {"IsFull", 0, 1},
		}},
		{"值全相同", 2, []op{
			{"EnQueue", 5, 1}, {"EnQueue", 5, 1}, {"Front", 0, 5}, {"Rear", 0, 5}, {"DeQueue", 0, 1},
			{"DeQueue", 0, 1}, {"DeQueue", 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Constructor(tt.k)
			for i, o := range tt.ops {
				if got := run(&q, o); got != o.want {
					t.Fatalf("第%d步 %s(%d) = %d, want %d", i, o.name, o.val, got, o.want)
				}
			}
		})
	}
}
