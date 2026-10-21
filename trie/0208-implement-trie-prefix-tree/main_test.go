package trie

import "testing"

// trier 两种实现共用的接口，同一张操作表跑两遍。
type trier interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

// op 一次操作：name 为 "insert"、"search" 或 "startsWith"；查询时 want 为期望返回值。
type op struct {
	name string
	arg  string
	want bool
}

func TestTrie(t *testing.T) {
	impls := []struct {
		name string
		new  func() trier
	}{
		{"数组Trie", func() trier { tr := Constructor(); return &tr }},
		{"mapTrie", func() trier { tr := NewTrieMap(); return &tr }},
	}
	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"insert", "apple", false}, {"search", "apple", true}, {"search", "app", false},
			{"startsWith", "app", true}, {"insert", "app", false}, {"search", "app", true},
		}},
		{"空树查询", []op{
			{"search", "a", false}, {"startsWith", "a", false},
		}},
		{"空前缀", []op{
			{"startsWith", "", true}, {"insert", "abc", false}, {"startsWith", "", true},
		}},
		{"单词比插入的长", []op{
			{"insert", "ab", false}, {"search", "abc", false}, {"startsWith", "abc", false},
		}},
		{"分叉路径", []op{
			{"insert", "car", false}, {"insert", "cat", false}, {"insert", "dog", false},
			{"search", "ca", false}, {"startsWith", "ca", true}, {"search", "cat", true},
			{"startsWith", "do", true}, {"startsWith", "dot", false},
		}},
		{"重复插入", []op{
			{"insert", "x", false}, {"insert", "x", false}, {"search", "x", true},
			{"search", "xx", false},
		}},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				tr := impl.new()
				for i, o := range tt.ops {
					switch o.name {
					case "insert":
						tr.Insert(o.arg)
					case "search":
						if got := tr.Search(o.arg); got != o.want {
							t.Errorf("op#%d Search(%q) = %v, want %v", i, o.arg, got, o.want)
						}
					case "startsWith":
						if got := tr.StartsWith(o.arg); got != o.want {
							t.Errorf("op#%d StartsWith(%q) = %v, want %v", i, o.arg, got, o.want)
						}
					}
				}
			})
		}
	}
}
