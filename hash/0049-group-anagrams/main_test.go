package groupanagrams

import (
	"slices"
	"sort"
	"testing"
)

// normalize 把 [][]string 规范化：每组内部排序，组按首元素排序，便于顺序无关比较。
func normalize(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, g := range groups {
		cp := slices.Clone(g)
		slices.Sort(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
	return out
}

func equalGroups(a, b [][]string) bool {
	na, nb := normalize(a), normalize(b)
	if len(na) != len(nb) {
		return false
	}
	for i := range na {
		if !slices.Equal(na[i], nb[i]) {
			return false
		}
	}
	return true
}

var solutions = map[string]func([]string) [][]string{
	"groupAnagrams":           groupAnagrams,
	"groupAnagramsCount":      groupAnagramsCount,
	"groupAnagramsSortStable": groupAnagramsSortStable,
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"示例1", []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"空串单元素", []string{""}, [][]string{{""}}},
		{"单字符", []string{"a"}, [][]string{{"a"}}},
		{"全是异位词", []string{"abc", "bca", "cab"}, [][]string{{"abc", "bca", "cab"}}},
		{"无异位词", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{"含空串与非空", []string{"", "b", ""}, [][]string{{"", ""}, {"b"}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.strs); !equalGroups(got, tt.want) {
					t.Errorf("%s(%v) = %v, want (顺序无关) %v", name, tt.strs, got, tt.want)
				}
			})
		}
	}
}

func TestGroupAnagramsEmptyInput(t *testing.T) {
	for name, fn := range solutions {
		t.Run(name, func(t *testing.T) {
			if got := fn([]string{}); len(got) != 0 {
				t.Errorf("%s([]) = %v, want 空结果", name, got)
			}
		})
	}
}
