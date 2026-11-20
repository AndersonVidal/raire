package ratelimit

import (
	"sync"
	"time"

	"github.com/raire/raire/internal/platform/clock"
)

type bucket struct {
	count   int
	resetAt time.Time
}

// Limiter is an in-process fixed-window limiter. It is intentionally simple for
// the single-instance MVP; a shared store is a documented scale-out follow-up.
type Limiter struct {
	mu      sync.Mutex
	buckets map[string]*bucket
	clock   clock.Clock
	maxKeys int
}

func New(c clock.Clock) *Limiter {
	if c == nil {
		c = clock.System{}
	}
	return &Limiter{buckets: make(map[string]*bucket), clock: c, maxKeys: 50_000}
}

// Allow records one attempt for key and reports whether it is allowed.
func (l *Limiter) Allow(key string, limit int, window time.Duration) (bool, int, time.Duration) {
	if limit <= 0 || window <= 0 {
		return true, 0, 0
	}
	now := l.clock.Now()
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.buckets) > l.maxKeys {
		for k, b := range l.buckets {
			if !now.Before(b.resetAt) {
				delete(l.buckets, k)
			}
		}
	}
	b, ok := l.buckets[key]
	if !ok || !now.Before(b.resetAt) {
		b = &bucket{count: 0, resetAt: now.Add(window)}
		l.buckets[key] = b
	}
	if b.count >= limit {
		return false, 0, b.resetAt.Sub(now)
	}
	b.count++
	return true, limit - b.count, b.resetAt.Sub(now)
}

// Reset clears a key, useful after a successful login.
func (l *Limiter) Reset(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.buckets, key)
}
