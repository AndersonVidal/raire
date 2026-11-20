package clock

import "time"

// Clock allows deterministic timestamps in tests.
type Clock interface {
	Now() time.Time
}

type System struct{}

func (System) Now() time.Time { return time.Now().UTC() }

// Fixed is a clock frozen at a point in time.
type Fixed struct {
	T time.Time
}

func (f Fixed) Now() time.Time { return f.T.UTC() }
