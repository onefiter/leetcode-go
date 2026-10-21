package climbstairs

import "testing"

var solutions = map[string]func(int) int{
	"climbStairs":      climbStairs,
	"climbStairsTable": climbStairsTable,
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"一阶", 1, 1},
		{"示例1", 2, 2},
		{"示例2", 3, 3},
		{"五阶", 5, 8},
		{"十阶", 10, 89},
		{"上限45阶", 45, 1836311903},
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
