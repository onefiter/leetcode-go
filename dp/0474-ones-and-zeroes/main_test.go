package onesandzeroes

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]string, int, int) int{
	"findMaxForm":     findMaxForm,
	"findMaxFormMemo": findMaxFormMemo,
}

func TestFindMaxForm(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		m, n int
		want int
	}{
		{"示例1", []string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{"示例2", []string{"10", "0", "1"}, 1, 1, 2},
		{"容量为零", []string{"0", "1"}, 0, 0, 0},
		{"只能放0", []string{"0", "0", "1", "1"}, 2, 0, 2},
		{"单个装不下", []string{"000"}, 2, 2, 0},
		{"全部可选", []string{"01", "10", "0", "1"}, 3, 3, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.strs), tt.m, tt.n); got != tt.want {
					t.Errorf("%s(%v, %d, %d) = %d, want %d", name, tt.strs, tt.m, tt.n, got, tt.want)
				}
			})
		}
	}
}
