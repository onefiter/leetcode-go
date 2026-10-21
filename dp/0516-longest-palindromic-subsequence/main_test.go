package lps

import "testing"

var solutions = map[string]func(string) int{
	"longestPalindromeSubseq":   longestPalindromeSubseq,
	"longestPalindromeSubseq2D": longestPalindromeSubseq2D,
}

func TestLongestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "bbbab", 4},
		{"示例2", "cbbd", 2},
		{"单字符", "a", 1},
		{"全不同", "abcd", 1},
		{"本身回文", "racecar", 7},
		{"两字符相同", "aa", 2},
		{"间隔回文", "agbdba", 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s); got != tt.want {
					t.Errorf("%s(%q) = %v, want %v", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
