package palindrome

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]string 规范化：外层按字典序排序（内层是切分顺序，不能排序）。
func normalize(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, g := range groups {
		out[i] = slices.Clone(g)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return slices.Compare(a, b)
	})
	return out
}

var solutions = map[string]func(string) [][]string{
	"partition":            partition,
	"partitionEnumeration": partitionEnumeration,
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{"示例1", "aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"示例2-单字符", "a", [][]string{{"a"}}},
		{"无多字符回文", "abc", [][]string{{"a", "b", "c"}}},
		{"整体回文", "aba", [][]string{{"a", "b", "a"}, {"aba"}}},
		{"全相同", "aaa", [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}}},
		{"嵌套回文", "abba", [][]string{{"a", "b", "b", "a"}, {"a", "bb", "a"}, {"abba"}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(tt.s))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%q) = %v, want %v", name, tt.s, got, want)
				}
			})
		}
	}
}
