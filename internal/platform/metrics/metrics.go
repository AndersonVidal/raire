package metrics

import (
	"sort"
	"sync"
)

// Registry is a minimal, dependency-free counter registry for the MVP. Export
// to the operational backend is a follow-up in FND-10.
type Registry struct {
	mu       sync.RWMutex
	counters map[string]int64
}

func New() *Registry {
	return &Registry{counters: make(map[string]int64)}
}

func (r *Registry) Inc(name string) {
	if r == nil {
		return
	}
	r.mu.Lock()
	r.counters[name]++
	r.mu.Unlock()
}

func (r *Registry) Snapshot() map[string]int64 {
	if r == nil {
		return nil
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[string]int64, len(r.counters))
	keys := make([]string, 0, len(r.counters))
	for k := range r.counters {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		out[k] = r.counters[k]
	}
	return out
}
