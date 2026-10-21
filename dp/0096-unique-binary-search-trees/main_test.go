package numtrees

import "testing"

var solutions = map[string]func(int) int{
	"numTrees":        numTrees,
	"numTreesCatalan": numTreesCatalan,
}

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"示例1", 3, 5},
		{"示例2", 1, 1},
		{"两个节点", 2, 2},
		{"四个节点", 4, 14},
		{"五个节点", 5, 42},
		{"十个节点", 10, 16796},
		{"最大值19", 19, 1767263190},
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
