package findthewinner

import "testing"

var solutions = map[string]func(int, int) int{
	"findTheWinner":            findTheWinner,
	"findTheWinnerEnumeration": findTheWinnerEnumeration,
}

func TestFindTheWinner(t *testing.T) {
	tests := []struct {
		name string
		n, k int
		want int
	}{
		{"示例1", 5, 2, 3},
		{"示例2", 6, 5, 1},
		{"单人", 1, 1, 1},
		{"单人k大", 1, 7, 1},
		{"k=1最后一人获胜", 4, 1, 4},
		{"两人k=2", 2, 2, 1},
		{"k大于n", 3, 5, 1},
		{"k等于n", 7, 7, 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, tt.k); got != tt.want {
					t.Errorf("%s(%d, %d) = %d, want %d", name, tt.n, tt.k, got, tt.want)
				}
			})
		}
	}
}

// TestFindTheWinnerAgainstEnumeration 在小范围内枚举所有 (n, k) 对拍。
func TestFindTheWinnerAgainstEnumeration(t *testing.T) {
	for n := 1; n <= 30; n++ {
		for k := 1; k <= 30; k++ {
			if got, want := findTheWinner(n, k), findTheWinnerEnumeration(n, k); got != want {
				t.Fatalf("findTheWinner(%d, %d) = %d, want %d", n, k, got, want)
			}
		}
	}
}
