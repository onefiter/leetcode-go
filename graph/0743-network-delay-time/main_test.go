package networkdelay

import "testing"

var solutions = map[string]func([][]int, int, int) int{
	"networkDelayTime":            networkDelayTime,
	"networkDelayTimeBellmanFord": networkDelayTimeBellmanFord,
}

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][]int
		n     int
		k     int
		want  int
	}{
		{"示例1", [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{"示例2", [][]int{{1, 2, 1}}, 2, 1, 1},
		{"示例3-反向不可达", [][]int{{1, 2, 1}}, 2, 2, -1},
		{"单节点", [][]int{}, 1, 1, 0},
		{"绕路更短", [][]int{{1, 2, 10}, {1, 3, 1}, {3, 2, 2}}, 3, 1, 3},
		{"有孤立节点", [][]int{{1, 2, 1}, {2, 3, 1}}, 4, 1, -1},
		{"含环与重边", [][]int{{1, 2, 5}, {1, 2, 2}, {2, 3, 2}, {3, 1, 1}, {3, 4, 4}}, 4, 1, 8},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.times, tt.n, tt.k); got != tt.want {
					t.Errorf("%s(%v, n=%d, k=%d) = %d, want %d", name, tt.times, tt.n, tt.k, got, tt.want)
				}
			})
		}
	}
}
