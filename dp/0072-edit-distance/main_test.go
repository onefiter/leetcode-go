package editdistance

import "testing"

var solutions = map[string]func(string, string) int{
	"minDistance":     minDistance,
	"minDistanceMemo": minDistanceMemo,
}

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name         string
		word1, word2 string
		want         int
	}{
		{"示例1", "horse", "ros", 3},
		{"示例2", "intention", "execution", 5},
		{"都为空", "", "", 0},
		{"源串为空", "", "abc", 3},
		{"目标串为空", "abc", "", 3},
		{"完全相同", "same", "same", 0},
		{"单字符替换", "a", "b", 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.word1, tt.word2); got != tt.want {
					t.Errorf("%s(%q, %q) = %v, want %v", name, tt.word1, tt.word2, got, tt.want)
				}
			})
		}
	}
}
