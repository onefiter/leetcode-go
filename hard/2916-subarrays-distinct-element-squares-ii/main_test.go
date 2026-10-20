package sumcounts

import "testing"

// bruteForce 暴力解，用于对拍。
func bruteForce(nums []int) int {
	ans := 0
	for i := range nums {
		set := map[int]bool{}
		for j := i; j < len(nums); j++ {
			set[nums[j]] = true
			ans = (ans + len(set)*len(set)) % mod
		}
	}
	return ans
}

func TestSumCounts(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1}, 15},
		{[]int{2, 2}, 3},
		{[]int{5}, 1},
	}
	for _, tt := range tests {
		if got := sumCounts(tt.nums); got != tt.want {
			t.Errorf("sumCounts(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}

	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4}
	if got, want := sumCounts(nums), bruteForce(nums); got != want {
		t.Errorf("sumCounts(%v) = %d, bruteForce = %d", nums, got, want)
	}
}
