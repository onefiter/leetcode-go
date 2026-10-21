package snapshotarray

import "testing"

// op 描述一次操作：Set(a, b) / Snap() / Get(a, b)，want 只对 Snap/Get 有意义。
type op struct {
	name string
	a, b int
	want int
}

func TestSnapshotArray(t *testing.T) {
	tests := []struct {
		name   string
		length int
		ops    []op
	}{
		{"示例1", 3, []op{
			{"Set", 0, 5, 0}, {"Snap", 0, 0, 0}, {"Set", 0, 6, 0}, {"Get", 0, 0, 5},
		}},
		{"未设置默认0", 2, []op{
			{"Get", 1, 0, 0}, {"Snap", 0, 0, 0}, {"Get", 0, 0, 0}, {"Get", 1, 0, 0},
		}},
		{"单元素多次快照", 1, []op{
			{"Set", 0, 1, 0}, {"Snap", 0, 0, 0}, {"Snap", 0, 0, 1}, {"Set", 0, 2, 0},
			{"Snap", 0, 0, 2}, {"Get", 0, 0, 1}, {"Get", 0, 1, 1}, {"Get", 0, 2, 2},
		}},
		{"同一周期多次Set取最后一次", 2, []op{
			{"Set", 1, 3, 0}, {"Set", 1, 4, 0}, {"Set", 1, 7, 0}, {"Snap", 0, 0, 0},
			{"Set", 1, 9, 0}, {"Get", 1, 0, 7},
		}},
		{"先快照后设置", 3, []op{
			{"Snap", 0, 0, 0}, {"Snap", 0, 0, 1}, {"Set", 2, 8, 0}, {"Snap", 0, 0, 2},
			{"Get", 2, 0, 0}, {"Get", 2, 1, 0}, {"Get", 2, 2, 8},
		}},
		{"设回0", 1, []op{
			{"Set", 0, 4, 0}, {"Snap", 0, 0, 0}, {"Set", 0, 0, 0}, {"Snap", 0, 0, 1},
			{"Get", 0, 0, 4}, {"Get", 0, 1, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := Constructor(tt.length)
			for i, o := range tt.ops {
				switch o.name {
				case "Set":
					sa.Set(o.a, o.b)
				case "Snap":
					if got := sa.Snap(); got != o.want {
						t.Fatalf("第%d步 Snap() = %d, want %d", i, got, o.want)
					}
				case "Get":
					if got := sa.Get(o.a, o.b); got != o.want {
						t.Fatalf("第%d步 Get(%d, %d) = %d, want %d", i, o.a, o.b, got, o.want)
					}
				}
			}
		})
	}
}
