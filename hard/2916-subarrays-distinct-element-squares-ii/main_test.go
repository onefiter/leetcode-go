package sumcounts

import "testing"

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
		if got := sumCountsEnumeration(tt.nums); got != tt.want {
			t.Errorf("sumCountsEnumeration(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}

	// 树状数组解法与穷举解法对拍。
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4}
	if got, want := sumCounts(nums), sumCountsEnumeration(nums); got != want {
		t.Errorf("sumCounts(%v) = %d, sumCountsEnumeration = %d", nums, got, want)
	}
}
