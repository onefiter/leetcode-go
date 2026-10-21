package wallsgates

import (
	"reflect"
	"slices"
	"testing"
)

// cloneGrid 深拷贝二维切片，避免原地修改污染测试数据。
func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i := range g {
		out[i] = slices.Clone(g[i])
	}
	return out
}

var solutions = map[string]func([][]int){
	"wallsAndGates":        wallsAndGates,
	"wallsAndGatesPerGate": wallsAndGatesPerGate,
}

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  [][]int
	}{
		{"示例1",
			[][]int{{INF, -1, 0, INF}, {INF, INF, INF, -1}, {INF, -1, INF, -1}, {0, -1, INF, INF}},
			[][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
		{"示例2-单墙", [][]int{{-1}}, [][]int{{-1}}},
		{"单个空房间无门", [][]int{{INF}}, [][]int{{INF}}},
		{"被墙隔开不可达", [][]int{{0, -1, INF}}, [][]int{{0, -1, INF}}},
		{"一行两个门", [][]int{{0, INF, INF, INF, 0}}, [][]int{{0, 1, 2, 1, 0}}},
		{"全是门", [][]int{{0, 0}, {0, 0}}, [][]int{{0, 0}, {0, 0}}},
		{"绕墙而行", [][]int{{0, -1, INF}, {INF, -1, INF}, {INF, INF, INF}},
			[][]int{{0, -1, 6}, {1, -1, 5}, {2, 3, 4}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				rooms := cloneGrid(tt.rooms)
				fn(rooms)
				if !reflect.DeepEqual(rooms, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.rooms, rooms, tt.want)
				}
			})
		}
	}
}
