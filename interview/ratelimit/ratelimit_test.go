package ratelimit

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTokenBucketAllow(t *testing.T) {
	// 容量 3，速率 0：先允许 3 个（突发），之后全部拒绝。
	b := NewTokenBucket(0, 3)
	for i := 0; i < 3; i++ {
		if !b.Allow() {
			t.Fatalf("第 %d 个令牌应被放行", i+1)
		}
	}
	if b.Allow() {
		t.Error("容量耗尽后仍被放行，限流失效")
	}
}

func TestTokenBucketRefill(t *testing.T) {
	// 容量 1，速率 100/s，即每 10ms 补满 1 个。
	b := NewTokenBucket(100, 1)
	if !b.Allow() {
		t.Fatal("初始令牌应放行")
	}
	if b.Allow() {
		t.Fatal("初始容量为 1，第二次应被拒绝")
	}
	// 手动回拨时间模拟过去 20ms，令牌已补充。
	b.last = b.last.Add(-20 * time.Millisecond)
	if !b.Allow() {
		t.Error("补充令牌后应被放行")
	}
}

func TestRateLimitHandler(t *testing.T) {
	h := RateLimit(NewTokenBucket(0, 1), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	h(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("第一次请求应放行，status = %d", rec.Code)
	}
	rec = httptest.NewRecorder()
	h(rec, req)
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("第二次请求应被限流 429，status = %d", rec.Code)
	}
}
