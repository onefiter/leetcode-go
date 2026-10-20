package ratelimit

import (
	"net/http"
	"sync"
	"time"
)

// TokenBucket 令牌桶限流器。
// 思想：桶里存令牌，容量为 cap；以固定速率 rate 往桶里加令牌（桶满则丢弃）。
// 每个请求要先拿走一个令牌，拿不到就被限流。令牌桶允许一定程度的突发：
// 桶满时最多能一次性放行 cap 个请求，之后按 rate 匀速补充。
type TokenBucket struct {
	rate   float64 // 每秒补充的令牌数
	cap    float64 // 桶容量（允许的最大突发）
	tokens float64 // 当前令牌数
	last   time.Time
	mu     sync.Mutex
}

func NewTokenBucket(rate float64, cap float64) *TokenBucket {
	return &TokenBucket{rate: rate, cap: cap, tokens: cap, last: time.Now()}
}

// Allow 尝试取一个令牌，拿到返回 true，否则返回 false。
func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	// 按经过的时间补充令牌，但不超过容量。
	b.tokens += now.Sub(b.last).Seconds() * b.rate
	if b.tokens > b.cap {
		b.tokens = b.cap
	}
	b.last = now
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

// RateLimit 用一个中间层给任意 Handler 加限流。
// 被限流的请求直接返回 429 Too Many Requests。
func RateLimit(bucket *TokenBucket, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bucket.Allow() {
			http.Error(w, "请求过于频繁，请稍后再试", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}
