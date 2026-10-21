package longestcommonprefix

import "testing"

var solutions = map[string]func([]string) string{
	"longestCommonPrefix":  longestCommonPrefix,
	"longestCommonPrefixH": longestCommonPrefixH,
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"示例1", []string{"flower", "flow", "flight"}, "fl"},
		{"示例2-无公共前缀", []string{"dog", "racecar", "car"}, ""},
		{"空数组", []string{}, ""},
		{"单个字符串", []string{"alone"}, "alone"},
		{"含空字符串", []string{"", "abc"}, ""},
		{"完全相同", []string{"abc", "abc", "abc"}, "abc"},
		{"一个是另一个前缀", []string{"ab", "abc"}, "ab"},
		{"首字符就不同", []string{"a", "b"}, ""},
		{"全部单字符相同", []string{"c", "c"}, "c"},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.strs); got != tt.want {
					t.Errorf("%s(%v) = %q, want %q", name, tt.strs, got, tt.want)
				}
			})
		}
	}
}
