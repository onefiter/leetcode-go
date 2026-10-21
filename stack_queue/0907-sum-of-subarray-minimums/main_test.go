package sumsubarraymins

import (
	"math/rand"
	"testing"
)

var solutions = map[string]func([]int) int{
	"sumSubarrayMins":            sumSubarrayMins,
	"sumSubarrayMinsTwoPass":     sumSubarrayMinsTwoPass,
	"sumSubarrayMinsEnumeration": sumSubarrayMinsEnumeration,
}

func TestSumSubarrayMins(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"示例1", []int{3, 1, 2, 4}, 17},
		{"示例2", []int{11, 81, 94, 43, 3}, 444},
		{"空数组", []int{}, 0},
		{"单元素", []int{7}, 7},
		{"全相同", []int{2, 2, 2}, 12},
		{"严格递增", []int{1, 2, 3}, 10},
		{"严格递减", []int{3, 2, 1}, 10},
		{"大值全相同", []int{30000, 30000, 30000, 30000, 30000, 30000, 30000, 30000, 30000, 30000}, 1650000},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.arr); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.arr, got, tt.want)
				}
			})
		}
	}
}

// TestLargeMod 大数组时结果会超过 1e9+7 需要取模，三种解法结果应一致。
func TestLargeMod(t *testing.T) {
	r := rand.New(rand.NewSource(907))
	arr := make([]int, 3000)
	for i := range arr {
		arr[i] = 1 + r.Intn(30000)
	}
	want := sumSubarrayMinsEnumeration(arr)
	for name, fn := range solutions {
		if got := fn(arr); got != want {
			t.Fatalf("%s = %d, want %d", name, got, want)
		}
	}
}

// TestAgainstEnumeration 随机小数组对拍（含大量重复值）。
func TestAgainstEnumeration(t *testing.T) {
	r := rand.New(rand.NewSource(9070))
	for k := 0; k < 500; k++ {
		arr := make([]int, r.Intn(25))
		for i := range arr {
			arr[i] = 1 + r.Intn(5)
		}
		want := sumSubarrayMinsEnumeration(arr)
		for name, fn := range solutions {
			if got := fn(arr); got != want {
				t.Fatalf("%s(%v) = %d, want %d", name, arr, got, want)
			}
		}
	}
}
