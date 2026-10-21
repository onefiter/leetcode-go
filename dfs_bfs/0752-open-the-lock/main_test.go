package openlock

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]string, string) int{
	"openLock":    openLock,
	"openLockBFS": openLockBFS,
}

func TestOpenLock(t *testing.T) {
	tests := []struct {
		name     string
		deadends []string
		target   string
		want     int
	}{
		{"示例1", []string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{"示例2", []string{"8888"}, "0009", 1},
		{"示例3-终点被围死", []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
		{"示例4-起点是死锁", []string{"0000"}, "8888", -1},
		{"目标即起点", []string{}, "0000", 0},
		{"无死锁最远状态", []string{}, "5555", 20},
		{"循环拨动", []string{}, "9999", 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.deadends), tt.target)
				if got != tt.want {
					t.Errorf("%s(%v, %q) = %d, want %d", name, tt.deadends, tt.target, got, tt.want)
				}
			})
		}
	}
}
