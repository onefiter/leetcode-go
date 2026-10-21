package uniquepaths

import "testing"

var solutions = map[string]func(int, int) int{
	"uniquePaths":     uniquePaths,
	"uniquePathsComb": uniquePathsComb,
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m, n int
		want int
	}{
		{"示例1", 3, 7, 28},
		{"示例2", 3, 2, 3},
		{"示例3", 7, 3, 28},
		{"单格", 1, 1, 1},
		{"单行", 1, 10, 1},
		{"方阵", 3, 3, 6},
		{"较大网格", 23, 12, 193536720},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.m, tt.n); got != tt.want {
					t.Errorf("%s(%d, %d) = %d, want %d", name, tt.m, tt.n, got, tt.want)
				}
			})
		}
	}
}
