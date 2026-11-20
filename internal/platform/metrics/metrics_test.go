package metrics

import "testing"

func TestRegistry(t *testing.T) {
	r := New()
	r.Inc("a")
	r.Inc("a")
	r.Inc("b")
	snap := r.Snapshot()
	if snap["a"] != 2 || snap["b"] != 1 {
		t.Fatalf("snapshot = %+v", snap)
	}
	var nilRegistry *Registry
	nilRegistry.Inc("x")
	if nilRegistry.Snapshot() != nil {
		t.Fatal("nil registry snapshot")
	}
}
