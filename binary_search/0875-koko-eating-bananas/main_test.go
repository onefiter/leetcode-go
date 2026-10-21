package koko

import "testing"

var solutions = map[string]func([]int, int) int{
	"minEatingSpeed":            minEatingSpeed,
	"minEatingSpeedEnumeration": minEatingSpeedEnumeration,
}

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"示例1", []int{3, 6, 7, 11}, 8, 4},
		{"示例2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"示例3", []int{30, 11, 23, 4, 20}, 6, 23},
		{"单堆-一小时", []int{10}, 1, 10},
		{"单堆-时间充裕", []int{10}, 100, 1},
		{"全相同", []int{5, 5, 5, 5}, 8, 3},
		{"h等于堆数", []int{1, 2, 3}, 3, 3},
		{"全是1", []int{1, 1, 1}, 5, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.piles, tt.h); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.piles, tt.h, got, tt.want)
				}
			})
		}
	}
}

// TestMinEatingSpeedAgainstEnumeration 对不同 h 与线性枚举结果比对。
func TestMinEatingSpeedAgainstEnumeration(t *testing.T) {
	piles := []int{7, 1, 19, 4, 12, 30, 2}
	for h := len(piles); h <= 100; h++ {
		want := minEatingSpeedEnumeration(piles, h)
		if got := minEatingSpeed(piles, h); got != want {
			t.Errorf("minEatingSpeed(h=%d) = %d, want %d", h, got, want)
		}
	}
}

// TestMinEatingSpeedLarge 大数据只跑二分。
func TestMinEatingSpeedLarge(t *testing.T) {
	piles := []int{1000000000, 1000000000}
	if got := minEatingSpeed(piles, 3); got != 1000000000 {
		t.Errorf("got %d, want 1000000000", got)
	}
	if got := minEatingSpeed([]int{1000000000}, 1000000000); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
