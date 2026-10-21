package treediameter

import "testing"

var solutions = map[string]func([][]int) int{
	"treeDiameter":            treeDiameter,
	"treeDiameterDP":          treeDiameterDP,
	"treeDiameterEnumeration": treeDiameterEnumeration,
}

func TestTreeDiameter(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  int
	}{
		{"示例1", [][]int{{0, 1}, {0, 2}}, 2},
		{"示例2", [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 5}}, 4},
		{"单节点", [][]int{}, 0},
		{"两个节点", [][]int{{0, 1}}, 1},
		{"一条链", [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 4},
		{"星形", [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, 2},
		{"直径不经过节点0", [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {4, 5}}, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.edges); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.edges, got, tt.want)
				}
			})
		}
	}
}

// TestAgainstEnumeration 用确定性随机树对拍三种解法。
func TestAgainstEnumeration(t *testing.T) {
	seed := 42
	for n := 1; n <= 60; n++ {
		edges := make([][]int, 0, n-1)
		for v := 1; v < n; v++ {
			seed = (seed*1103515245 + 12345) % 2147483648
			edges = append(edges, []int{seed % v, v}) // 连到之前任意一个节点，保证是树
		}
		want := treeDiameterEnumeration(edges)
		for name, fn := range solutions {
			if got := fn(edges); got != want {
				t.Fatalf("%s n=%d edges=%v: got %d, want %d", name, n, edges, got, want)
			}
		}
	}
}
