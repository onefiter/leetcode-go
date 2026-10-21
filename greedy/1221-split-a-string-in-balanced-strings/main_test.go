package balancedstring

import "testing"

var solutions = map[string]func(string) int{
	"balancedStringSplit":   balancedStringSplit,
	"balancedStringSplitDP": balancedStringSplitDP,
}

func TestBalancedStringSplit(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "RLRRLLRLRL", 4},
		{"示例2", "RLRRRLLRLL", 2},
		{"示例3", "LLLLRRRR", 1},
		{"最短平衡串", "LR", 1},
		{"交替出现", "RLRLRLRL", 4},
		{"嵌套结构", "RRLLRL", 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.s)
				if got != tt.want {
					t.Errorf("%s(%q) = %d, want %d", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
