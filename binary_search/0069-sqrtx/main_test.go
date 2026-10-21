package sqrtx

import "testing"

var solutions = map[string]func(int) int{
	"mySqrt":            mySqrt,
	"mySqrtNewton":      mySqrtNewton,
	"mySqrtEnumeration": mySqrtEnumeration,
}

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"示例1", 4, 2},
		{"示例2", 8, 2},
		{"零", 0, 0},
		{"一", 1, 1},
		{"二", 2, 1},
		{"三", 3, 1},
		{"完全平方数", 49, 7},
		{"完全平方数减一", 48, 6},
		{"最大值", 2147483647, 46340},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.x); got != tt.want {
					t.Errorf("%s(%d) = %d, want %d", name, tt.x, got, tt.want)
				}
			})
		}
	}
}

// TestMySqrtRange 在一段连续区间上校验 r*r <= x < (r+1)*(r+1)。
func TestMySqrtRange(t *testing.T) {
	for x := 0; x <= 2000; x++ {
		for name, fn := range solutions {
			r := fn(x)
			if r*r > x || (r+1)*(r+1) <= x {
				t.Errorf("%s(%d) = %d, not floor sqrt", name, x, r)
			}
		}
	}
}
