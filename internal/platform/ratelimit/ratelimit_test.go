package ratelimit

import (
	"testing"
	"time"

	"github.com/raire/raire/internal/platform/clock"
)

func TestLimiter(t *testing.T) {
	now := time.Unix(1_700_000_000, 0).UTC()
	c := clock.Fixed{T: now}
	l := New(c)
	for i := 0; i < 5; i++ {
		ok, remaining, _ := l.Allow("k", 5, time.Minute)
		if !ok {
			t.Fatalf("attempt %d blocked", i+1)
		}
		if remaining != 4-i {
			t.Fatalf("remaining = %d", remaining)
		}
	}
	ok, _, retry := l.Allow("k", 5, time.Minute)
	if ok {
		t.Fatal("sixth attempt allowed")
	}
	if retry <= 0 {
		t.Fatal("expected retry duration")
	}
	// Fixed window based on the injected clock; advancing requires a new limiter
	// instance or a mutable test clock, so verify isolation instead.
	if ok, _, _ := l.Allow("other", 5, time.Minute); !ok {
		t.Fatal("key isolation failed")
	}
}

func TestLimiterReset(t *testing.T) {
	l := New(clock.System{})
	if ok, _, _ := l.Allow("k", 1, time.Minute); !ok {
		t.Fatal("first attempt blocked")
	}
	if ok, _, _ := l.Allow("k", 1, time.Minute); ok {
		t.Fatal("second attempt allowed")
	}
	l.Reset("k")
	if ok, _, _ := l.Allow("k", 1, time.Minute); !ok {
		t.Fatal("reset did not clear")
	}
}
