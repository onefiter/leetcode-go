package randomset

import "testing"

// randomSet 两种实现共用的接口，同一张操作表跑两遍。
type randomSet interface {
	Insert(val int) bool
	Remove(val int) bool
	GetRandom() int
}

// op 一次操作：name 为 "insert"、"remove" 或 "random"。
// insert/remove 时 want 为期望返回值；random 时 members 为当前集合，校验返回值属于它，
// 并多次调用检查每个成员都至少出现过一次（分布覆盖）。
type op struct {
	name    string
	val     int
	want    bool
	members []int
}

func TestRandomizedSet(t *testing.T) {
	impls := []struct {
		name string
		new  func() randomSet
	}{
		{"数组加哈希", func() randomSet { s := Constructor(); return &s }},
		{"仅哈希", func() randomSet { s := NewRandomizedSetMap(); return &s }},
	}
	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"insert", 1, true, nil}, {"remove", 2, false, nil}, {"insert", 2, true, nil},
			{"random", 0, false, []int{1, 2}}, {"remove", 1, true, nil},
			{"insert", 2, false, nil}, {"random", 0, false, []int{2}},
		}},
		{"重复插入", []op{
			{"insert", 5, true, nil}, {"insert", 5, false, nil}, {"random", 0, false, []int{5}},
		}},
		{"删除不存在", []op{
			{"remove", 1, false, nil}, {"insert", 1, true, nil},
			{"remove", 1, true, nil}, {"remove", 1, false, nil},
		}},
		{"删除末尾元素", []op{
			{"insert", 1, true, nil}, {"insert", 2, true, nil}, {"insert", 3, true, nil},
			{"remove", 3, true, nil}, {"random", 0, false, []int{1, 2}},
			{"insert", 3, true, nil}, {"random", 0, false, []int{1, 2, 3}},
		}},
		{"删除中间元素", []op{
			{"insert", 10, true, nil}, {"insert", 20, true, nil}, {"insert", 30, true, nil},
			{"insert", 40, true, nil}, {"remove", 20, true, nil}, {"remove", 10, true, nil},
			{"random", 0, false, []int{30, 40}}, {"remove", 40, true, nil},
			{"random", 0, false, []int{30}},
		}},
		{"负数与清空后重建", []op{
			{"insert", -1, true, nil}, {"insert", 0, true, nil}, {"remove", -1, true, nil},
			{"remove", 0, true, nil}, {"insert", -1, true, nil}, {"insert", 7, true, nil},
			{"random", 0, false, []int{-1, 7}},
		}},
	}
	const calls = 1000
	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				s := impl.new()
				for i, o := range tt.ops {
					switch o.name {
					case "insert":
						if got := s.Insert(o.val); got != o.want {
							t.Errorf("op#%d Insert(%d) = %v, want %v", i, o.val, got, o.want)
						}
					case "remove":
						if got := s.Remove(o.val); got != o.want {
							t.Errorf("op#%d Remove(%d) = %v, want %v", i, o.val, got, o.want)
						}
					case "random":
						count := map[int]int{}
						for _, m := range o.members {
							count[m] = 0
						}
						for c := 0; c < calls; c++ {
							got := s.GetRandom()
							if _, ok := count[got]; !ok {
								t.Fatalf("op#%d GetRandom() = %d, not in %v", i, got, o.members)
							}
							count[got]++
						}
						for m, c := range count {
							if c == 0 {
								t.Errorf("op#%d GetRandom() never returned %d in %d calls", i, m, calls)
							}
						}
					}
				}
			})
		}
	}
}
