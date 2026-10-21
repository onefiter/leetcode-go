package wordbreak

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string, []string) bool{
	"wordBreak":    wordBreak,
	"wordBreakBFS": wordBreakBFS,
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{"示例1", "leetcode", []string{"leet", "code"}, true},
		{"示例2单词复用", "applepenapple", []string{"apple", "pen"}, true},
		{"示例3", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"单字符匹配", "a", []string{"a"}, true},
		{"单字符不匹配", "a", []string{"b"}, false},
		{"需要回溯", "aaaaaaab", []string{"a", "aa", "aaa"}, false},
		{"多种拆分", "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s, slices.Clone(tt.wordDict)); got != tt.want {
					t.Errorf("%s(%q, %v) = %v, want %v", name, tt.s, tt.wordDict, got, tt.want)
				}
			})
		}
	}
}
