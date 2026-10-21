package foursumii

import "testing"

var solutions = map[string]func([]int, []int, []int, []int) int{
	"fourSumCount":      fourSumCount,
	"fourSumCountBrute": fourSumCountBrute,
}

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		name           string
		n1, n2, n3, n4 []int
		want           int
	}{
		{"示例1", []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{"示例2", []int{0}, []int{0}, []int{0}, []int{0}, 1},
		{"无解", []int{1}, []int{1}, []int{1}, []int{1}, 0},
		{"全零多元素", []int{0, 0}, []int{0, 0}, []int{0, 0}, []int{0, 0}, 16},
		{"含负数组合", []int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, 1}, 8},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n1, tt.n2, tt.n3, tt.n4); got != tt.want {
					t.Errorf("%s(%v,%v,%v,%v) = %d, want %d", name, tt.n1, tt.n2, tt.n3, tt.n4, got, tt.want)
				}
			})
		}
	}
}
