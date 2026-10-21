package median

import "testing"

// finder 两种实现共同的接口。
type finder interface {
	AddNum(num int)
	FindMedian() float64
}

var solutions = map[string]func() finder{
	"MedianFinder": func() finder {
		f := Constructor()
		return &f
	},
	"MedianFinderSorted": func() finder {
		f := ConstructorSorted()
		return &f
	},
}

// op 一次操作：name 为 "add" 或 "find"；add 时 num 为加入的数，find 时 want 为期望中位数。
type op struct {
	name string
	num  int
	want float64
}

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"add", 1, 0}, {"add", 2, 0}, {"find", 0, 1.5},
			{"add", 3, 0}, {"find", 0, 2},
		}},
		{"单个元素", []op{{"add", 5, 0}, {"find", 0, 5}}},
		{"降序加入", []op{
			{"add", 5, 0}, {"find", 0, 5}, {"add", 4, 0}, {"find", 0, 4.5},
			{"add", 3, 0}, {"find", 0, 4}, {"add", 2, 0}, {"find", 0, 3.5},
		}},
		{"含负数", []op{
			{"add", -1, 0}, {"find", 0, -1}, {"add", -2, 0}, {"find", 0, -1.5},
			{"add", -3, 0}, {"find", 0, -2},
		}},
		{"重复元素", []op{
			{"add", 2, 0}, {"add", 2, 0}, {"add", 2, 0}, {"find", 0, 2},
			{"add", 8, 0}, {"find", 0, 2},
		}},
		{"乱序加入", []op{
			{"add", 6, 0}, {"add", 10, 0}, {"add", 2, 0}, {"find", 0, 6},
			{"add", 6, 0}, {"find", 0, 6}, {"add", 5, 0}, {"add", 0, 0}, {"find", 0, 5.5},
		}},
	}
	for name, newFinder := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				f := newFinder()
				for i, o := range tt.ops {
					switch o.name {
					case "add":
						f.AddNum(o.num)
					case "find":
						if got := f.FindMedian(); got != o.want {
							t.Errorf("op#%d FindMedian() = %v, want %v", i, got, o.want)
						}
					}
				}
			})
		}
	}
}
