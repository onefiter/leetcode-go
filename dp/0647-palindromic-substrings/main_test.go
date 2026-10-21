package palindromic

import "testing"

var solutions = map[string]func(string) int{
	"countSubstrings":   countSubstrings,
	"countSubstringsDP": countSubstringsDP,
}

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "abc", 3},
		{"示例2", "aaa", 6},
		{"单字符", "a", 1},
		{"两字符相同", "aa", 3},
		{"两字符不同", "ab", 2},
		{"奇偶混合", "abba", 6},
		{"较长串", "racecar", 10},
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
