package recentcounter

import "testing"

func TestRecentCounter(t *testing.T) {
	tests := []struct {
		name  string
		pings []int
		want  []int
	}{
		{"示例1", []int{1, 100, 3001, 3002}, []int{1, 2, 3, 3}},
		{"单次请求", []int{5}, []int{1}},
		{"恰好3000边界", []int{1, 3001}, []int{1, 2}},
		{"超出3000边界", []int{1, 3002}, []int{1, 1}},
		{"间隔很大全部过期", []int{1, 10000, 20000, 30000}, []int{1, 1, 1, 1}},
		{"密集请求", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"逐步滑出", []int{1000, 2000, 3000, 4000, 5000, 7001}, []int{1, 2, 3, 4, 4, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := Constructor()
			for i, p := range tt.pings {
				if got := rc.Ping(p); got != tt.want[i] {
					t.Fatalf("第%d步 Ping(%d) = %d, want %d", i, p, got, tt.want[i])
				}
			}
		})
	}
}
