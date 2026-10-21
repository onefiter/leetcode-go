package provinces

import "testing"

var solutions = map[string]func([][]int) int{
	"findCircleNum":    findCircleNum,
	"findCircleNumDFS": findCircleNumDFS,
}

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{"示例1", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{"示例2-全部孤立", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
		{"单个城市", [][]int{{1}}, 1},
		{"全部相连", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1},
		{"间接相连成链", [][]int{{1, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 1, 1}, {0, 0, 1, 1}}, 1},
		{"首尾相连中间孤立", [][]int{{1, 0, 0, 1}, {0, 1, 0, 0}, {0, 0, 1, 0}, {1, 0, 0, 1}}, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.isConnected); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.isConnected, got, tt.want)
				}
			})
		}
	}
}
