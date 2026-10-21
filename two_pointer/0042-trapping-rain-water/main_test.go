package trapwater

import "testing"

var solutions = map[string]func([]int) int{
	"trap":               trap,
	"trapMonotonicStack": trapMonotonicStack,
	"trapPrefixMax":      trapPrefixMax,
	"trapEnumeration":    trapEnumeration,
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"示例1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"示例2", []int{4, 2, 0, 3, 2, 5}, 9},
		{"空数组", []int{}, 0},
		{"单个元素", []int{5}, 0},
		{"两个元素", []int{3, 1}, 0},
		{"全部相同", []int{2, 2, 2, 2}, 0},
		{"单调递增", []int{1, 2, 3, 4}, 0},
		{"单调递减", []int{4, 3, 2, 1}, 0},
		{"V字形", []int{3, 0, 3}, 3},
		{"全为0", []int{0, 0, 0}, 0},
		{"两侧不等高", []int{5, 1, 2, 1, 3}, 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.height); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.height, got, tt.want)
				}
			})
		}
	}
}
