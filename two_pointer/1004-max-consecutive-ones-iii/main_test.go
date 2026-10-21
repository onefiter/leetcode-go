package longestones

import "testing"

var solutions = map[string]func([]int, int) int{
	"longestOnes":            longestOnes,
	"longestOnesEnumeration": longestOnesEnumeration,
}

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"示例2", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{"空数组", []int{}, 1, 0},
		{"单个0-k为0", []int{0}, 0, 0},
		{"单个0-k为1", []int{0}, 1, 1},
		{"全为1", []int{1, 1, 1, 1}, 0, 4},
		{"全为0-k为0", []int{0, 0, 0}, 0, 0},
		{"全为0-k为2", []int{0, 0, 0}, 2, 2},
		{"k大于0的个数", []int{1, 0, 1, 0}, 5, 4},
		{"k为0取最长1段", []int{1, 1, 0, 1, 1, 1, 0, 1}, 0, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
