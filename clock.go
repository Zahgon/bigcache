package bigcache

type clock interface {
	Epoch() int64
}

type systemClock struct {
}

func (c systemClock) Epoch() int64 { _ = "STUB: not implemented"; return 0 }
