package decodeways

import "testing"

var solutions = map[string]func(string) int{
	"numDecodings":     numDecodings,
	"numDecodingsMemo": numDecodingsMemo,
}

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "12", 2},
		{"示例2", "226", 3},
		{"示例3前导零", "06", 0},
		{"单个零", "0", 0},
		{"含10", "10", 1},
		{"连续零", "100", 0},
		{"超过26", "27", 1},
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
