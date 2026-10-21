package guessnumber

import "testing"

var solutions = map[string]func(int) int{
	"guessNumber":            guessNumber,
	"guessNumberEnumeration": guessNumberEnumeration,
}

// setPick 设置 guess 接口，使其按 pick 返回提示。
func setPick(pick int) {
	guess = func(num int) int {
		switch {
		case num > pick:
			return -1
		case num < pick:
			return 1
		default:
			return 0
		}
	}
}

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		name    string
		n, pick int
	}{
		{"示例1", 10, 6},
		{"示例2-单个数", 1, 1},
		{"示例3", 2, 1},
		{"最大", 2, 2},
		{"中间", 100, 50},
		{"末尾", 100, 100},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				setPick(tt.pick)
				if got := fn(tt.n); got != tt.pick {
					t.Errorf("%s(n=%d, pick=%d) = %d", name, tt.n, tt.pick, got)
				}
			})
		}
	}
}

// TestGuessNumberLarge 大 n 只跑二分，验证不溢出。
func TestGuessNumberLarge(t *testing.T) {
	for _, pick := range []int{1, 1702766719, 2147483647} {
		setPick(pick)
		if got := guessNumber(2147483647); got != pick {
			t.Errorf("guessNumber(pick=%d) = %d", pick, got)
		}
	}
}
