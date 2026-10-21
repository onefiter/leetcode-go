package islands

import "testing"

// toGrid 把字符串数组转成 [][]byte，每次调用都生成新网格（解法会原地修改）。
func toGrid(rows []string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, r := range rows {
		grid[i] = []byte(r)
	}
	return grid
}

var solutions = map[string]func([][]byte) int{
	"numIslands":          numIslands,
	"numIslandsUnionFind": numIslandsUnionFind,
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid []string
		want int
	}{
		{"示例1-一座大岛", []string{"11110", "11010", "11000", "00000"}, 1},
		{"示例2-三座岛", []string{"11000", "11000", "00100", "00011"}, 3},
		{"全是水", []string{"000", "000"}, 0},
		{"全是陆地", []string{"111", "111"}, 1},
		{"对角不相连", []string{"101", "010", "101"}, 5},
		{"单格陆地", []string{"1"}, 1},
		{"U 形岛", []string{"101", "101", "111"}, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(toGrid(tt.grid)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.grid, got, tt.want)
				}
			})
		}
	}
}
