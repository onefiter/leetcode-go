package allone

import "testing"

// allOner 两种实现共用的接口，同一张操作表跑两遍。
type allOner interface {
	Inc(key string)
	Dec(key string)
	GetMaxKey() string
	GetMinKey() string
}

// op 一次操作：name 为 "inc"、"dec"、"max"、"min"。
// max/min 时 want 为期望的计数（0 表示应返回 ""）；因为可能多个 key 并列，
// 只校验返回 key 的真实计数等于 want，真实计数由测试里的 model 维护。
type op struct {
	name string
	key  string
	want int
}

func TestAllOne(t *testing.T) {
	impls := []struct {
		name string
		new  func() allOner
	}{
		{"计数桶链表", func() allOner { a := Constructor(); return &a }},
		{"扫描哈希", func() allOner { a := NewAllOneMap(); return &a }},
	}
	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"inc", "hello", 0}, {"inc", "hello", 0}, {"max", "", 2}, {"min", "", 2},
			{"inc", "leet", 0}, {"max", "", 2}, {"min", "", 1},
		}},
		{"空结构", []op{
			{"max", "", 0}, {"min", "", 0},
		}},
		{"减到0后删除", []op{
			{"inc", "a", 0}, {"dec", "a", 0}, {"max", "", 0}, {"min", "", 0},
			{"inc", "a", 0}, {"max", "", 1},
		}},
		{"并列最值", []op{
			{"inc", "a", 0}, {"inc", "b", 0}, {"inc", "c", 0},
			{"max", "", 1}, {"min", "", 1}, {"inc", "b", 0}, {"inc", "c", 0},
			{"max", "", 2}, {"min", "", 1},
		}},
		{"桶需新建与合并", []op{
			{"inc", "a", 0}, {"inc", "a", 0}, {"inc", "a", 0}, {"inc", "b", 0},
			{"max", "", 3}, {"min", "", 1}, {"dec", "a", 0}, {"inc", "b", 0},
			{"max", "", 2}, {"min", "", 2}, {"dec", "a", 0}, {"max", "", 2}, {"min", "", 1},
		}},
		{"最大者被减到最小", []op{
			{"inc", "x", 0}, {"inc", "x", 0}, {"inc", "y", 0}, {"inc", "y", 0},
			{"inc", "z", 0}, {"dec", "x", 0}, {"dec", "x", 0},
			{"max", "", 2}, {"min", "", 1}, {"dec", "y", 0}, {"dec", "y", 0},
			{"max", "", 1}, {"min", "", 1}, {"dec", "z", 0}, {"max", "", 0},
		}},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				a := impl.new()
				model := map[string]int{}
				check := func(i int, fn string, got string, want int) {
					if want == 0 {
						if got != "" {
							t.Errorf("op#%d %s() = %q, want \"\"", i, fn, got)
						}
						return
					}
					if model[got] != want {
						t.Errorf("op#%d %s() = %q (count %d), want count %d", i, fn, got, model[got], want)
					}
				}
				for i, o := range tt.ops {
					switch o.name {
					case "inc":
						a.Inc(o.key)
						model[o.key]++
					case "dec":
						a.Dec(o.key)
						if model[o.key]--; model[o.key] == 0 {
							delete(model, o.key)
						}
					case "max":
						check(i, "GetMaxKey", a.GetMaxKey(), o.want)
					case "min":
						check(i, "GetMinKey", a.GetMinKey(), o.want)
					}
				}
			})
		}
	}
}
