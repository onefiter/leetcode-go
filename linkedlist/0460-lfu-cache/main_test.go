package lfucache

import "testing"

// op 一次操作：name 为 "get" 或 "put"；get 时 want 为期望返回值。
type op struct {
	name       string
	key, value int
	want       int
}

func TestLFUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []op
	}{
		{"示例1", 2, []op{
			{"put", 1, 1, 0}, {"put", 2, 2, 0}, {"get", 1, 0, 1},
			{"put", 3, 3, 0}, {"get", 2, 0, -1}, {"get", 3, 0, 3},
			{"put", 4, 4, 0}, {"get", 1, 0, -1}, {"get", 3, 0, 3}, {"get", 4, 0, 4},
		}},
		{"容量为0", 0, []op{{"put", 0, 0, 0}, {"get", 0, 0, -1}}},
		{"容量为1", 1, []op{
			{"put", 1, 1, 0}, {"get", 1, 0, 1}, {"put", 2, 2, 0},
			{"get", 1, 0, -1}, {"get", 2, 0, 2},
		}},
		{"空缓存get", 2, []op{{"get", 1, 0, -1}}},
		{"同频次淘汰最久未使用", 2, []op{
			{"put", 1, 1, 0}, {"put", 2, 2, 0}, {"put", 3, 3, 0},
			{"get", 1, 0, -1}, {"get", 2, 0, 2}, {"get", 3, 0, 3},
		}},
		{"更新值也增加频次", 2, []op{
			{"put", 1, 1, 0}, {"put", 2, 2, 0}, {"put", 1, 10, 0},
			{"put", 3, 3, 0}, {"get", 2, 0, -1}, {"get", 1, 0, 10}, {"get", 3, 0, 3},
		}},
		{"新key插入后minFreq重置为1", 2, []op{
			{"put", 1, 1, 0}, {"get", 1, 0, 1}, {"get", 1, 0, 1},
			{"put", 2, 2, 0}, {"put", 3, 3, 0}, // 淘汰频次为 1 的 key 2
			{"get", 2, 0, -1}, {"get", 1, 0, 1}, {"get", 3, 0, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor(tt.capacity)
			for i, o := range tt.ops {
				switch o.name {
				case "put":
					c.Put(o.key, o.value)
				case "get":
					if got := c.Get(o.key); got != o.want {
						t.Errorf("op#%d Get(%d) = %d, want %d", i, o.key, got, o.want)
					}
				}
			}
		})
	}
}
