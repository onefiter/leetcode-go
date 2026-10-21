package kdistinct

import "testing"

var solutions = map[string]func(string, int) int{
	"lengthOfLongestSubstringKDistinct":            lengthOfLongestSubstringKDistinct,
	"lengthOfLongestSubstringKDistinctEnumeration": lengthOfLongestSubstringKDistinctEnumeration,
}

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"示例1", "eceba", 2, 3},
		{"示例2", "aa", 1, 2},
		{"空串", "", 3, 0},
		{"k=0", "abc", 0, 0},
		{"单字符", "a", 1, 1},
		{"全相同", "bbbb", 1, 4},
		{"k大于种类数", "abcab", 10, 5},
		{"k=1", "aabbbcc", 1, 3},
		{"k=3", "abcadcacacaca", 3, 11},
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
