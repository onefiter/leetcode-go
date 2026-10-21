package charreplace

import "testing"

var solutions = map[string]func(string, int) int{
	"characterReplacement":            characterReplacement,
	"characterReplacementEnumeration": characterReplacementEnumeration,
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"示例1", "ABAB", 2, 4},
		{"示例2", "AABABBA", 1, 4},
		{"空串", "", 2, 0},
		{"单字符", "Z", 0, 1},
		{"全相同", "AAAA", 0, 4},
		{"k=0", "AABBBCC", 0, 3},
		{"k足够大", "ABCDE", 10, 5},
		// 整串有 7 个 B、3 个非 B，k=2 不够；"BBBBACBBB" 替换 A、C 得到 9。
		{"中间替换", "ABBBBACBBB", 2, 9},
		{"全不同k=1", "ABCD", 1, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s, tt.k); got != tt.want {
					t.Errorf("%s(%q, %d) = %d, want %d", name, tt.s, tt.k, got, tt.want)
				}
			})
		}
	}
}
