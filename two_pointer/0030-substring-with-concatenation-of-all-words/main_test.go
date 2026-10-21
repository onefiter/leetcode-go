package substrconcat

import (
	"slices"
	"strings"
	"testing"
)

var solutions = map[string]func(string, []string) []int{
	"findSubstring":            findSubstring,
	"findSubstringEnumeration": findSubstringEnumeration,
}

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		{"示例1", "barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"示例2-无解", "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"示例3", "barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"空串", "", []string{"a"}, []int{}},
		{"单字符单单词", "a", []string{"a"}, []int{0}},
		{"s比串联短", "ab", []string{"ab", "cd"}, []int{}},
		{"全相同字符", "aaaaaa", []string{"aa", "aa"}, []int{0, 1, 2}},
		{"重复单词", "wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"超量回缩", "foofoobarfoo", []string{"foo", "bar"}, []int{3, 6}},
		{"单单词多处", "abcabc", []string{"bc"}, []int{1, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.s, slices.Clone(tt.words))
				slices.Sort(got) // 下标顺序不唯一，排序后比较
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%q, %v) = %v, want %v", name, tt.s, tt.words, got, tt.want)
				}
			})
		}
	}
}

// TestFindSubstringAgainstEnumeration 用暴力解核对滑窗解（字母表很小，易出现重叠匹配）。
func TestFindSubstringAgainstEnumeration(t *testing.T) {
	s := strings.Repeat("abba", 6) + "ab"
	wordSets := [][]string{
		{"ab", "ba"}, {"ab", "ab"}, {"bb", "aa"}, {"a", "b", "b"}, {"abb", "aab"}, {"ba"},
	}
	for _, words := range wordSets {
		want := findSubstringEnumeration(s, words)
		got := findSubstring(s, words)
		slices.Sort(want)
		slices.Sort(got)
		if !slices.Equal(got, want) {
			t.Errorf("findSubstring(%q, %v) = %v, want %v", s, words, got, want)
		}
	}
}
