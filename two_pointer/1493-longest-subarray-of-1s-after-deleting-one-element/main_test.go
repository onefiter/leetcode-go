package deleteone

import "testing"

var solutions = map[string]func([]int) int{
	"longestSubarray":            longestSubarray,
	"longestSubarrayDP":          longestSubarrayDP,
	"longestSubarrayEnumeration": longestSubarrayEnumeration,
}

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{1, 1, 0, 1}, 3},
		{"示例2", []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{"示例3-全1必须删一个", []int{1, 1, 1}, 2},
		{"空数组", []int{}, 0},
		{"单个1", []int{1}, 0},
		{"单个0", []int{0}, 0},
		{"全为0", []int{0, 0, 0}, 0},
		{"两个0相邻", []int{1, 0, 0, 1}, 1},
		{"0在开头", []int{0, 1, 1}, 2},
		{"0在末尾", []int{1, 1, 1, 0}, 3},
		{"多段取最优拼接", []int{1, 0, 1, 1, 0, 1, 1, 1}, 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
