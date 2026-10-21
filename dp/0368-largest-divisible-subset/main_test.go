package divisible

import (
	"slices"
	"testing"
)

// valid 校验 got 是 nums 的子集（无重复）且任意两元素之间都满足整除关系。
func valid(nums, got []int) bool {
	seen := map[int]bool{}
	for _, x := range got {
		if seen[x] || !slices.Contains(nums, x) {
			return false
		}
		seen[x] = true
	}
	for i := 0; i < len(got); i++ {
		for j := i + 1; j < len(got); j++ {
			if got[i]%got[j] != 0 && got[j]%got[i] != 0 {
				return false
			}
		}
	}
	return true
}

var solutions = map[string]func([]int) []int{
	"largestDivisibleSubset":    largestDivisibleSubset,
	"largestDivisibleSubsetDFS": largestDivisibleSubsetDFS,
}

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantLen int
	}{
		{"示例1", []int{1, 2, 3}, 2},
		{"示例2", []int{1, 2, 4, 8}, 4},
		{"单元素", []int{7}, 1},
		{"两两互质", []int{2, 3, 5, 7}, 1},
		{"乱序输入", []int{8, 1, 4, 2, 3, 9}, 4},
		{"多条链取最长", []int{3, 4, 16, 8, 9, 27, 81}, 4},
		{"不含1", []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.nums))
				if len(got) != tt.wantLen || !valid(tt.nums, got) {
					t.Errorf("%s(%v) = %v, want valid subset of len %d", name, tt.nums, got, tt.wantLen)
				}
			})
		}
	}
}
