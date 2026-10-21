package getfood

import "testing"

// toGrid 把字符串切片转成 [][]byte，每次调用都会生成新的副本（解法一会原地修改网格）。
func toGrid(rows []string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, r := range rows {
		grid[i] = []byte(r)
	}
	return grid
}

var solutions = map[string]func([][]byte) int{
	"getFood":         getFood,
	"getFoodFromFood": getFoodFromFood,
}

func TestGetFood(t *testing.T) {
	tests := []struct {
		name string
		grid []string
		want int
	}{
		{"示例1", []string{"XXXXXX", "X*OOOX", "XOO#OX", "XXXXXX"}, 3},
		{"示例2-被围住", []string{"XXXXX", "X*XOX", "XOX#X", "XXXXX"}, -1},
		{"示例3-多个食物", []string{"XXXXXXXX", "X*OXO#OX", "XOOXOOXX", "XOOOO#OX", "XXXXXXXX"}, 6},
		{"紧邻食物", []string{"*#"}, 1},
		{"没有食物", []string{"*OO"}, -1},
		{"取最近的食物", []string{"#OO*O#"}, 2},
		{"绕过障碍", []string{"*XO", "OXO", "OO#"}, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(toGrid(tt.grid))
				if got != tt.want {
					t.Errorf("%s(%q) = %d, want %d", name, tt.grid, got, tt.want)
				}
			})
		}
	}
}
