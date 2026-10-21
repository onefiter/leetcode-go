package nqueens2

import "testing"

var solutions = map[string]func(int) int{
	"totalNQueens":            totalNQueens,
	"totalNQueensEnumeration": totalNQueensEnumeration,
}

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"示例1-n=4", 4, 2},
		{"示例2-n=1", 1, 1},
		{"n=2 无解", 2, 0},
		{"n=3 无解", 3, 0},
		{"n=6 解比 n=5 少", 6, 4},
		{"n=8 经典八皇后", 8, 92},
		{"边界-n=9", 9, 352},
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
