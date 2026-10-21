package findanagrams

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string, string) []int{
	"findAnagrams":      findAnagrams,
	"findAnagramsBrute": findAnagramsBrute,
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{"示例1", "cbaebabacd", "abc", []int{0, 6}},
		{"示例2", "abab", "ab", []int{0, 1, 2}},
		{"无解", "abcdef", "xyz", []int{}},
		{"空s", "", "a", []int{}},
		{"空p", "abc", "", []int{}},
		{"p比s长", "ab", "abc", []int{}},
		{"全相同", "aaaa", "aa", []int{0, 1, 2}},
		{"整串异位", "bca", "abc", []int{0}},
		{"单字符多次", "aaa", "a", []int{0, 1, 2}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s, tt.p); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%q, %q) = %v, want %v", name, tt.s, tt.p, got, tt.want)
				}
			})
		}
	}
}
