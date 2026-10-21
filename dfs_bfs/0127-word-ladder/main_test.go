package wordladder

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string, string, []string) int{
	"ladderLength":    ladderLength,
	"ladderLengthBFS": ladderLengthBFS,
}

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name     string
		begin    string
		end      string
		wordList []string
		want     int
	}{
		{"示例1", "hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"示例2-终点不在字典", "hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		{"一步到达", "a", "c", []string{"a", "b", "c"}, 2},
		{"不可达", "hot", "dog", []string{"hot", "dog"}, 0},
		{"字典含起点", "hot", "dog", []string{"hot", "dog", "dot"}, 3},
		{"多条路径取最短", "abc", "xyz", []string{"xbc", "xyc", "xyz", "abz", "ayz"}, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.begin, tt.end, slices.Clone(tt.wordList))
				if got != tt.want {
					t.Errorf("%s(%q, %q, %v) = %d, want %d", name, tt.begin, tt.end, tt.wordList, got, tt.want)
				}
			})
		}
	}
}
