package shortestsubarray

import "testing"

var solutions = map[string]func([]int, int) int{
	"shortestSubarray":            shortestSubarray,
	"shortestSubarrayEnumeration": shortestSubarrayEnumeration,
}

func TestShortestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1-单元素满足", []int{1}, 1, 1},
		{"示例2-不存在", []int{1, 2}, 4, -1},
		{"示例3-含负数", []int{2, -1, 2}, 3, 3},
		{"单元素不满足", []int{5}, 6, -1},
		{"全相同", []int{3, 3, 3, 3}, 7, 3},
		{"答案在末尾", []int{1, 1, 1, 10}, 10, 1},
		{"答案在开头", []int{10, 1, 1, 1}, 10, 1},
		{"负数隔断", []int{84, -37, 32, 40, 95}, 167, 3},
		{"全负数", []int{-1, -2, -3}, 1, -1},
		{"整个数组", []int{1, 2, 3}, 6, 3},
		{"负数后反弹", []int{48, 99, 37, 4, -31}, 140, 2},
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

// TestShortestSubarrayAgainstEnumeration 用确定性伪随机数组对拍两种解法。
func TestShortestSubarrayAgainstEnumeration(t *testing.T) {
	seed := 12345
	next := func() int {
		seed = (seed*1103515245 + 12345) % (1 << 31)
		return seed
	}
	for round := 0; round < 300; round++ {
		n := next()%12 + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = next()%21 - 8 // [-8, 12]
		}
		k := next()%30 + 1
		want := shortestSubarrayEnumeration(nums, k)
		if got := shortestSubarray(nums, k); got != want {
			t.Fatalf("shortestSubarray(%v, %d) = %d, want %d", nums, k, got, want)
		}
	}
}
