package worddict

import "testing"

// dictionary 两种实现共用的接口，同一张操作表跑两遍。
type dictionary interface {
	AddWord(word string)
	Search(word string) bool
}

// op 一次操作：name 为 "add" 或 "search"；search 时 want 为期望返回值。
type op struct {
	name string
	arg  string
	want bool
}

func TestWordDictionary(t *testing.T) {
	impls := []struct {
		name string
		new  func() dictionary
	}{
		{"前缀树", func() dictionary { d := Constructor(); return &d }},
		{"暴力比对", func() dictionary { d := NewWordDictionaryBrute(); return &d }},
	}
	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"add", "bad", false}, {"add", "dad", false}, {"add", "mad", false},
			{"search", "pad", false}, {"search", "bad", true},
			{"search", ".ad", true}, {"search", "b..", true},
		}},
		{"空字典", []op{
			{"search", "a", false}, {"search", ".", false},
		}},
		{"全是通配符", []op{
			{"add", "abc", false}, {"search", "...", true},
			{"search", "..", false}, {"search", "....", false},
		}},
		{"前缀不算完整单词", []op{
			{"add", "abcd", false}, {"search", "abc", false}, {"search", "ab.", false},
			{"search", "abc.", true},
		}},
		{"通配符需回溯", []op{
			{"add", "ab", false}, {"add", "cd", false},
			{"search", ".d", true}, {"search", ".b", true}, {"search", ".c", false},
		}},
		{"单字母", []op{
			{"add", "a", false}, {"add", "a", false},
			{"search", ".", true}, {"search", "a", true}, {"search", "aa", false},
			{"search", ".a", false},
		}},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				d := impl.new()
				for i, o := range tt.ops {
					switch o.name {
					case "add":
						d.AddWord(o.arg)
					case "search":
						if got := d.Search(o.arg); got != o.want {
							t.Errorf("op#%d Search(%q) = %v, want %v", i, o.arg, got, o.want)
						}
					}
				}
			})
		}
	}
}
