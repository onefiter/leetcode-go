package firstuniqchar

import "testing"

var solutions = map[string]func(string) int{
	"firstUniqChar":      firstUniqChar,
	"firstUniqCharBrute": firstUniqCharBrute,
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "leetcode", 0},
		{"示例2", "loveleetcode", 2},
		{"示例3-无解", "aabb", -1},
		{"空串", "", -1},
		{"单字符", "z", 0},
		{"全相同", "aaaa", -1},
		{"唯一在末尾", "aabbc", 4},
		{"唯一在中间", "aabcc", 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s); got != tt.want {
					t.Errorf("%s(%q) = %d, want %d", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
