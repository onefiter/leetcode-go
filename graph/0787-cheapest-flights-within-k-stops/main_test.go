package cheapestflights

import "testing"

var solutions = map[string]func(int, [][]int, int, int, int) int{
	"findCheapestPrice":    findCheapestPrice,
	"findCheapestPriceBFS": findCheapestPriceBFS,
}

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		{"示例1", 4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1, 700},
		{"示例2", 3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200},
		{"示例3-不允许中转", 3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0, 500},
		{"不可达", 3, [][]int{{0, 1, 100}}, 0, 2, 1, -1},
		{"中转次数不够", 4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1, -1},
		{"更便宜路线受中转限制", 5, [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}, 0, 2, 2, 7},
		{"最便宜路线超出中转数", 4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, tt.flights, tt.src, tt.dst, tt.k); got != tt.want {
					t.Errorf("%s(n=%d, %v, src=%d, dst=%d, k=%d) = %d, want %d",
						name, tt.n, tt.flights, tt.src, tt.dst, tt.k, got, tt.want)
				}
			})
		}
	}
}
