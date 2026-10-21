package dicerolls

import "testing"

var solutions = map[string]func(int, int, int) int{
	"numRollsToTarget":     numRollsToTarget,
	"numRollsToTargetMemo": numRollsToTargetMemo,
}

func TestNumRollsToTarget(t *testing.T) {
	tests := []struct {
		name         string
		n, k, target int
		want         int
	}{
		{"示例1", 1, 6, 3, 1},
		{"示例2", 2, 6, 7, 6},
		{"示例3取模", 30, 30, 500, 222616187},
		{"目标过小", 3, 6, 2, 0},
		{"目标过大", 2, 6, 13, 0},
		{"全取最大", 3, 4, 12, 1},
		{"单面骰子", 5, 1, 5, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, tt.k, tt.target); got != tt.want {
					t.Errorf("%s(%d, %d, %d) = %d, want %d", name, tt.n, tt.k, tt.target, got, tt.want)
				}
			})
		}
	}
}
