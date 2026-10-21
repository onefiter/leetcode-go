package firstbadversion

import "testing"

var solutions = map[string]func(int) int{
	"firstBadVersion":            firstBadVersion,
	"firstBadVersionEnumeration": firstBadVersionEnumeration,
}

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name   string
		n, bad int
	}{
		{"示例1", 5, 4},
		{"示例2-单版本", 1, 1},
		{"第一个就坏", 10, 1},
		{"最后一个坏", 10, 10},
		{"两个版本-第二个坏", 2, 2},
		{"中间", 100, 37},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				isBadVersion = func(v int) bool { return v >= tt.bad }
				if got := fn(tt.n); got != tt.bad {
					t.Errorf("%s(n=%d, bad=%d) = %d", name, tt.n, tt.bad, got)
				}
			})
		}
	}
}

// TestFirstBadVersionLarge 大 n 只跑二分，验证不溢出、调用次数为对数级。
func TestFirstBadVersionLarge(t *testing.T) {
	n, bad := 2147483647, 2147483647
	calls := 0
	isBadVersion = func(v int) bool { calls++; return v >= bad }
	if got := firstBadVersion(n); got != bad {
		t.Errorf("firstBadVersion(%d) = %d, want %d", n, got, bad)
	}
	if calls > 32 {
		t.Errorf("isBadVersion called %d times, want <= 32", calls)
	}
}
