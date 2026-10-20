package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	c := New(2)
	c.Put("a", 1)
	c.Put("b", 2)

	if got := c.Get("a"); got != 1 {
		t.Fatalf("Get(a) = %d, want 1", got)
	}

	// 插入 c 后容量超限，应淘汰最久未使用的 b（a 刚被访问过）。
	c.Put("c", 3)
	if got := c.Get("b"); got != -1 {
		t.Errorf("b 应被淘汰，Get(b) = %d, want -1", got)
	}
	if got := c.Get("a"); got != 1 {
		t.Errorf("Get(a) = %d, want 1", got)
	}
	if got := c.Get("c"); got != 3 {
		t.Errorf("Get(c) = %d, want 3", got)
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	c := New(2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 10) // 更新 a，不触发淘汰

	if got := c.Get("a"); got != 10 {
		t.Errorf("Get(a) = %d, want 10", got)
	}
	if got := c.Get("b"); got != 2 {
		t.Errorf("Get(b) = %d, want 2", got)
	}
}
