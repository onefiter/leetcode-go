package circulardeque

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

func run(d *MyCircularDeque, o op) int {
	switch o.name {
	case "InsertFront":
		return b2i(d.InsertFront(o.val))
	case "InsertLast":
		return b2i(d.InsertLast(o.val))
	case "DeleteFront":
		return b2i(d.DeleteFront())
	case "DeleteLast":
		return b2i(d.DeleteLast())
	case "GetFront":
		return d.GetFront()
	case "GetRear":
		return d.GetRear()
	case "IsEmpty":
		return b2i(d.IsEmpty())
	case "IsFull":
		return b2i(d.IsFull())
	}
	panic("unknown op " + o.name)
}

func TestMyCircularDeque(t *testing.T) {
	tests := []struct {
		name string
		k    int
		ops  []op
	}{
		{"示例1", 3, []op{
			{"InsertLast", 1, 1}, {"InsertLast", 2, 1}, {"InsertFront", 3, 1}, {"InsertFront", 4, 0},
			{"GetRear", 0, 2}, {"IsFull", 0, 1}, {"DeleteLast", 0, 1}, {"InsertFront", 4, 1}, {"GetFront", 0, 4},
		}},
		{"空队列", 2, []op{
			{"IsEmpty", 0, 1}, {"GetFront", 0, -1}, {"GetRear", 0, -1},
			{"DeleteFront", 0, 0}, {"DeleteLast", 0, 0}, {"IsFull", 0, 0},
		}},
		{"容量1", 1, []op{
			{"InsertFront", 5, 1}, {"GetFront", 0, 5}, {"GetRear", 0, 5}, {"InsertLast", 6, 0},
			{"DeleteLast", 0, 1}, {"IsEmpty", 0, 1}, {"InsertLast", 6, 1}, {"GetFront", 0, 6},
			{"DeleteFront", 0, 1}, {"IsEmpty", 0, 1},
		}},
		{"只从头部插入", 3, []op{
			{"InsertFront", 1, 1}, {"InsertFront", 2, 1}, {"InsertFront", 3, 1},
			{"GetFront", 0, 3}, {"GetRear", 0, 1}, {"DeleteFront", 0, 1}, {"GetFront", 0, 2},
			{"DeleteLast", 0, 1}, {"GetRear", 0, 2}, {"GetFront", 0, 2},
		}},
		{"两端交替绕环", 4, []op{
			{"InsertLast", 1, 1}, {"InsertFront", 0, 1}, {"InsertLast", 2, 1}, {"InsertFront", -1, 1},
			{"IsFull", 0, 1}, {"GetFront", 0, -1}, {"GetRear", 0, 2}, {"DeleteFront", 0, 1},
			{"DeleteFront", 0, 1}, {"InsertLast", 3, 1}, {"InsertLast", 4, 1}, {"GetFront", 0, 1},
			{"GetRear", 0, 4}, {"DeleteLast", 0, 1}, {"GetRear", 0, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Constructor(tt.k)
			for i, o := range tt.ops {
				if got := run(&d, o); got != o.want {
					t.Fatalf("第%d步 %s(%d) = %d, want %d", i, o.name, o.val, got, o.want)
				}
			}
		})
	}
}
