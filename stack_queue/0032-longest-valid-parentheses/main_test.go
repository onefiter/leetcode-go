package longestvalid

import "testing"

var solutions = map[string]func(string) int{
	"longestValidParentheses":            longestValidParentheses,
	"longestValidParenthesesDP":          longestValidParenthesesDP,
	"longestValidParenthesesCount":       longestValidParenthesesCount,
	"longestValidParenthesesEnumeration": longestValidParenthesesEnumeration,
}

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "(()", 2},
		{"示例2", ")()())", 4},
		{"示例3-空串", "", 0},
		{"单个左括号", "(", 0},
		{"单个右括号", ")", 0},
		{"一对", "()", 2},
		{"全是左括号", "((((", 0},
		{"全是右括号", "))))", 0},
		{"嵌套加并列", "()(())", 6},
		{"中间断开", "()(()", 2},
		{"整串有效", "(()())", 6},
		{"右侧更长", "())((()))", 6},
		{"反序", ")(", 0},
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

// TestAgainstEnumeration 枚举长度 <= 10 的所有括号串，用暴力结果核对其余解法。
func TestAgainstEnumeration(t *testing.T) {
	for n := 0; n <= 10; n++ {
		for mask := 0; mask < 1<<n; mask++ {
			b := make([]byte, n)
			for i := 0; i < n; i++ {
				if mask>>i&1 == 1 {
					b[i] = '('
				} else {
					b[i] = ')'
				}
			}
			s := string(b)
			want := longestValidParenthesesEnumeration(s)
			for name, fn := range solutions {
				if got := fn(s); got != want {
					t.Fatalf("%s(%q) = %d, want %d", name, s, got, want)
				}
			}
		}
	}
}
