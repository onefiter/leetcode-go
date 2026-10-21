package nearestexit

import (
	"slices"
	"testing"
)

// toMaze 把字符串切片转成 [][]byte，每次调用都会生成新的副本（解法一会原地修改迷宫）。
func toMaze(rows []string) [][]byte {
	maze := make([][]byte, len(rows))
	for i, r := range rows {
		maze[i] = []byte(r)
	}
	return maze
}

var solutions = map[string]func([][]byte, []int) int{
	"nearestExit":          nearestExit,
	"nearestExitFromExits": nearestExitFromExits,
}

func TestNearestExit(t *testing.T) {
	tests := []struct {
		name     string
		maze     []string
		entrance []int
		want     int
	}{
		{"示例1", []string{"++.+", "...+", "+++."}, []int{1, 2}, 1},
		{"示例2", []string{"+++", "...", "+++"}, []int{1, 0}, 2},
		{"示例3-单格无出口", []string{".+"}, []int{0, 0}, -1},
		{"被墙围死", []string{"+++", "+.+", "+++"}, []int{1, 1}, -1},
		{"入口旁边就是出口", []string{".."}, []int{0, 0}, 1},
		{"迷宫中心", []string{".....", ".....", ".....", ".....", "....."}, []int{2, 2}, 2},
		{"绕路出迷宫", []string{"+++++", "+...+", "+.+.+", "+.+.+", "+++.+"}, []int{2, 1}, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(toMaze(tt.maze), slices.Clone(tt.entrance))
				if got != tt.want {
					t.Errorf("%s(%q, %v) = %d, want %d", name, tt.maze, tt.entrance, got, tt.want)
				}
			})
		}
	}
}
