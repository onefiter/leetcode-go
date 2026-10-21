package perfectsquares

import "testing"

var solutions = map[string]func(int) int{
	"numSquares":     numSquares,
	"numSquaresMath": numSquaresMath,
}

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"示例1", 12, 3},
		{"示例2", 13, 2},
		{"最小值", 1, 1},
		{"本身是平方数", 16, 1},
		{"需要四个", 7, 4},
		{"4的幂乘7", 28, 4},
		{"较大值", 9999, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n); got != tt.want {
					t.Errorf("%s(%d) = %d, want %d", name, tt.n, got, tt.want)
				}
			})
		}
	}
}
