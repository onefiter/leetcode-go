package firstunique

import "testing"

// op：Show 表示 ShowFirstUnique()（断言 want），Add 表示 Add(val)。
type op struct {
	name string
	val  int
	want int
}

func TestFirstUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		ops  []op
	}{
		{"示例1", []int{2, 3, 5}, []op{
			{"Show", 0, 2}, {"Add", 5, 0}, {"Show", 0, 2}, {"Add", 2, 0}, {"Show", 0, 3},
			{"Add", 3, 0}, {"Show", 0, -1},
		}},
		{"示例2", []int{7, 7, 7, 7, 7, 7}, []op{
			{"Show", 0, -1}, {"Add", 7, 0}, {"Add", 3, 0}, {"Add", 3, 0}, {"Add", 7, 0},
			{"Add", 17, 0}, {"Show", 0, 17},
		}},
		{"示例3-单元素", []int{809}, []op{
			{"Show", 0, 809}, {"Add", 809, 0}, {"Show", 0, -1},
		}},
		{"空初始化", []int{}, []op{
			{"Show", 0, -1}, {"Add", 1, 0}, {"Show", 0, 1}, {"Add", 2, 0}, {"Show", 0, 1},
			{"Add", 1, 0}, {"Show", 0, 2},
		}},
		{"唯一数在末尾", []int{1, 2, 1, 2, 3}, []op{
			{"Show", 0, 3}, {"Add", 3, 0}, {"Show", 0, -1}, {"Add", 4, 0}, {"Show", 0, 4},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := Constructor(tt.nums)
			for i, o := range tt.ops {
				switch o.name {
				case "Add":
					fu.Add(o.val)
				case "Show":
					if got := fu.ShowFirstUnique(); got != o.want {
						t.Fatalf("第%d步 ShowFirstUnique() = %d, want %d", i, got, o.want)
					}
				}
			}
		})
	}
}
